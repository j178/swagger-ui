package swagger_ui

import "github.com/j178/swagger-ui/internal/statik"

//go:generate github.com/FZambia/statik/fs@master -src=./swagger-ui -dest=./internal -f
var FS = statik.FS
