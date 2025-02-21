resource "aws_dynamodb_table" "this" {
  name         = format("%s-tf-lock", var.tags.project)
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
