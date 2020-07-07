package controllers

import (
	"encoding/json"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/usecases/goodsusecase"
	"liveCodeAPI/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GoodsHandler struct {
	goodsUseCase goodsusecase.GoodsUsecase
}

func GoodsController(r *mux.Router, service goodsusecase.GoodsUsecase) {
	goodsHandler := GoodsHandler{goodsUseCase: service}

	goods := r.PathPrefix("/goods").Subrouter()
	goods.HandleFunc("", goodsHandler.listGoods).Methods(http.MethodGet)
	goods.Use(middleware.TokenValidationMiddleware)
	goods.HandleFunc("/{id}", goodsHandler.goods).Methods(http.MethodGet)
	goods.HandleFunc("/add", goodsHandler.addGoods).Methods(http.MethodPost)
	goods.HandleFunc("/delete/{id}", goodsHandler.deleteGoods).Methods(http.MethodDelete)
	goods.HandleFunc("/update/{id}", goodsHandler.updateGoods).Methods(http.MethodPut)
}

func (s *GoodsHandler) listGoods(w http.ResponseWriter, r *http.Request) {
	goods, err := s.goodsUseCase.GetGoods()
	if err != nil {
		log.Println(err)
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = goods
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *GoodsHandler) goods(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	goods, err := s.goodsUseCase.GetGoodsByID(id)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = goods
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *GoodsHandler) addGoods(w http.ResponseWriter, r *http.Request) {
	var goodsRequest *models.Goods
	_ = json.NewDecoder(r.Body).Decode(&goodsRequest)
	_, err := s.goodsUseCase.AddGoods(goodsRequest)
	if err != nil {
		w.Write([]byte("Cannot Add Data"))
	}
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = goodsRequest
	byteData, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Something Wrong on Marshalling Data"))
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(byteData)
}

func (s *GoodsHandler) deleteGoods(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := s.goodsUseCase.DeleteGoods(id)
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

func (s *GoodsHandler) updateGoods(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data models.Goods
	_ = json.NewDecoder(r.Body).Decode(&data)

	category, err := s.goodsUseCase.UpdateGoods(id, &data)
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
