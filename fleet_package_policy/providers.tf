terraform {
  required_providers {
    elasticstack = {
      source  = "elastic/elasticstack"
      version = ">= 0.11.0"
    }
    restapi = {
      source  = "Mastercard/restapi"
      version = ">= 1.18.0"
    }
  }
}
