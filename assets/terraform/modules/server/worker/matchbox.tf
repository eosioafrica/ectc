provider "matchbox" {
  endpoint    = "${var.matchbox_rpc_endpoint}"
  client_cert = "${file("/home/ubuntu/.matchbox/infra-local/client.crt")}"
  client_key  = "${file("/home/ubuntu/.matchbox/infra-local/client.key")}"
  ca          = "${file("/home/ubuntu/.matchbox/infra-local/ca.crt")}"
}

resource "matchbox_group" "worker" {
  name    = "${var.hostname}"
  profile = "${matchbox_profile.worker.name}"

  selector {
    mac = "${var.mac_address}"
  }
}

resource "matchbox_profile" "worker" {
  name   = "${var.hostname}"
  kernel = "/assets/coreos/${var.container_linux_version}/coreos_production_pxe.vmlinuz"

  initrd = [
    "/assets/coreos/${var.container_linux_version}/coreos_production_pxe_image.cpio.gz",
  ]

  args = [
    "coreos.config.url=http://${var.matchbox_http_endpoint}/ignition?uuid=$${uuid}&mac=$${mac:hexhyp}",
    "coreos.first_boot=yes",
    "console=tty0",
    "initrd=coreos_production_pxe_image.cpio.gz",
  ]

  raw_ignition = "${data.ignition_config.worker.rendered}"
}
