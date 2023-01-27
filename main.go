package main

import (
    "fmt"
    "net/http"
    "os"
	//"crypto/md5"
)

func encryptoHandler(w, http.ResponseWriter, r *http.Request){
	path := r.URL.Path
	parts := strings.Split(path, "/")

	secret := parts[3]
	word := parts[4]
}

func main() {
    var port = "8080"
    if len(os.Args) > 1 {
        port = os.Args[1]
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, World!")
    })
    fmt.Printf("Server is active and listening on port %s\n", port)
    http.ListenAndServe(":"+port, mux)
}
