variable "agent_policy_id" {
  description = "ID of the agent policy to add the package policy to."
  type        = string
}

variable "package_policy_name" {
  description = "Name of the package policy. Defaults to \"{package_name}-{namespace}.\""
  default     = null
}

variable "package_name" {
  description = "Name of the Fleet package to install (e.g. ti_recordedfuture)."
  type        = string
}

variable "package_version" {
  description = "Package version."
  type        = string
  validation {
    condition     = can(regex("^\\d+\\.\\d+\\.\\d+", var.package_version))
    error_message = "Invalid package_version version."
  }
}

variable "namespace" {
  description = "Namespace to apply to data streams."
  type        = string
  default     = "default"
}

variable "description" {
  description = "Description to apply to the package policy."
  type        = string
  default     = ""
}

variable "policy_template" {
  description = "Name of the policy template (see the integration's manifest.yml)."
  type        = string
}

variable "data_stream" {
  description = "Name of the data_stream within the integration (e.g. \"log\")."
}

variable "input_type" {
  description = "Input type."
  type        = string
}

variable "package_variables_json" {
  description = "JSON encoded package level variables that are shared by each stream."
  type        = string
  default     = "{}"
  sensitive   = true
  validation {
    condition = (
      can(jsondecode(var.package_variables_json))
    )
    error_message = "Must be valid JSON."
  }
}

variable "data_stream_variables_json" {
  description = "JSON encoded data stream specific variables."
  type        = string
  sensitive   = true
  default     = "{}"
  validation {
    condition = (
      can(jsondecode(var.data_stream_variables_json))
    )
    error_message = "Must be valid JSON."
  }
}

variable "input_variables_json" {
  description = "JSON encoded input level variables."
  type        = string
  default     = "{}"
  sensitive   = true
  validation {
    condition = (
      can(jsondecode(var.input_variables_json))
    )
    error_message = "Must be valid JSON."
  }
}

variable "all_data_streams" {
  description = "List of all data streams associated to the input type in the policy template. This is necessary to disable all data streams except the one being used."
  type        = list(string)
  default     = []
}

variable "all_input_types" {
  description = "List of all input types in the package policy template. This is necessary to disable all inputs except the one being used."
  type        = list(string)
  default     = []
}


