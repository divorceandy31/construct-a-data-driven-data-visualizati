package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type DataPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Dataset struct {
	Name    string     `json:"name"`
	Data    []DataPoint `json:"data"`
	Options map[string]string `json:"options"`
}

type Analysis struct {
	Dataset Dataset `json:"dataset"`
	Results []Result `json:"results"`
}

type Result struct {
	Type    string  `json:"type"`
	Value   float64 `json:"value"`
	Message string  `json:"message"`
}

type API struct {
	log *zerolog.Logger
}

func NewAPI(log *zerolog.Logger) *API {
	return &API{log: log}
}

func (a *API) InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/datasets", a.getDatasets).Methods("GET")
	router.HandleFunc("/datasets/{id}", a.getDataset).Methods("GET")
	router.HandleFunc("/analyze", a.analyzeDataset).Methods("POST")
}

func (a *API) getDatasets(w http.ResponseWriter, r *http.Request) {
	// implement getting all datasets
}

func (a *API) getDataset(w http.ResponseWriter, r *http.Request) {
	// implement getting a single dataset
}

func (a *API) analyzeDataset(w http.ResponseWriter, r *http.Request) {
	var dataset Dataset
	err := json.NewDecoder(r.Body).Decode(&dataset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results := a.analyze(dataset)
	json.NewEncoder(w).Encode(Analysis{dataset, results})
}

func (a *API) analyze(dataset Dataset) []Result {
	// implement data analysis logic here
	results := make([]Result, 0)
	// ...
	return results
}

func main() {
	log := zerolog.NewConsoleWriter()
	api := NewAPI(log)
	router := mux.NewRouter()
	api.InitializeRoutes(router)
	http.ListenAndServe(":8080", router)
}