//通过 本地 memeory 的方式来存储beers 和 reviws 的值

package storage

import "mvc_layered/models"

type Memory struct{
	cellar []models.Beer
	reviews []models.Review
}

//在这里不是很理解这个 []*models.Beer， 我知道单个beer是 models.Beer， 然后这个 []*models.Beer 是复数吗？应该不是
func (s *Memory) FindBeer(beer models.Beer) ([]*models.Beer, error){
	var beers []*models.Beer
	for idx := range s.cellar{
		if s.cellar[idx].ID == beer.ID{
			beers = append(beers, &s.cellar[idx])
		}
	}
	return beers, nil
}

//FindReview locate full data set based on given criteria
func (s *Memory) FindReview(review models.Review) ([]*models.Review, error){
	var reviews []*models.Review
	for idx := range s.reviews{
		if s.reviews[idx].BeerID == review.BeerID || s.reviews[idx].ID == review.ID{
			reviews = append(reviews, &s.reviews[idx])
		}
	}
	return reviews, nil
}

// FindBeers reutrn all beers
func (s *Memory) FindBeers() []models.Beer{
	return s.cellar
}

//FindReviews return all reviews
func (s *Memory) FindReviews() []models.Review{
	return s.reviews
}

// SaveBeer insert or update beers
func (s *Memory) SaveBeer(beers ...models.Beer) error{
	for _, beer := range beers{
		var err error
		beersFound, err := s.FindBeer(beer)
		if err != nil{
			return err
		}
		if len(beersFound) == 1{
			*beersFound[0] = beer
			return nil
		}

		beer.ID = len(s.cellar) + 1
		s.cellar = append(s.cellar, beer)
	}

	return nil
}

//SaveReview insert or update reviews
func (s *Memory) SaveReview(reviews ...models.Review) error{
	for _, review := range reviews{

		var err error
		reviewsFound, err := s.FindReview(review)
		if err != nil{
			return err
		}
		if len(reviewsFound) == 1{
			*reviewsFound[0] = review
		}
		review.ID = len(s.reviews) + 1
		s.reviews = append(s.reviews, review)
	}
	return nil
}