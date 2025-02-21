locals {
  tags = {
    project     = var.project
    department  = var.department
    environment = var.environment
    managed-by  = var.managed-by
    created-at  = var.created-at
  }
}
