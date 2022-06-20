package handler

import (
	"assignment-dua/entity"
	"assignment-dua/helper"
	"assignment-dua/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	service service.ServicesInterface
}

func NewHandler(service service.ServicesInterface) *Handler {
	return &Handler{
		service: service,
	}
}

func (h Handler) CreateOrders(writer http.ResponseWriter, req *http.Request) {
	var order entity.Orders
	jsonDecoder := json.NewDecoder(req.Body)
	err := jsonDecoder.Decode(&order)

	if err != nil {
		return
	}

	newOrder, errCreate := h.service.Create(req.Context(), order)

	orderResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   &newOrder,
	}

	if errCreate != nil {
		return
	}

	jsonData, _ := json.Marshal(orderResponse)
	writer.Header().Add("Content-Type", "application/json")
	_, errWrite := writer.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (h Handler) GetAllData(writer http.ResponseWriter, req *http.Request) {
	getAllOrder, err := h.service.GetAll(req.Context())

	fmt.Println("cek-data-handler", getAllOrder)

	getAllOrderResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   &getAllOrder,
	}

	jsonData, _ := json.Marshal(&getAllOrderResponse)
	writer.Header().Add("Content-Type", "application/json")
	_, err = writer.Write(jsonData)
	if err != nil {
		return
	}
}

func (h Handler) DeleteData(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]

	err := h.service.Delete(req.Context(), id)

	if err != nil {
		writer.Write([]byte("error when deleted"))
	}
	writer.Write([]byte("success when deleted"))
}

func (h Handler) UpdateData(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	var orders entity.Orders

	jsonDecoder := json.NewDecoder(req.Body)
	err := jsonDecoder.Decode(&orders)

	if err != nil {
		return
	}

	order, err := h.service.Update(req.Context(), orders, id)

	updateOrderResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   &order,
	}

	jsonData, _ := json.Marshal(&updateOrderResponse)
	writer.Header().Add("Content-Type", "application/json")
	_, err = writer.Write(jsonData)
	if err != nil {
		return
	}
}
