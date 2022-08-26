package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"price/internal/data"
)

type Handler interface {
	Register(router *mux.Router)
}

type handler struct {
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/", h.GetAll)
	router.HandleFunc("/msk", h.GetMSK)
	router.HandleFunc("/sh", h.GetSH)
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {

	filename := "priceMSK"
	filename2 := "productMSK"
	priceMSK := data.ImportXML(filename)
	productMSK := data.ImportXML(filename2)
	msk := data.CollectProduct(productMSK, priceMSK)

	filename3 := "priceSH"
	filename4 := "productSH"
	priceSH := data.ImportXML(filename3)
	productSH := data.ImportXML(filename4)
	sh := data.CollectProduct(productSH, priceSH)

	var collect data.Price
	collect.ProductMSK = msk
	collect.ProductSH = sh

	resultJson, err := json.MarshalIndent(collect, " ", " ")
	if err != nil {
		errors.New(fmt.Sprintf("не удалось перекодировать данные. ошибка: %v", err))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resultJson)
}

func (h *handler) GetMSK(w http.ResponseWriter, r *http.Request) {
	filename := "priceMSK"
	filename2 := "productMSK"
	priceMSK := data.ImportXML(filename)
	productMSK := data.ImportXML(filename2)
	price := data.CollectProduct(productMSK, priceMSK)
	var collect data.Price
	collect.ProductMSK = price

	resultJson, err := json.MarshalIndent(collect.ProductMSK, " ", " ")
	if err != nil {
		errors.New(fmt.Sprintf("не удалось перекодировать данные. ошибка: %v", err))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resultJson)
}

func (h *handler) GetSH(w http.ResponseWriter, r *http.Request) {
	filename := "priceSH"
	filename2 := "productSH"
	priceSH := data.ImportXML(filename)
	productSH := data.ImportXML(filename2)
	price := data.CollectProduct(productSH, priceSH)
	var collect data.Price
	collect.ProductSH = price

	resultJson, err := json.MarshalIndent(collect.ProductSH, " ", " ")
	if err != nil {
		errors.New(fmt.Sprintf("не удалось перекодировать данные. ошибка: %v", err))
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resultJson)
}
