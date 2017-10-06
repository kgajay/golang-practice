package main


import (
	"fmt"
	//"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"encoding/json"
	"time"
)

type Todo struct {
	Name      string 			`json:"name"`
	Completed bool				`json:"completed"`
	Due       time.Time			`json:"due"`
}

type Todos []Todo

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//fmt.Fprintln(w, "Welcome")
	todos := Todos{
		Todo{
			Name: "Write presentation",
		},
		Todo{
			Name: "Host meetup",
		},
	}

	json.NewEncoder(w).Encode(todos)

}

func TodoIndex(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "ToDo Index")
}

func TodoShow(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	fmt.Println(vars)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "ToDo Show:", todoId)
}