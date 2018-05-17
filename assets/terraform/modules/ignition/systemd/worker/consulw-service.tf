data "template_file" "consulw" {
  template = "${file("${path.module}/files/consulw.service.tpl")}"

}



data "ignition_systemd_unit" "consulw" {

  name = "consulw.service"
  enabled = true
  content = "${data.template_file.consulw.rendered}"
}

# Ignition config include the previous defined systemd unit data resource
data "ignition_config" "consulw" {

  systemd = [
       "${data.ignition_systemd_unit.consulw.id}",
  ]
}


output "consulw" {
  value = "${data.ignition_systemd_unit.consulw.id}"
}
