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

type GroupController struct {
	Controller
}

func (g GroupController) CreateGroup(w http.ResponseWriter, r * http.Request){
	fmt.Println("Endpoint Hit: create group")
	var group models.Group
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Could not parse body")
	}
	json.Unmarshal(reqBody, &group)

	groupDao := service.GroupService{}
	err = groupDao.Create(&group)
	if err != nil {
		json.NewEncoder(w).Encode("Could not create group")
		return
	}
	json.NewEncoder(w).Encode("Group created")
}

func (g GroupController) GetGroup(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: get group by id")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil { Controller{}.checkErr(err) }

	groupDao := service.GroupService{}
	group, err := groupDao.GetById(id)
	if err != nil {
		json.NewEncoder(w).Encode("Could not create group")
		return
	}
	json.NewEncoder(w).Encode(group)
}

func (g GroupController) GetGroups(w http.ResponseWriter, r *http.Request){
	groupDao := service.GroupService{}
	groups, err := groupDao.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode("Could not retrieve users")
	}
	json.NewEncoder(w).Encode(groups)
}

func (g GroupController) UpdateGroup(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: update group")
	var group models.Group
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Could not parse body")
	}
	json.Unmarshal(reqBody, &group)

	groupDao := service.GroupService{}
	err = groupDao.Update(&group)
	if err != nil {
		json.NewEncoder(w).Encode("Could not update group")
		return
	}
	json.NewEncoder(w).Encode("Group updated")
}

func (g GroupController) DeleteGroup(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: delete group by id")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil { Controller{}.checkErr(err) }

	groupDao := service.GroupService{}
	err = groupDao.Delete(id)
	if err != nil {
		json.NewEncoder(w).Encode("Could not delete user")
		return
	}
	json.NewEncoder(w).Encode("Group deleted")
}