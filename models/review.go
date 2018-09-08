package models

import "time"

type Review struct{
	ID int `json:"id"`
	BeerID int `json:"beer_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Score float32 `json:"score"`
	Text string `json:"text"`
	Created time.Time `json:"created"`
}