#!/bin/bash
set -e

oapi-codegen -generate types -o "./internal/handlers/http/openapi_types.gen.go" -package "http" "./api/app/http.yml"
oapi-codegen -generate chi-server -o "./internal/handlers/http/openapi_api.gen.go" -package "http" "./api/app/http.yml"
