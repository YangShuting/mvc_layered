//使用 JSON 的方式来存储 beers 和 reviews 的值

package storage

import (
	"encoding/json"
	"fmt"
	"github.com/nanobox-io/golang-scribble"
	"mvc_layered/models"
	"strconv"
)

//上面的那个 golang-scribble 包是干嘛的？是golang一个JSON 的数据库：A tiny Golang JSON database
// JSON  是一个使用 JSON file 来存储数据的 数据层存储器：
type JSON struct{
	db *scribble.Driver
}

const (
	CollectionBeer int = iota
	CollectionReview
)

//使用 golang-scribble JSO你数据库来创建一个JSON
func NewJSON(location string) (*JSON, error){
	var err error

	stg := new(JSON)
	//疑问： new(JSON) 的作用, 看这篇文章，就是分配了一个没有内存的空间，但是返回了一个地址指针
    // https://stackoverflow.com/questions/34543430/golang-basics-struct-and-new-keyword
    //相当于： stg := &JSON{}
    stg.db, err = scribble.New(location, nil)
    if err != nil{
    	return nil, err
	}
    return stg, nil
}

//接下来的是， 6个CRUD函数
func (s *JSON) SaveBeer(beers ...models.Beer) error{
	for _, beer := range beers{
		var resource = strconv.Itoa(beer.ID)
		var collection = strconv.Itoa(CollectionBeer)
		allBeers := s.FindBeers()
		for _, b := range allBeers{
			if beer.Abv == b.Abv &&
				beer.Brewery == b.Brewery &&
				beer.Name == b.Name{
					return fmt.Errorf("Beer already exists")
			}
		}
		beer.ID = len(allBeers) + 1
		if err := s.db.Write(collection, resource,  &beer); err!= nil{
			return err
		}
	}
	return nil
}

//SaveReviews insert or update reviews
func(s *JSON) SaveReviews(reviews []*models.Review) error{
	for _, review := range reviews{
		var resource = strconv.Itoa(review.ID)
		var collection = strconv.Itoa(CollectionReview)

		allReviews := s.FindReviews()
		for _, r := range allReviews{
			if review.BeerID == r.BeerID &&
				review.FirstName == r.FirstName &&
				review.LastName == r.LastName &&
				review.Text == r.Text{
					return fmt.Errorf("Review already exists")
			}
		}
		review.ID = len(allReviews) + 1
		if err := s.db.Write(collection, resource, &review);err!=nil{
			return err
		}
	}
	return nil
}

//FindBeers find all beers
func (s *JSON)FindBeers() []*models.Beer{
	var beers []*models.Beer
	var collection = strconv.Itoa(CollectionBeer)

	records, err := s.db.ReadAll(collection)
	if err != nil{
		return beers
	}
	for _, b := range records{
		var beer models.Beer

		if err := json.Unmarshal([]byte(b), &beer); err!= nil{
			return beers
		}
		beers = append(beers, &beer)
	}
	return beers
}

// FindReviews find all reviews
func (s *JSON) FindReviews() []*models.Review{
	var reviews []*models.Review
	var collection = strconv.Itoa(CollectionReview)

	records, err := s.db.ReadAll(collection)
	if err != nil{
		return reviews
	}

	for _, r := range records{
		var review models.Review
		if err := json.Unmarshal([]byte(r), &review);err!=nil{
			return reviews
		}
		reviews = append(reviews, &review)
	}
	return reviews
}

// FindBeer Find a beer
func (s *JSON) FindBeer(criteria models.Beer) ([]*models.Beer, error){
	var beers []*models.Beer
	var beer models.Beer
	var resource = strconv.Itoa(criteria.ID)
	var collection = strconv.Itoa(CollectionBeer)

	if err := s.db.Read(collection, resource, &beer); err!=nil{
		return beers, nil
	}
	beers = append(beers, &beer)
	return beers, nil
}

//FindReview find a review
func (s *JSON) FindReview(criteria models.Review) ([]*models.Review, error){
	var reviews []*models.Review
	var review models.Review
	var resource = strconv.Itoa(criteria.ID)
	var collection = strconv.Itoa(CollectionReview)

	if err := s.db.Read(collection, resource, &review);err!=nil{
		return reviews, nil
	}

	reviews = append(reviews, &review)
	return reviews, nil
}


