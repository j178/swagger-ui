# Go package that embeds swagger-ui source

## Usage
```shell
go get github.com/j178/swagger-ui
```

```go
package main

import "net/http"
import "github.com/j178/swagger-ui"

func main() {
	http.Handle("/swagger/", http.StripPrefix("/swagger/", swagger_ui.FS))
	http.ListenAndServe(":8000", nil)
}
```
