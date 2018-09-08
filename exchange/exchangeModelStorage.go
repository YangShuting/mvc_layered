package exchange


//StorageType defines available storage types
import (
	"mvc_layered/models"
	"mvc_layered/storage"
)
type StorageType int

const (
	//JSON will store data in JSOn files saved on disk
	JSON StorageType = iota

	// Memory will store data in memory
	Memory
)

//DB is an interface to interact with data on multiple layered of data storage

var DB Storage

//Storage represents all possible actions available to deal with data
type Storage interface{
	SaveBeer(...models.Beer) error
	SaveReview(...models.Review) error
	FindBeer(models.Beer)([]*models.Beer, error)
	FindReview(models.Review)([]*models.Review, error)
	FindBeers() []models.Beer
	FindReviews() []models.Review
}

func NewStorage(storageType StorageType) error{
	switch storageType{
	case Memory:
		DB = new(storage.Memory)

	case JSON:
		_, err := storage.NewJSON("./data/")
		if err != nil{
			return err
		}
	}
	return nil
}