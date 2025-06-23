package model

type Destination struct {
	ID            int     `json:"id"`
	Country       string  `json:"country"`
	Destination   string  `json:"destination"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	TourDuration  string  `json:"tour_duration"`
	Rating        int     `json:"rating"`
	Images        string  `json:"images"`
	People        string  `json:"people"`
	FlagImage     string  `json:"flag_image"`
	PriceDiscount float64 `json:"price_discount"`
}