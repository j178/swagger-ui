# Swagger UI embedded in Go binary

This a golang package that embeds swagger-ui source, making it easy to serve swagger-ui from a golang application.

## Usage
```shell
go get github.com/j178/swagger-ui
```

```go
package main

import (
	"net/http"
	"github.com/j178/swagger-ui"
)

func main() {
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(swagger_ui.FS)))
	http.ListenAndServe(":8000", nil)
}
```
