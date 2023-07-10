package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/MelihEmreGuler/go-psql-redis-cities/brokers"
	"github.com/MelihEmreGuler/go-psql-redis-cities/cache"
	"github.com/MelihEmreGuler/go-psql-redis-cities/entity"
	"github.com/MelihEmreGuler/go-psql-redis-cities/repository"
	"io"
	"net/http"
	"strconv"
)

// rabbitmq broker instance for publish
var rabbitmq = brokers.NewRabbitMQ()

// Create a new redis client
var redisCache = cache.NewRedis()

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
	} else if getByName == true {
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

	//if the request is not for a specific city, we will write all cities

	// Get city list from redis cache
	cityCache := redisCache.Get()
	if cityCache != nil {
		fmt.Println("City list fetched from redis cache")
		writer.Write(cityCache)
		return
	}

	fmt.Println("City list is not in redis cache")

	// Get city list from database
	cityList := repository.CityRepo.List()
	cityListBytes, _ := json.Marshal(cityList)

	go func(data []byte) {
		fmt.Println("City is caching to redis cache")

		// Set city list to redis cache
		redisCache.Put(cityListBytes)

		// Publish message to rabbitmq
		rabbitmq.Publish([]byte("City list fetched"))
	}(cityListBytes)

	// Write response header with status code 200 (OK)
	writer.Write(cityListBytes)
	writer.WriteHeader(http.StatusOK)

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

	// Publish message to rabbitmq
	rabbitmq.Publish([]byte("City created with name: " + city.Name))

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
