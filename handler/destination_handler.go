package handler

import (
	"app-destination/model"
	"app-destination/service"
	"app-destination/utils"
	"encoding/json"
	"net/http"
)

type DestinationHandler struct {
	Service *service.DestinationService
}

func NewDestinationHandler(s *service.DestinationService) *DestinationHandler {
	return &DestinationHandler{Service: s}
}

func (h *DestinationHandler) GetAllDestinations (w http.ResponseWriter, r *http.Request) {
	destinations, err := h.Service.GetAllDestinations()
	if err != nil {
		http.Error(w, "failed to fetch data", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status": true,
		"message": "succes get data",
		"data": destinations,
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *DestinationHandler) CreateDestination (w http.ResponseWriter, r *http.Request) {
	var d model.Destination
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		// http.Error(w, "Invalid input", http.StatusBadRequest)
		utils.RespondError(w, http.StatusBadRequest, "invalid input")
		return
	}

	err := h.Service.CreateDestination(d)
	if err != nil {
		// http.Error(w, "Failed to create destination", http.StatusInternalServerError)
		utils.RespondError(w, http.StatusBadRequest,"failed to create destination")
		return
	}

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"status": true,
	// 	"message": "destination created successfully",
	// })

	utils.RespondJson(w, http.StatusCreated, map[string]interface{}{
		"status": true,
		"message": "destination created succesfully",
	})

}

