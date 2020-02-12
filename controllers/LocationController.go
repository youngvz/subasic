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

type LocationController struct {
	Controller
}

func (l LocationController) CreateLocation(w http.ResponseWriter, r * http.Request){
	fmt.Println("Endpoint Hit: create location")
	var location models.Location
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Could not parse body")
	}
	json.Unmarshal(reqBody, &location)

	locationDao := service.LocationService{}
	err = locationDao.Create(&location)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Could not create location")
		return
	}

	json.NewEncoder(w).Encode("Location created")
}

func (l LocationController) GetLocation(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: get location by id")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil { Controller{}.checkErr(err) }

	locationDao := service.LocationService{}
	location, err := locationDao.GetById(id)
	if err != nil {
		Controller{}.checkErr(err)
		json.NewEncoder(w).Encode("Could not retrieve location")
		return
	}
	json.NewEncoder(w).Encode(location)
}

func (l LocationController) GetLocations(w http.ResponseWriter, r *http.Request){
	locationDao := service.LocationService{}
	locations, err := locationDao.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode("Could not retrieve locations")
	}
	json.NewEncoder(w).Encode(locations)
}

func (l LocationController) UpdateLocation(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: update location")
	var location models.Location
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("Could not parse body")
	}
	json.Unmarshal(reqBody, &location)

	locationDao := service.LocationService{}
	err = locationDao.Update(&location)
	if err != nil {
		json.NewEncoder(w).Encode("Could not update location")
		return
	}
	json.NewEncoder(w).Encode("Location updated")
}

func (l LocationController) DeleteLocation(w http.ResponseWriter, r* http.Request){
	fmt.Println("Endpoint Hit: delete location by id")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil { Controller{}.checkErr(err) }

	locationDao := service.LocationService{}
	err = locationDao.Delete(id)
	if err != nil {
		Controller{}.checkErr(err)
		json.NewEncoder(w).Encode("Could not delete location")
		return
	}
	json.NewEncoder(w).Encode("Location deleted")
}