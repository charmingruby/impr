resource "aws_cognito_user_pool_client" "this" {
  name         = "${var.tags.project}-${var.tags.environment}-user-pool"
  user_pool_id = aws_cognito_user_pool.this.id
  explicit_auth_flows = [
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_PASSWORD_AUTH"
  ]
  token_validity_units {
    access_token  = "hours"
    refresh_token = "hours"
  }

  access_token_validity  = 3
  refresh_token_validity = 1440 # 60 days
}
