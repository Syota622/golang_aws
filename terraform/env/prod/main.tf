# module call
module "network" {
  source = "../../module/network"
  pj     = local.pj
  env    = local.env
}
