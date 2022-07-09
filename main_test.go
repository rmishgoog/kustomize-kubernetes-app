package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetCarsHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/cars", GetCarsHandler)
	req, _ := http.NewRequest("GET", "/cars", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var cars []Car
	json.Unmarshal(w.Body.Bytes(), &cars)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, cars)
}

func TestAddCarsHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/cars", AddCarsHandler)
	car_id := xid.New().String()
	car := Car{
		ID:     car_id,
		Name:   "Test Vehicle",
		Stock:  30,
		Price:  555.95,
		Origin: "USA",
		Engine: "v6",
	}
	jsonValue, _ := json.Marshal(car)
	req, _ := http.NewRequest("POST", "/cars", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
