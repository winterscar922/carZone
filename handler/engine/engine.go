package engine

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/winterscar922/carZone/models"
	"github.com/winterscar922/carZone/service"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{service: service}
}

func (h *EngineHandler) GetEngineById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	res, err := h.service.GetEngineById(ctx, int64(id))
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

func (h *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var engineReq models.EngineRequest
	err := json.NewDecoder(r.Body).Decode(&engineReq)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := models.ValidateEngineRequest(engineReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateEngine(ctx, engineReq)
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

func (h *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var engineReq models.EngineRequest
	err = json.NewDecoder(r.Body).Decode(&engineReq)

	if err != nil {
		log.Fatal("error while decoding request body")
		return
	}

	if err := models.ValidateEngineRequest(engineReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateEngine(ctx, engineReq, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteEngine(ctx, int64(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
