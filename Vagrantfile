Vagrant.configure("2") do |config|
  arch=`uname -m`.strip

  if arch == "arm64"
    puts "Detected Apple Silicon (ARM)"
    config.vm.box = "bento/ubuntu-24.04"
    config.vm.box_version = "202510.26.0"
  else
    puts "Detected Intel/AMD (x86_64)"
    config.vm.box = "ubuntu/jammy64"
  end

  # Sync current host directory to /vagrant in the VM
  config.vm.synced_folder ".", "/home/vagrant/shared"
end

# sudo apt install libbpf-dev
# sudo apt update
# sudo apt install build-essential
# sudo add-apt-repository ppa:longsleep/golang-backports
# sudo apt update
# sudo apt install golang-go