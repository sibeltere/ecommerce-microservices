package controllers

import (
	"ecommerce-microservices/services/product/internal/application/services"
	"ecommerce-microservices/services/product/internal/domain/models"
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

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product *models.CreateProductModel

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isSuccess, serviceerr := ph.service.CreateProduct(product)

	if !isSuccess || serviceerr != nil {
		http.Error(w, "product not created", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isSuccess)
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product *models.UpdateProductModel

	variables := mux.Vars(r)
	idStr := variables["id"]

	id, _ := strconv.Atoi(idStr)

	//firstly get a product Is it exits?
	existingProduct, err := ph.service.GetByID(id)

	if existingProduct == nil || err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	//now you can parse your request to product
	if er := json.NewDecoder(r.Body).Decode(&product); er != nil {
		http.Error(w, "product model is not readable", http.StatusBadRequest)
		return
	}

	isSuccess, serviceerr := ph.service.UpdateProduct(product, id)

	if !isSuccess || serviceerr != nil {
		http.Error(w, "product not updated", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isSuccess)

}

func (ph *ProductHandler) DeleteProuct(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	IdStr := variables["id"]
	Id, _ := strconv.Atoi(IdStr)

	//firstly get a product Is it exits?
	existingProduct, err := ph.service.GetByID(Id)

	if existingProduct == nil || err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	isSuccess, serviceerr := ph.service.DeleteProuct(Id)

	if !isSuccess || serviceerr != nil {
		http.Error(w, "product not deleted", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isSuccess)

}
