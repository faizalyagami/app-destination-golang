package dto

import "app-destination/model"

type DestinationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []model.Destination `json:"data"`
}