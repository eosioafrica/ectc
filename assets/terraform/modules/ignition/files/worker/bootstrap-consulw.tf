data "ignition_file" "bootstrap-consulw" {
  filesystem = "root"
  path       = "/etc/bootstrap-consulw"
  mode       = "0644"

  content {
	content = "${file("${path.module}/bootstrap-consulw.conf")}"
  }
}

output "bootstrap-consulw" {
  value = "${data.ignition_file.bootstrap-consulw.id}"
}
