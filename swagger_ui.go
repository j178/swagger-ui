package swagger_ui

import "github.com/j178/swagger-ui/internal/statik"

//go:generate go run github.com/FZambia/statik@master -src=./dist -dest=./internal -f
var FS = statik.FS
