resource "aws_cognito_user_pool" "this" {
  name                     = "${var.tags.project}-${var.tags.environment}-user-pool"
  mfa_configuration        = "OFF"
  deletion_protection      = "ACTIVE"
  auto_verified_attributes = ["email"]
  username_attributes      = ["email"]

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }

  admin_create_user_config {
    allow_admin_create_user_only = false
  }

  password_policy {
    minimum_length                   = 8
    require_lowercase                = true
    require_numbers                  = true
    require_symbols                  = true
    require_uppercase                = true
    temporary_password_validity_days = 2
  }

  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "given_name"
    required            = false
    string_attribute_constraints {
      min_length = 1
      max_length = 2048
    }
  }

  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "family_name"
    required            = false
    string_attribute_constraints {
      min_length = 1
      max_length = 2048
    }
  }

  schema {
    attribute_data_type = "DateTime"
    mutable             = true
    name                = "birthdate"
    required            = false
  }


  tags = merge(var.tags, {
    "Name" = "${var.tags.project}-${var.tags.environment}-user-pool"
  })
}
