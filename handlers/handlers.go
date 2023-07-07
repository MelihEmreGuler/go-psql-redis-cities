package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/MelihEmreGuler/go-psql-redis-cities/entity"
	"github.com/MelihEmreGuler/go-psql-redis-cities/repository"
	"io"
	"net/http"
	"strconv"
)

// GetCity returns all cities
func GetCity(writer http.ResponseWriter, request *http.Request) {

	var getByID = request.URL.Query().Has("id")
	var getByName = request.URL.Query().Has("name")

	// Get city by id or name from database
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
	if getByName == true {
		cityName := request.URL.Query().Get("name")
		city := repository.CityRepo.GetByName(cityName)
		if city == nil {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte("City not found"))
			return
		}
		cityBytes, _ := json.Marshal(city)
		writer.Write(cityBytes)
		return
	}

	// Get city list from database
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

// PutCity updates a city
func PutCity(writer http.ResponseWriter, request *http.Request) {
	var city entity.City

	var putByID = request.URL.Query().Has("id")
	var putByName = request.URL.Query().Has("name")

	//check if id or name is given
	if putByID {
		cityIdStr := request.URL.Query().Get("id")
		cityId, err := strconv.Atoi(cityIdStr)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		city.Id = cityId
	}
	if putByName {
		cityName := request.URL.Query().Get("name")
		city.Id = repository.CityRepo.GetByName(cityName).Id
	}

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

	// Update city in database
	repository.CityRepo.Update(city)

	// Write response header with status code 200 (OK)
	writer.WriteHeader(http.StatusOK)
}

// DeleteCity deletes a city
func DeleteCity(writer http.ResponseWriter, request *http.Request) {
	var city entity.City

	var deleteByID = request.URL.Query().Has("id")
	var deleteByName = request.URL.Query().Has("name")

	//check if id or name is given
	if deleteByID {
		cityIdStr := request.URL.Query().Get("id")
		cityId, err := strconv.Atoi(cityIdStr)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		city.Id = cityId
	}
	if deleteByName {
		cityName := request.URL.Query().Get("name")
		city.Id = repository.CityRepo.GetByName(cityName).Id
		fmt.Println("cityId:", city.Id)
	}

	// Delete city in database
	repository.CityRepo.Delete(city)

	// Write response header with status code 200 (OK)
	writer.WriteHeader(http.StatusOK)
}
