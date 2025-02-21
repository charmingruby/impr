variable "project" {
  description = "The project name"
  type        = string
}

variable "department" {
  description = "The project name"
  type        = string
}

variable "environment" {
  description = "The project name"
  type        = string
  default     = "dev"
}

variable "managed-by" {
  description = "The project name"
  type        = string
  default     = "Terraform"
}

variable "created-at" {
  description = "The project name"
  type        = string
}
