data "ignition_config" "worker" {
  networkd = [
    "${module.ignition_networkd.10_internal0_link}",
    "${module.ignition_networkd.20_internal0_network}",
  ]

  users = [
    "${module.ignition_users.demo}",
  ]

  systemd = [
    "${module.ignition_systemd.consulw}",
  ]

  files = [
    "${module.ignition_files.hostname}",
    "${module.ignition_files.bootstrap-consulw}",
  ]
}

module "ignition_networkd" {
  source               = "../../ignition/networkd"
  internal_mac_address = "${var.mac_address}"
}

module "ignition_users" {
  source = "../../ignition/users"
}

module "ignition_systemd" {
  source               = "../../ignition/systemd/worker"
  consul_name          = "${element(split(".", var.hostname), 0)}"
  consul_server_cluster = "${var.consul_server_cluster}"
}

module "ignition_files" {
  source   = "../../ignition/files/worker"
  hostname = "${var.hostname}"
}
