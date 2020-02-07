package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "subasic/controllers"
)

var router *mux.Router
var userController = controllers.UserController{}
var groupController = controllers.GroupController{}

func main() {
   handleRequests()
}

func setupRouter() {
    router = mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homePage)
    setupUserMappings()
    setupGroupMappings()
}

func handleRequests() {
    setupRouter()
    http.ListenAndServe(":8000", router)
}

func setupUserMappings(){
    router.HandleFunc("/users", userController.GetUsers).Methods("GET")
    router.HandleFunc("/user", userController.CreateUser).Methods("POST")
    router.HandleFunc("/user/{id}", userController.GetUser).Methods("GET")
    router.HandleFunc("/user", userController.UpdateUser).Methods("PUT")
    router.HandleFunc("/user/{id}", userController.DeleteUser).Methods("DELETE")
}

func setupGroupMappings(){
    router.HandleFunc("/groups", groupController.GetGroups).Methods("GET")
    router.HandleFunc("/group", groupController.CreateGroup).Methods("POST")
    router.HandleFunc("/group/{id}", groupController.GetGroup).Methods("GET")
    router.HandleFunc("/group", groupController.UpdateGroup).Methods("PUT")
    router.HandleFunc("/group/{id}", groupController.DeleteGroup).Methods("DELETE")
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: homePage")
    fmt.Fprintf(w, "Welcome to the HomePage!")
}
