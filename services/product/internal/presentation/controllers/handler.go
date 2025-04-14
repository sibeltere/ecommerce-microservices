package controllers

import (
	"ecommerce-microservices/services/product/internal/application/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) (productHandler *ProductHandler) {
	return &ProductHandler{service: service}
}

func (ph *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	variables := mux.Vars(r)
	idStr := variables["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	product, serviceerr := ph.service.GetByID(id)
	if serviceerr != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (ph *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	productList, serviceerr := ph.service.GetAllProduct()
	if serviceerr != nil {
		http.Error(w, "there is no any product", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productList)
}
