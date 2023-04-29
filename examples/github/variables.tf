variable "kibana_username" {
  description = "Kibana username"
  type        = string
  default     = "elastic"
}

variable "kibana_password" {
  description = "Kibana password"
  type        = string
  default     = "changeme"
}

variable "kibana_url" {
  description = "Kibana URL"
  type        = string
  default     = "https://localhost:5601"
}

variable "kibana_insecure" {
  description = "Skip the TLS verification step and proceed without checking."
  type        = bool
  default     = true
}

variable "github_token" {
  description = "GitHub personal access token that allows reading issues."
  sensitive   = true
  type        = string
}
