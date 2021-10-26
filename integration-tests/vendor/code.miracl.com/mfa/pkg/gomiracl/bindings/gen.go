package bindings

//go:generate go run ../gen/wrappers/main.go ecdsa ecdsa_wrappers.go.tmpl
//go:generate go run ../gen/wrappers/main.go mpin mpin_wrappers.go.tmpl
//go:generate go run ../gen/wrappers/main.go pbc pbc_wrappers.go.tmpl
//go:generate go run ../gen/wrappers/main.go rand rand_wrappers.go.tmpl
//go:generate go run ../gen/wrappers/main.go rsa rsa_wrappers.go.tmpl
//go:generate go run ../gen/wrappers/main.go wcc wcc_wrappers.go.tmpl
//go:generate gofmt -s -w .
