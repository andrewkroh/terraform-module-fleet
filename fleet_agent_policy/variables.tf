variable "name" {
  description = "Name of Agent policy."
  type        = string
}

variable "description" {
  description = "Description of Agent policy."
  type        = string
}

variable "namespace" {
  description = "Namespace for the policy. Controls namespace of monitoring data."
  type        = string
  default     = "default"
}

variable "monitor_logs" {
  default = true
}

variable "monitor_metrics" {
  default = true
}

variable "inactivity_timeout_sec" {
  description = "An agent will automatically change to inactive status and be filtered out"
  default     = 3600
}

variable "data_output_id" {
  description = "ID of the Fleet data output. Use this to choose a non-default output."
  type        = string
  default     = null
}

variable "monitoring_output_id" {
  description = "ID of the Fleet monitoring output. Use this to choose a non-default output."
  type        = string
  default     = null
}

variable "fleet_server_host_id" {
  description = "ID of the Fleet server host. Use this to choose a non-default Fleet server address."
  type        = string
  default     = null
}