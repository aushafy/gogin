variable "project_name" {
    type = string
    default = "sirclo-data-nonprod"
}

variable "region_name" {
    type = string
    default = "asia-southeast1"
}

variable "zone_name" {
    type = string
    default = "asia-southeast1-c"
}

variable "vpc_name" {
    type = string
    default = "default"
}

variable "cluster_name" {
    type = string
    default = "data-prod"
}

variable "disk_type" {
    type = string
    default = "pd-standard"
}

variable "machine_type" {
    type = string
    default = "n2-standard-4"
}

variable "vm_api_name" {
    type = string
    default = "vm-api"
}

variable "vm_mysql_name" {
    type = string
    default = "vm-mysql"
}

variable "preemptible" {
    type = bool
    default = false
}

variable "spot" {
    type = bool
    default = true
}

variable "node_count" {
    type = number
    default = 1
}

variable "max_node_count" {
    type = number
    default = 3
}

variable "min_node_count" {
    type = number
    default = 1
}

variable "disk_size_gb" {
    type = number
    default = 40
}

variable "kubernetes_version" {
    type = string
    default = "1.24.2-gke.1900"
}

variable "cluster_autoscaling" {
    type = bool
    default = "true"
}

variable "memory_resource_limits" {
    type = number
    default = 72
}

variable "cpu_resource_limits" {
    type = number
    default = 36
}

variable "max_node_autoscale" {
    type = number
    default = 3
}