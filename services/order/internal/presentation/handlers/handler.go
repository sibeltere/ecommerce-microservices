package handlers

import (
	"encoding/json"
	"net/http"
	"orderservice/internal/application/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderService(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	IdStr := variables["Id"]

	objID, err := primitive.ObjectIDFromHex(IdStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrder(objID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
