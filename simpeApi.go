package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type People []Person

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Simple Rest Go Api .")
}

func listPage(w http.ResponseWriter, r *http.Request) {
	people := People{
		Person{Name: "Enes", Surname: "Ays", Age: 55},
		Person{Name: "Emir", Surname: "ER", Age: 55},
		Person{Name: "Cem", Surname: "Yar", Age: 55},
	}
	fmt.Print("Simple Rest Go Api list of persons .")
	json.NewEncoder(w).Encode(people)

}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/people", listPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
