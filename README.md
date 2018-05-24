Automates the creation of a emulated cluster environment for EOS.IO high availabity. The purpose is to aid the development of a native high availability solution for EOS.IO.

Virtualbox for machines. Consul for consensus.

1 * Provisioner (Runs Ubuntu 16.04. Stores full provisioning data for cluster. Allows nodes to provision fully with MAC address). This means that, at deploy, iPXE will be suffecient to maintain the booting of cluster nodes. All node setup will reside on this machine(s)

3 * Consul cluster server nodes (Runs CoreOS. Producer election and nodeos health checks)

2 * EOS Producers nodes (Runs CoreOS. Eos runs on docker, at startup. Consul clients connected to server cluster)

1 * Private network



