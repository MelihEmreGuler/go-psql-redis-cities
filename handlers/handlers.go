package handlers

import (
	"encoding/json"
	"github.com/MelihEmreGuler/go-psql-redis-cities/entity"
	"github.com/MelihEmreGuler/go-psql-redis-cities/repository"
	"io"
	"net/http"
	"strconv"
)

// GetCity returns all cities
func GetCity(writer http.ResponseWriter, request *http.Request) {

	var getByID = request.URL.Query().Has("id")
	if getByID == true {
		cityIdStr := request.URL.Query().Get("id")
		cityId, err := strconv.Atoi(cityIdStr)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		city := repository.CityRepo.GetById(cityId)
		if city == nil {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("City not found"))
			return
		}
		cityBytes, _ := json.Marshal(city)
		writer.Write(cityBytes)
		return
	}

	// Get cityList from database
	cityList := repository.CityRepo.List()

	// Encode cityList to json and write to response writer
	err := json.NewEncoder(writer).Encode(cityList)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PostCity inserts a new city
func PostCity(writer http.ResponseWriter, request *http.Request) {
	var city entity.City

	// Read request body as bytes (bodyBytes)
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal bodyBytes to City struct
	err = json.Unmarshal(bodyBytes, &city)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert city to database
	repository.CityRepo.Insert(city)

	// Write response header with status code 201 (Created)
	writer.WriteHeader(http.StatusCreated)
}
