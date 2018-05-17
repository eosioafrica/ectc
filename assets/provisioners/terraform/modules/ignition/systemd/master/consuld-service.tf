data "template_file" "consuld" {
  template = "${file("${path.module}/files/consuld.service.tpl")}"
}



data "ignition_systemd_unit" "consuld" {

  name = "consuld.service"
  enabled = true
  content = "${data.template_file.consuld.rendered}"
}

# Ignition config include the previous defined systemd unit data resource
data "ignition_config" "consuld" {

  systemd = [
       "${data.ignition_systemd_unit.consuld.id}",
  ]
}


output "consuld" {
  value = "${data.ignition_systemd_unit.consuld.id}"
}
