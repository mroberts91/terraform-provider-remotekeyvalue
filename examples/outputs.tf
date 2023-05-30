output "foo_value" {
  value = data.remotekeyvalue_pair.foo.value
}

output "foo_key" {
  value = data.remotekeyvalue_pair.foo.key
}

output "foo_sensitive" {
  value = data.remotekeyvalue_pair.foo.sensitive
}

output "biz_value" {
  value = data.remotekeyvalue_pair.biz.value
}

output "biz_key" {
  value = data.remotekeyvalue_pair.biz.key
}

output "biz_sensitive" {
  value = data.remotekeyvalue_pair.biz.sensitive
}
