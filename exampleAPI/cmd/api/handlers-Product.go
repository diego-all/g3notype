package main

import (
	"net/http"
	"strconv"
	"time"

	models "github.com/diego-all/products-API/internal"

	"github.com/go-chi/chi"
)

type productRequest struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Price	int	`json:"price"`
	Quantity	int	`json:"quantity"`
}

type productResponse struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Price	int	`json:"price"`
	Quantity	int	`json:"quantity"`
}


func (app *application) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var productReq productRequest
	var payload jsonResponse

	err := app.readJSON(w, r, &productReq)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}


	var product = models.Product{
	Name:	productReq.Name,
	Description:	productReq.Description,
	Price:	productReq.Price,
	Quantity:	productReq.Quantity,
}

	_, err = app.models.Product.Insert(product)
	if err != nil {
		app.errorJSON(w, err)
		return
	}


	payload = jsonResponse{
	    Error:   false,
	    Message: "Product successfully created",
	    Data:    envelope{"product": product.Name},
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *application) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	product, err := app.models.Product.GetOneById(productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var productResponse = productResponse{
	Name:	product.Name,
	Description:	product.Description,
	Price:	product.Price,
	Quantity:	product.Quantity,
}

	payload := jsonResponse{
		Error:   false,
		Message: "Product successfully obtained",
		Data:    envelope{"product": productResponse},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}


func (app *application) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var productReq productRequest
	var payload jsonResponse

	err := app.readJSON(w, r, &productReq)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_, err = app.models.Product.GetOneById(productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var product = models.Product{
	Name:	productReq.Name,
	Description:	productReq.Description,
	Price:	productReq.Price,
	Quantity:	productReq.Quantity,
	UpdatedAt:   time.Now(),
 	Id:          productID,
}

	_, err = app.models.Product.Update(product)
	if err != nil {
		app.errorJSON(w, err)
		return
	}


	payload = jsonResponse{
	    Error:   false,
	    Message: "Product successfully updated",
	    Data:    envelope{"product": product.Name},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}


func (app *application) AllProducts(w http.ResponseWriter, r *http.Request) {
	var products models.Product
	all, err := products.GetAll() 
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Data successfully obtained",
		Data:    envelope{"products": all},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.Product.DeleteByID(productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Product successfully deleted",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}