# module call
module "shared" {
  source = "../../module/shared"
  pj     = local.pj
  env    = local.env
}
