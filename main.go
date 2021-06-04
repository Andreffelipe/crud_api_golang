package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func allUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"All users endpoint hit")
}

func newUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"New user Endpoint hit")
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Delete user endpoint hit")
}

func updateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Update user endpoint hit")
}

func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/",deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}",updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}",newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	fmt.Println("go tutorial")

	handleRequest()
}