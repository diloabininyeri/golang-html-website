package main

import (
	"net/http"
	"test/router"
)

func static() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

}

func main() {

	static()
	http.HandleFunc("/", router.Anasayfa)
	http.HandleFunc("/home", router.Home)
	err := http.ListenAndServe(":8004", nil)
	if err != nil {
		panic(err)
	}
}
