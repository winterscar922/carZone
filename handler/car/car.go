package car

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/winterscar922/carZone/models"
	"github.com/winterscar922/carZone/service"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	res, err := h.service.GetCarById(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var carReq models.CarRequest
	err := json.NewDecoder(r.Body).Decode(&carReq)

	if err != nil {
		log.Fatal("error while decoding request body")
		return
	}

	res, err := h.service.CreateCar(ctx, carReq)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
