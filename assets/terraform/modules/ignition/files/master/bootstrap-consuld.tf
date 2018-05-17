data "ignition_file" "bootstrap-consuld" {
  filesystem = "root"
  path       = "/etc/bootstrap-consuld"
  mode       = "0644"

  content {
	content = "${file("${path.module}/bootstrap-consuld.conf")}"
  }
}

output "bootstrap-consuld" {
  value = "${data.ignition_file.bootstrap-consuld.id}"
}
