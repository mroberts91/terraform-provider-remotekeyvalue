# Remote Key Value Examples
This example consists of 2 pieces:
1. Go Rest API that is the remote source used in the terraform project. `./api`
2. Terraform Project

## 1. Run the API
Run the API locally.

```bash
# From the root of the repo
cd ./examples/api

# Run the API with API Keys
go run main.go -port 8080 -keyname API_KEY -key 12345

# go run main.go -port 8080 -keyname API_KEY -key 12345
# 2023/05/30 15:44:54 Starting Server at 0.0.0.0:8080...
# 2023/05/30 15:44:54 Supported Routes: [ /api/v1/{key} ]
# 2023/05/30 15:44:54 Available Keys: [ Foo, Biz ]
# 2023/05/30 15:45:03 "GET http://localhost:8080/api/v1/foo HTTP/1.1" from [::1]:38028 - 200 55B in 41.406µs
# 2023/05/30 15:45:04 "GET http://localhost:8080/api/v1/foo HTTP/1.1" from [::1]:38028 - 200 55B in 11.693µs
# 2023/05/30 15:45:08 "GET http://localhost:8080/api/v1/foo HTTP/1.1" from [::1]:38028 - 200 55B in 17.57µs
# 2023/05/30 15:45:08 "GET http://localhost:8080/api/v1/foo HTTP/1.1" from [::1]:38028 - 200 55B in 17.609µs
```

## 2. Execute the terraform project
The terraform project is setup to register the KVP data sources and set outputs from the data sources.

```bash
terraform init
terraform validate
terraform apply --auto-approve

# Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

# Outputs:

# biz_key = "Biz"
# biz_sensitive = true
# biz_value = "Baz"
# foo_key = "Foo"
# foo_sensitive = false
# foo_value = "Bar"
```
