***The key idea behind this project is to create easy to install and repeatable testing and development environments for the EOS.IO community. If we have a less than overly onerous test environment, we will be able to get as many people testing all the new features that will be added to the ecosystem.***

Virtualbox for machines. Consul for consensus.

1 * Provisioner (Runs Ubuntu 16.04. Stores full provisioning data for cluster. Allows nodes to provision fully with MAC address). This means that, at deploy, iPXE will be suffecient to maintain the booting of cluster nodes. All node setup will reside on this machine(s)

3 * Consul cluster server nodes (Runs CoreOS. Producer election and nodeos health checks)

2 * EOS Producers nodes (Runs CoreOS. Eos runs on docker, at startup. Consul clients connected to server cluster)

1 * Private network

***Instructions***

- On Ubuntu 18.04. 
- Need 6 Gigs of RAM to run the entire cluster. The point of this exercise will be made by just 3 Gigs. Run one of each category.
- I am forgetting one more thing. Must not be important!!! Moving on...


1 . [Install VirtualBox and Vagrant. Follow this tutorial.](http://www.codebind.com/linux-tutorials/install-vagrant-ubuntu-18-04-lts-linux/)

2. [Install golang from this script] (https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh)

3. git clone github.com/eosioafrica/ectc

4. Add user {ecte}. No special priviledges needed. 

5. Inside the project folder run as sudo (sorry!) : {sudo go run main.go}

6. Pray. hehehehe!

7. If all successful run VirtualBox as sudo and all the machines will be there.

### Task list

- [ ] Change provisioner to CoreOS. Downloads are too big and takes too long to setup. A docker image would work better.
- [ ] Enable EOS systemd service
- [ ] Change consul to handle block.one HA switch
- [ ] Refactor code for composability. For example, consul must be replaceable by etcd.



