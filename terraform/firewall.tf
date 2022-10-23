module "firewall_rules" {
  source       = "terraform-google-modules/network/google//modules/firewall-rules"
  project_id   = var.project_name
  network_name = var.vpc_name

  rules = [
    {
        name                    = "allow-http-ingress"
        description             = "allow http and https to access api"
        direction               = "INGRESS"
        priority                = 1000
        ranges                  = ["0.0.0.0/0"]
        source_tags             = null
        source_service_accounts = null
        target_tags             = ["allow-http"]
        target_service_accounts = null
        allow = [{
        protocol = "tcp"
        ports    = ["22","80","443"]
        }]
        deny = []
        log_config = {
        metadata = "INCLUDE_ALL_METADATA"
        }
    },
    {
        name                    = "allow-mysql-ingress"
        description             = "allow mysql access from api instance only"
        direction               = "INGRESS"
        priority                = 1000
        ranges                  = ["10.128.0.0/9"]
        source_tags             = null
        source_service_accounts = null
        target_tags             = ["allow-mysql"]
        target_service_accounts = null
        allow = [{
        protocol = "tcp"
        ports    = ["3306"]
        }]
        deny = []
        log_config = {
        metadata = "INCLUDE_ALL_METADATA"
        }
    },
    {
        name                    = "allow-ssh-ingress-temp"
        description             = "allow ssh access from internet"
        direction               = "INGRESS"
        priority                = 1000
        ranges                  = ["0.0.0.0/0"]
        source_tags             = null
        source_service_accounts = null
        target_tags             = ["allow-mysql"]
        target_service_accounts = null
        allow = [{
        protocol = "tcp"
        ports    = ["22"]
        }]
        deny = []
        log_config = {
        metadata = "INCLUDE_ALL_METADATA"
        }
    }
    ]
}