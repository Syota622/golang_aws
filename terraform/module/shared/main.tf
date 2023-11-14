### terraform バックエンド管理サービス ###
# module "tf_backend" {
#   source = "../tf_backend"
#   pj     = var.pj
#   env    = var.env
# }

### ネットワーク類 ###
module "network" {
  source             = "../rails_network"
  pj                 = var.pj
  env                = var.env
}

### ランディングページ ###
module "landing_page" {
  source             = "../landing_page"
  pj                 = var.pj
  env                = var.env
}