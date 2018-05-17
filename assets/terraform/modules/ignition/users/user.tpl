data "ignition_user" "demo" {
  name                = "demo"
  ssh_authorized_keys = ["${PUBKEY}"]
  groups              = ["sudo", "docker"]
}

output "demo" {
  value = "${data.ignition_user.demo.id}"
}