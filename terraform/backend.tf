terraform {
  backend "gcs" {
    credentials = "./service-account.json"
    bucket  = "tf-state-1192"
    prefix  = "state"
  }
}