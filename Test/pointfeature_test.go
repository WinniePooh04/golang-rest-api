package main

import (
	"bytes"
	"encoding/json"
	"go-rest-api/config"
	"go-rest-api/models"
	"go-rest-api/routes"
	"go-rest-api/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	if router == nil {
		panic("Router Initialization Failed")
	}

	// Load Env
	config.LoadEnv()

	routes.SetUpRouter(router)

	// Initialize DB (check for errors)
	services.InitDBAdapter()

	return router
}

// Test FTTx search POST /search (Get FTTx by filter)
func TestGetFTTx(t *testing.T) {
	r := SetUpRouter()
	// Mock Request body
	fttxBody := models.FTTxRequestBody{
		SWLat:               16.8660524224038,
		SWLng:               96.15272251393414,
		NELat:               16.870944747638752,
		NELng:               96.15958360459423,
		Zoom:                18,
		Type:                "CA2",
		LineString:          "",
		NodeName:            "",
		ShowSearchHierarchy: "false",
		ShowPolygonCoverage: "true",
		DeviceStatus:        "",
		PortStatus:          "",
		Tags:                "",
		ConnectionTypes:     "",
		BoxType:             "",
	}
	jsonValue, _ := json.Marshal(fttxBody)
	req, _ := http.NewRequest("POST", "/fttx/search", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
