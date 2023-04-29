variable "name" {
  description = "Name of output."
  type        = string
}

variable "insecure" {
  description = "Disable TLS verification."
  default     = false
}

variable "elasticsearch_url" {
  description = "Elasticsearch URL"
  type        = string
}
