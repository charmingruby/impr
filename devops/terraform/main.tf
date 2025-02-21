module "remote-state" {
  source = "./modules/aws/remote-state"
  tags   = local.tags
}

module "identity" {
  source = "./modules/aws/identity"
  tags   = local.tags
}
