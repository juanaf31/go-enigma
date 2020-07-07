package controllers

import (
	"encoding/json"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/usecases/warehouseusecase"
	"liveCodeAPI/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WHHandler struct {
	whUseCase warehouseusecase.WHUsecase
}

func WarehouseController(r *mux.Router, service warehouseusecase.WHUsecase) {

	whHandler := WHHandler{whUseCase: service}

	warehouse := r.PathPrefix("/warehouse").Subrouter()
	warehouse.HandleFunc("", whHandler.listWarehouses).Methods(http.MethodGet)
	warehouse.Use(middleware.TokenValidationMiddleware)
	warehouse.HandleFunc("/{id}", whHandler.warehouse).Methods(http.MethodGet)
	warehouse.HandleFunc("/add", whHandler.addWarehouse).Methods(http.MethodPost)
	warehouse.HandleFunc("/delete/{id}", whHandler.deleteWarehouse).Methods(http.MethodDelete)
	warehouse.HandleFunc("/update/{id}", whHandler.updateWarehouse).Methods(http.MethodPut)
	r.HandleFunc("/warehouseinfo", whHandler.info).Methods(http.MethodGet)
}

func (s *WHHandler) listWarehouses(w http.ResponseWriter, r *http.Request) {
	warehouse, err := s.whUseCase.GetWarehouses()
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = warehouse
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *WHHandler) warehouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	warehouse, err := s.whUseCase.GetWarehouseByID(id)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = warehouse
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *WHHandler) addWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouseRequest *models.Warehouse
	_ = json.NewDecoder(r.Body).Decode(&warehouseRequest)
	_, err := s.whUseCase.AddWarehouse(warehouseRequest)
	if err != nil {
		w.Write([]byte("Cannot Add Data"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = warehouseRequest
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *WHHandler) deleteWarehouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := s.whUseCase.DeleteWarehouse(id)
	if err != nil {
		w.Write([]byte("Delete Data Failed!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success Deleted Data"
	byteData, err := json.Marshal(response)
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *WHHandler) updateWarehouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data models.Warehouse
	_ = json.NewDecoder(r.Body).Decode(&data)

	category, err := s.whUseCase.UpdateWarehouse(id, &data)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Update Data Failed!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = category
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *WHHandler) info(w http.ResponseWriter, r *http.Request) {

	warehouse, err := s.whUseCase.GetWHInfo()

	if err != nil {
		log.Println(err)
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = warehouse
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}
