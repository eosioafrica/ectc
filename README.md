Automates the creation of a cluster environment for EOS high availabity test.

Virtualbox for machines. Consul for consensus.

1 * Provisioner (Runs Ubuntu 16.04. Stores provisioning data for cluster. Allows nodes to provision fully with MAC address)

3 * Consul cluster server nodes (Runs CoreOS. Producer election and nodeos health checks)

2 * EOS Producers (Runs CoreOS. Eos runs on docker, at startup. Consul clients connected to server cluster)



