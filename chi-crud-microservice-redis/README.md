# 1
```
### Do after doing the following

$ go build main.go && ./main
```
```go
package main

import (
	"fmt"
	"net/http"
)

func main () {
	server := &http.Server {
		Addr: ":10086",
		Handler: http.HandlerFunc( handler ),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("err/listen & serve", err)
	}
}

func handler (writer http.ResponseWriter, req *http.Request) {
	writer.Write( []byte("Hello, World!") )
}
```
