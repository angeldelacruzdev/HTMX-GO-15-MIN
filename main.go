package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id      int
	Message string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{
				Id:      1,
				Message: "Buy Milk",
			},
		},
	}

	todosHandle := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))

		templ.Execute(w, data)
	}

	addTodoHandle := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		templ := template.Must(template.ParseFiles("index.html"))

		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}

		data["Todos"] = append(data["Todos"], todo)

		templ.ExecuteTemplate(w, "todo-list-element", todo)

	}

	http.HandleFunc("/", todosHandle)

	http.HandleFunc("/api/add-todo", addTodoHandle)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
