provider "google" {
  credentials = file("service-account.json")
  project     = var.project_name
  region      = var.region_name
  zone        = var.zone_name
}