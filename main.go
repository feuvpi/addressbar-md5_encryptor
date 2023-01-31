package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func encryptoHandler(w http.ResponseWriter, r *http.Request) {

	// extract variables from URL path
	vars := mux.Vars(r)
	secret := vars["secret"]
	word := vars["word"]

	// encrypt
	sum := md5.Sum([]byte(secret + word))
	encrypted := fmt.Sprintf("%x", sum)

	// write in response
	//w.Write([]byte(encrypted))

	// redirect to encrypted
	redirect := "/encrypted/" + encrypted
	http.Redirect(w, r, redirect, http.StatusMovedPermanently)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

}

func main() {
	var port = "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, application is running.")
	})

	router.HandleFunc("/encrypt/{secret}/{word}", encryptoHandler)

	router.HandleFunc("/encrypted/{encrypted}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Encryption done! Check address bar.")
	})

	http.ListenAndServe(":"+port, router)
	fmt.Printf("Server is active and listening on port %s\n", port)
}
