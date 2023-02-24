package v1

import "net/http"

func PageHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello World!"))
}
