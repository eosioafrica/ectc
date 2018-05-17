data "ignition_user" "demo" {
  name                = "demo"
  ssh_authorized_keys = ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC/6sK+8lPoC6F3d7qsj8dwVdysL3TwW8F1PGvqc6ewWyhSQUgw9NOqlY9YmqUL8jozBLTk2uepLXpUhdSAO7ki1C9Y2k+MNyuXxJNkWBULMvfuO4z/qVwmZe3ZmELUi2G0bWxx3OFoZdYmqrzZiETDlRW/S7JLJWjZG/iaTy6tZPB/Nk+058zkbRrLEEzSQ5JfwzgTDbMsKOquW0A0Wa1epJ1w1aFtckJgfs6P2nYDoe4G9vZG28x3iEnRJjqy9JLu6K5ZsHVSNmFu5SW3Gf4aUPxpq5/JrKE8Oo3JJCGGtjmIFq57aZAjpeYOPEX9uzbTczyciTJv4aXsy9hIpFH5 vagrant@provisioner-1"]
  groups              = ["sudo", "docker"]
}

output "demo" {
  value = "${data.ignition_user.demo.id}"
}