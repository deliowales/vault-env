# vault-env

Golang shim to populate envrioment variables from injected secrets via Vault sidecar on Kubernetes

## Usage

### Getting Values

The application relies on 1 environment variable `SECRETPATH` which is used to create the search path. An example could be `/vault/secrets`. To run it simply call `vault-env` without any parameters and this will cause the variables to printed out to STDOUT. To make them environment variables wrap the command in eval, `eval $(vault-env)`.

## Build

To build the binary run `go build . ` from the directory. 
