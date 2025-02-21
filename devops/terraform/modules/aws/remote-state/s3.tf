resource "aws_s3_bucket" "this" {
  bucket = format("%s-tf-state", var.tags.project)
}
