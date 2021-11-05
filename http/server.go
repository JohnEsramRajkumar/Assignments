package main
import (
         "net/http"
)
type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    // create response binary data
    data := []byte("Hello World....!") 
    // write `data` to response
    res.Write(data)
}
func main() {
    // create a new handler
    handler := HttpHandler{}
    // listen and serve
    http.ListenAndServe(":9010", handler)
}