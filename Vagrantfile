Vagrant.configure("2") do |config|
  config.vm.box = "etherframe/alpine64"
  config.vm.network "public_network"
  config.vm.provider "hyperv" do |h|
    h.memory = "2048"
    h.cpus = "2"
    config.ssh.shell = "sh"
    h.vm_integration_services = {
      guest_service_interface: true#,
      #CustomVMSRV: true
    }

  end
end