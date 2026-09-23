provider "rancher2" {
  api_url    = local.rancher_url
  access_key = var.rancher_credentials.access_key
  secret_key = var.rancher_credentials.secret_key
  insecure   = var.rancher_insecure
}

locals {
  cluster_name       = var.cluster_name
  rancher_url        = var.rancher_url
  kubernetes_version = var.kubernetes_version

  pve_cloud_credential_name = var.pve_cloud_credential_name

  server = {
    template_vmid = var.pve_server_template_vmid
    sockets       = var.pve_server_sockets
    cores         = var.pve_server_cores
    memory        = var.pve_server_memory
    quantity      = var.server_quantity
  }

  worker = {
    template_vmid = var.pve_worker_template_vmid
    sockets       = var.pve_worker_sockets
    cores         = var.pve_worker_cores
    memory        = var.pve_worker_memory
    quantity      = var.worker_quantity
  }

  pve = {
    allowed_nodes = var.pve_allowed_nodes
    net_device    = var.pve_net_device
    ssh_user      = var.pve_ssh_user
    linked_clone  = var.pve_linked_clone
    tags          = var.pve_tags
  }

}

# Look up existing cloud credential already configured in Rancher
# (URL, API token and resource pool all live on this credential)
data "rancher2_cloud_credential" "pve" {
  name = local.pve_cloud_credential_name
}

resource "rancher2_machine_config_v2" "server" {
  generate_name = "${local.cluster_name}-server"

  pve_config {
    pve_template_vmid = local.server.template_vmid

    # Host selection: restrict server nodes to this subset of the cluster
    pve_allowed_nodes = local.pve.allowed_nodes

    pve_net_device   = local.pve.net_device
    pve_ssh_user     = local.pve.ssh_user
    pve_sockets      = local.server.sockets
    pve_cores        = local.server.cores
    pve_memory       = local.server.memory
    pve_linked_clone = local.pve.linked_clone
    pve_tags         = local.pve.tags
  }
}

resource "rancher2_machine_config_v2" "worker" {
  generate_name = "${local.cluster_name}-worker"

  pve_config {
    pve_template_vmid = local.worker.template_vmid

    # Host selection: restrict worker nodes to this subset of the cluster
    pve_allowed_nodes = local.pve.allowed_nodes

    pve_net_device   = local.pve.net_device
    pve_ssh_user     = local.pve.ssh_user
    pve_sockets      = local.worker.sockets
    pve_cores        = local.worker.cores
    pve_memory       = local.worker.memory
    pve_linked_clone = local.pve.linked_clone
    pve_tags         = local.pve.tags
  }
}

resource "rancher2_cluster_v2" "cluster" {
  name               = local.cluster_name
  kubernetes_version = local.kubernetes_version

  rke_config {
    machine_global_config = <<-EOT
      cni: cilium
      disable-kube-proxy: false
      etcd-expose-metrics: false
    EOT

    chart_values = <<-EOT
      rke2-cilium: {}
    EOT

    machine_selector_config {
      config = <<-EOT
        kube-proxy-arg:
          - proxy-mode=nftables
        protect-kernel-defaults: false
      EOT
    }

    machine_pools {
      name                         = "server"
      cloud_credential_secret_name = data.rancher2_cloud_credential.pve.id
      control_plane_role           = true
      etcd_role                    = true
      worker_role                  = false
      quantity                     = local.server.quantity
      drain_before_delete          = true

      machine_config {
        kind = rancher2_machine_config_v2.server.kind
        name = rancher2_machine_config_v2.server.name
      }
    }

    machine_pools {
      name                         = "worker"
      cloud_credential_secret_name = data.rancher2_cloud_credential.pve.id
      control_plane_role           = false
      etcd_role                    = false
      worker_role                  = true
      quantity                     = local.worker.quantity
      drain_before_delete          = true

      machine_config {
        kind = rancher2_machine_config_v2.worker.kind
        name = rancher2_machine_config_v2.worker.name
      }
    }

    upgrade_strategy {
      control_plane_concurrency = "1"
      worker_concurrency        = "1"

      control_plane_drain_options {
        enabled                              = false
        force                                = false
        ignore_daemon_sets                   = true
        delete_empty_dir_data                = true
        disable_eviction                     = false
        grace_period                         = -1
        timeout                              = 120
        skip_wait_for_delete_timeout_seconds = 0
      }

      worker_drain_options {
        enabled                              = false
        force                                = false
        ignore_daemon_sets                   = true
        delete_empty_dir_data                = true
        disable_eviction                     = false
        grace_period                         = -1
        timeout                              = 120
        skip_wait_for_delete_timeout_seconds = 0
      }
    }
  }
}
