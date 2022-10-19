package prefixedclient

// This is an example of how to add a prefix to the name of the generated Client struct
// See https://github.com/diebietse/oapi-codegen/issues/785 for why this might be necessary

//go:generate go run github.com/diebietse/oapi-codegen/cmd/oapi-codegen -config cfg.yaml api.yaml
