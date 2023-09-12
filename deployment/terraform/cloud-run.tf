resource "google_cloud_run_v2_service" "server" {
  name     = var.cloud_run_name
  location = var.cloud_run_location
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    labels = {
      app = var.cloud_run_name
    }
    scaling {
      min_instance_count = 0
    }
    containers {
      name  = var.cloud_run_name
      image = var.gcr_image_name
      env {
        name  = "MC_CACHE_DIALECT"
        value = var.mc_cache_dialect
      }
      env {
        name  = "MC_CACHE_ADDRESS"
        value = var.mc_cache_address
      }
      env {
        name  = "MC_CACHE_PASSWORD"
        value = var.mc_cache_password
      }
      env {
        name  = "MC_CACHE_TIME"
        value = var.mc_cache_time
      }
      env {
        name  = "MC_LOG_LEVEL"
        value = var.mc_log_level
      }
      env {
        name  = "MC_LOG_JSON"
        value = var.mc_log_json
      }
      env {
        name  = "MC_NEWRELIC_LICENSE_KEY"
        value = var.mc_newrelic_license_key
      }
    }
  }
}

resource "google_cloud_run_service_iam_binding" "noauth" {
  service  = google_cloud_run_v2_service.server.name
  location = google_cloud_run_v2_service.server.location
  role     = "roles/run.invoker"
  members  = ["allUsers"]
}
