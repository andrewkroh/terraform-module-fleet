terraform {
  required_version = ">= 1.4"
  required_providers {
    elasticstack = {
      source  = "elastic/elasticstack"
      version = "~> 0.11"
    }
    restapi = {
      source  = "Mastercard/restapi"
      version = "~> 1.18"
    }
  }
}
