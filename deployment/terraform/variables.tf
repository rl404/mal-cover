variable "gcp_project_id" {
  type        = string
  description = "GCP project id"
}

variable "gcp_region" {
  type        = string
  description = "GCP project region"
}

variable "gcr_image_name" {
  type        = string
  description = "GCR image name"
}

variable "cloud_run_name" {
  type        = string
  description = "Google cloud run name"
}

variable "cloud_run_location" {
  type        = string
  description = "Google cloud run location"
}

variable "mc_cache_dialect" {
  type        = string
  description = "Cache dialect"
}

variable "mc_cache_address" {
  type        = string
  description = "Cache address"
}

variable "mc_cache_password" {
  type        = string
  description = "Cache password"
}

variable "mc_cache_time" {
  type        = string
  description = "Cache time"
}

variable "mc_log_level" {
  type        = number
  description = "Log level"
}

variable "mc_log_json" {
  type        = bool
  description = "Log json"
}

variable "mc_newrelic_license_key" {
  type        = string
  description = "Newrelic license key"
}
