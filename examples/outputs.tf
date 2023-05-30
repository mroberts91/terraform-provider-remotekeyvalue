output "connection_string" {
  value = data.remotekeyvalue_pair.foo.value
}

output "connection_string_key" {
  value = data.remotekeyvalue_pair.foo.key
}

output "connection_string_sensitive" {
  value = data.remotekeyvalue_pair.foo.sensitive
}
