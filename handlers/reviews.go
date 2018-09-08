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

//对 review 的操作：
//查询一个Beer所有的review的操作
func GetBeerReviews(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil{
		http.Error(w, fmt.Sprintf("%s is not a valid beer review id.", ps.ByName("id")), http.StatusBadRequest)
		return
	}

	results, _ := exchange.DB.FindReview(models.Review{BeerID:ID})
	//results, _ := models.DB.FindReview(models.Review{BeerID:ID})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

//AddBeerReview adds a new review for a beer
func AddBeerReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil{
		http.Error(w, fmt.Sprintf("%s is not a valid beer id"), http.StatusBadRequest)
		return
	}


	var newReview models.Review
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(newReview);err!=nil{
		http.Error(w, "Failed to parse a review", http.StatusBadRequest)
		return
	}

	newReview.BeerID = ID

	if err := exchange.DB.SaveReview(newReview);err != nil{
		http.Error(w, fmt.Sprintf("Failed to save review."), http.StatusBadRequest)
		return
	}
	//if err := models.DB.SaveReview(newReview);err!=nil{
	//}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New review has been added.")
}



