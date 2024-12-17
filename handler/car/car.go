package car

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/winterscar922/carZone/service"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{service: service}
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("went inside get car()")
	ctx := r.Context()
	vars := mux.Vars(r)
	fmt.Println("var id is -->")
	fmt.Println(vars["id"])
	id, err := strconv.Atoi(vars["id"])

	fmt.Println("id is ")
	fmt.Println(id)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	res, err := h.service.GetCarById(ctx, id)

	fmt.Println("res got is")
	fmt.Println(res)
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
