### terraform management S3 and DynamoDB ###
# module "tf_backend" {
#   source = "../tf_backend"
#   pj     = var.pj
#   env    = var.env
# }

### network ###
module "network" {
  source             = "../network"
  pj                 = var.pj
  env                = var.env
}
