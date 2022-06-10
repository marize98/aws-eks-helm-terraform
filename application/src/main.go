package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", HolaEndpoint)
    http.ListenAndServe(":8080", nil)
}

func HelloEndpoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s, %s!", os.Getenv("HOLA"), r.URL.Path[1:])
}
