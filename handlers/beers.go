package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"mvc_layered/exchange"
	"mvc_layered/models"
	"net/http"
	"strconv"
)

//Get all beers
func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Header().Set("Content-Type", "application/josn")
	//cellar := models.DB.FindBeers()
	cellar := exchange.DB.FindBeers()
	json.NewEncoder(w).Encode(cellar)
}

// GetBeer returns a beer from the cellar
func GetBeer(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil{
		http.Error(w, fmt.Sprintf("%s is not a vaild Beer ID, it must be a number.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	//cellar, _ := models.DB.FindBeer(models.Beer{ID:ID})
	cellar, _ := exchange.DB.FindBeer(models.Beer{ID:ID})
	if len(cellar) == 1{
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cellar[0])
		return
	}
	http.Error(w, "The Beer you requestd does not exist.", http.StatusBadRequest)
}

//AddBeer adds a new beer to the cellar
func AddBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	decoder := json.NewDecoder(r.Body)

	var newBeer models.Beer
	err := decoder.Decode(newBeer)

	if err != nil{
		http.Error(w, "Bad beer - this will be a HTTP status code soon!", http.StatusBadRequest)
		return
	}

	exchange.DB.SaveBeer(newBeer)
	//models.DB.SaveBeer(newBeer)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New beer added.")
}
