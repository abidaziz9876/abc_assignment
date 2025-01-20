package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateClassesController(t *testing.T) {


	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name       string
		query      string
		statusCode int
		response   string
	}{
		{
			name:       "Missing class_name",
			query:      "start_date=2025-01-20&end_date=2025-01-25&capacity=30",
			statusCode: http.StatusBadRequest,
			response:   `{"error":"class_name is required"}`,
		},
		{
			name:       "Missing start_date",
			query:      "class_name=Math&end_date=2025-01-25&capacity=30",
			statusCode: http.StatusBadRequest,
			response:   `{"error":"start_date is required"}`,
		},
		{
			name:       "Invalid capacity format",
			query:      "class_name=Math&start_date=2025-01-20&end_date=2025-01-25&capacity=abc",
			statusCode: http.StatusBadRequest,
			response:   `{"error":"Invalid capacity format"}`,
		},
		{
			name:       "Successful creation",
			query:      "class_name=Math&start_date=2025-01-20&end_date=2025-01-25&capacity=30",
			statusCode: http.StatusCreated,
			response:   `{"message":"Class created successfully"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			
			router := gin.Default()
			router.GET("/create-class", CreateClassesController())

			// Creating a test request
			req := httptest.NewRequest(http.MethodGet, "/create-class?"+test.query, nil)
			rec := httptest.NewRecorder()

			// Performing the request
			router.ServeHTTP(rec, req)

			// Checking the status code
			assert.Equal(t, test.statusCode, rec.Code)

			// Checking the response body
			if test.statusCode == http.StatusCreated {
				// For successful cases, validating parts of the response
				var resp map[string]interface{}
				err := json.Unmarshal(rec.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, "Class created successfully", resp["message"])
				assert.NotNil(t, resp["class"])
			} else {
				assert.JSONEq(t, test.response, rec.Body.String())
			}
		})
	}
}


func TestCreateBooking(t *testing.T) {

	gin.SetMode(gin.TestMode)

	// Test cases
	tests := []struct {
		name       string
		query      string
		statusCode int
		response   string
	}{
		{
			name:       "Missing name",
			query:      "date=2025-01-20",
			statusCode: http.StatusBadRequest,
			response:   `{"error":"name is required"}`,
		},
		{
			name:       "Missing date",
			query:      "name=John",
			statusCode: http.StatusBadRequest,
			response:   `{"error":"date is required"}`,
		},
		{
			name:       "Invalid date format",
			query:      "name=John&date=invalid-date",
			statusCode: http.StatusBadRequest,
			response:   `{"error":"Invalid date format for date"}`, 
		},
		{
			name:       "Successful booking",
			query:      "name=John&date=2025-01-20",
			statusCode: http.StatusCreated,
			response:   `{"message":"Booking successful"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			
			router := gin.Default()
			router.GET("/create-booking", CreateBooking())

			// Creating a test request
			req := httptest.NewRequest(http.MethodGet, "/create-booking?"+test.query, nil)
			rec := httptest.NewRecorder()

			// Performing the request
			router.ServeHTTP(rec, req)

			// Checking the status code
			assert.Equal(t, test.statusCode, rec.Code)

			// Checking the response body
			if test.statusCode == http.StatusCreated {
				// For successful cases, validating parts of the response
				var resp map[string]interface{}
				err := json.Unmarshal(rec.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, "Booking successful", resp["message"])
				assert.NotNil(t, resp["booking"])
			} else {
				// For error cases, match the exact error message
				assert.JSONEq(t, test.response, rec.Body.String())
			}
		})
	}
}
