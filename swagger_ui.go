package swagger_ui

import (
	"net/http"

	_ "github.com/j178/swagger-ui/statik"
	"github.com/rakyll/statik/fs"
)

var FS http.FileSystem

func init() {
	FS, _ = fs.New()
}
