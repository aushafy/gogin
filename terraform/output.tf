output "ip_api_vm" {
    value = google_compute_instance.vm-api.network_interface.0.access_config.0.nat_ip
}

output "ip_mysql_vm" {
    value = google_compute_instance.vm-mysql.network_interface.0.access_config.0.nat_ip
}