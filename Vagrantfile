# note: the IP addresses used by these VMs may need some configuration for MacOS
# create this file /etc/vbox/networks.conf with content:
# * 10.0.0.0/8 192.168.0.0/16
# * 2001::/64
# https://www.virtualbox.org/manual/ch06.html#network_hostonly

Vagrant.configure("2") do |config|
  config.vm.box = "hashicorp/bionic64"
  config.vm.provider "virtualbox" do |v|
    v.memory = 2048
    v.cpus = 2
  end


  config.vm.provision "shell", inline: "echo -e '192.168.56.3 node01\n192.168.56.4 node02\n192.168.56.5 node03' | tee -a /etc/hosts"
  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "setup_env.yml"
  end


  config.vm.define "node01" do |node01|
    node01.vm.network "forwarded_port", guest: 80, host: 8000, host_ip: "0.0.0.0"
    node01.vm.network :private_network, ip: "192.168.56.3"
    node01.vm.hostname = "node01"
    node01.vm.network "forwarded_port", id: "ssh", host: 2201, guest: 22
  end

  # config.vm.define "node02" do |node02|
  #   node02.vm.network "forwarded_port", guest: 80, host: 8001, host_ip: "127.0.0.1"
  #   node02.vm.network :private_network, ip: "192.168.56.4"
  #   node02.vm.hostname = "node02"
  #   node02.vm.network "forwarded_port", id: "ssh", host: 2202, guest: 22
  # end

  # config.vm.define "node03" do |node03|
  #   node03.vm.network "forwarded_port", guest: 80, host: 8002, host_ip: "127.0.0.1"
  #   node03.vm.network :private_network, ip: "192.168.56.5"
  #   node03.vm.hostname = "node03"
  #   node03.vm.network "forwarded_port", id: "ssh", host: 2203, guest: 22
  # end
end
