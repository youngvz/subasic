package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"subasic/models"
	"subasic/service"
)

type UserController struct {
	Controller
}

func (u UserController) CreateUser(w http.ResponseWriter, r * http.Request){
	fmt.Println("Endpoint Hit: create user")
	var user models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Could not parse body")
	}
	json.Unmarshal(reqBody, &user)

	userDao := service.UserService{}
	err = userDao.Create(&user)
	if err != nil {
		json.NewEncoder(w).Encode("Could not create user")
		return
	}

	json.NewEncoder(w).Encode("User created")
}

func (u UserController) GetUser(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: get user by id")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil { Controller{}.checkErr(err) }

	userDao := service.UserService{}
	user, err := userDao.GetById(id)
	if err != nil {
		Controller{}.checkErr(err)
		json.NewEncoder(w).Encode("Could not retrieve user")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (u UserController) GetUsers(w http.ResponseWriter, r *http.Request){
	userDao := service.UserService{}
	users, err := userDao.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode("Could not retrieve users")
	}
	json.NewEncoder(w).Encode(users)
}

func (u UserController) UpdateUser(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: update user")
	var user models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Could not parse body")
	}
	json.Unmarshal(reqBody, &user)

	userDao := service.UserService{}
	err = userDao.Update(&user)
	if err != nil {
		json.NewEncoder(w).Encode("Could not update user")
		return
	}
	json.NewEncoder(w).Encode("User updated")
}

func (u UserController) DeleteUser(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: delete user by id")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil { Controller{}.checkErr(err) }

	userDao := service.UserService{}
	err = userDao.Delete(id)
	if err != nil {
		json.NewEncoder(w).Encode("Could not delete user")
		return
	}
	json.NewEncoder(w).Encode("User deleted")
}