resource "google_compute_instance" "vm-api" {
  name         = var.vm_api_name
  machine_type = var.machine_type
  zone         = var.zone_name

  tags = ["allow-http"]

  boot_disk {
    initialize_params {
      // gcloud compute image list
      image = "ubuntu-os-cloud/ubuntu-2004-focal-v20221018"
      labels = {
        my_label = "value"
      }
    }
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral public IP
    }
  }

  metadata = {
    foo = "bar"
  }

  metadata_startup_script = "echo hi > /test.txt"


}

