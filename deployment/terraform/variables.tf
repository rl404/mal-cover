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
