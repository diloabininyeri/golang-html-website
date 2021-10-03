package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Surname string
	Address string
	Array   [4]string
	Data    TodoPageData
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Todo struct {
	Title string
	Done  bool
}

func anasayfa(w http.ResponseWriter, r *http.Request) {

	user := User{Name: "dılo", Surname: "sürücü"}

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	user.Data = data

	user.Array[0] = "bir"
	user.Array[1] = "iki"
	user.Array[2] = "uc"
	user.Array[3] = "dort"

	name := "DILO SURUCU"

	user.Address = name
	t, err := template.ParseFiles("sablon.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func main() {

	http.HandleFunc("/", anasayfa)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))

	err := http.ListenAndServe(":8004", nil)
	if err != nil {
		panic(err)
	}
}
