package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// fungsi newRouter untuk melempar ke root
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	//fungsi route untuk menggunakan staticfile
	staticFileDirectory := http.Dir("./portofolio/")
	staticFileHandler := http.StripPrefix("/portofolio/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/portofolio/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	// memanggil fungsi konstruktor `newRouter` yang telah definisikan di atas
	// membuat server berjalan pada port 8080
	r := newRouter()
	log.Print("Listening on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

// fungsi untuk membuat tulisan "helloworld" yang akan di tampilkan pada route ("/")
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

//Buat route portofolio ke root server("/")
// func main() {
// 	fs := http.FileServer(http.Dir("./portofolio"))
// 	http.Handle("/", fs)

// 	log.Print("Listening on :3000...")
// 	err := http.ListenAndServe(":3000", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
