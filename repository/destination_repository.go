package repository

import (
	"app-destination/model"
	"database/sql"
)

type DestinationRepository interface {
	GetAllDestinations() ([]model.Destination, error)
	InsertDestination(d model.Destination) error
}

type destinationRepo struct {
	database *sql.DB
}

//constructor
func NewDestinationRepository(database *sql.DB) DestinationRepository {
	return &destinationRepo{database: database}
}

func (r *destinationRepo) GetAllDestinations() ([]model.Destination, error) {
	rows, err := r.database.Query(`SELECT id, country, destination, description, price, tour_duration, rating, images, people, flag_image, price_discount 
		FROM destinations`)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var destinations []model.Destination

	for rows.Next(){
		var d model.Destination
		err := rows.Scan(
			&d.ID, &d.Country, &d.Destination, &d.Description,
			&d.Price, &d.TourDuration, &d.Rating,
			&d.Images, &d.People, &d.FlagImage, &d.PriceDiscount,
		)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, d)
	}
	return destinations, nil
}

func (r *destinationRepo) InsertDestination(d model.Destination) error {
	query := `INSERT INTO destinations (country, destination, description, price, tour_duration, rating, images, people, flag_image, price_discount)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.database.Exec(query,
		d.Country,
		d.Destination,
		d.Description,
		d.Price,
		d.TourDuration,
		d.Rating,
		d.Images,
		d.People,
		d.FlagImage,
		d.PriceDiscount,)

	if err != nil {
		return err
	}
	return nil

}
