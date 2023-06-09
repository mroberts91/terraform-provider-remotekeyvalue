provider "remotekeyvalue" {
  uri                  = var.uri
  api_key_header_name  = var.api_key_header_name
  api_key_header_value = var.api_key_header_value
  timeout              = var.timeout
}

data "remotekeyvalue_pair" "foo" {
  path = "/api/v1/foo"
}

data "remotekeyvalue_pair" "biz" {
  path = "/api/v1/biz"
}
