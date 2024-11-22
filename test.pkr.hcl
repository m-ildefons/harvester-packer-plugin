packer {
  required_plugins {
    harvester = {
      version = "0.0.1"
      source = "packer.local/local/harvester"
    }
  }
}

source "harvester" "test" {
  kubeconfig = "./kubeconfig.yaml"

  namespace = "custom-build"

  communicator {
    communicator = "ssh"
    ssh_username = "ubuntu"
  }

  cpu = 2
  memory = "4GiB"

  volume {
    source {
      type = "cloud-init"
    }
  }
  volume {
    # alias = "artifact"

    source {
      type = "download"
      url = "https://cloud-images.ubuntu.com/releases/releases/22.04/release/ubuntu-22.04-server-cloudimg-amd64.img"
    }
  }
}

build {
  sources = [
    "source.harvester.test",
  ]

  provisioner "shell" {
    inline = [
      "echo foobar > /etc/issue",
      "ip -br a",
    ]
  }
}
