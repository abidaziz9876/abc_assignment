package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"demo/models"
	"demo/services"

	"github.com/gin-gonic/gin"
)

// Helper function for parsing and validating dates
func parseDate(dateStr, fieldName string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("%s must be in the format YYYY-MM-DD", fieldName)
	}
	return date, nil
}

// CreateClassesController handles creating new classes
func CreateClassesController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Validate and parse inputs
		className := ctx.Query("class_name")
		if className == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "class_name is required"})
			return
		}

		startDateStr := ctx.Query("start_date")
		if startDateStr == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "start_date is required"})
			return
		}
		startDate, err := parseDate(startDateStr, "start_date")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		endDateStr := ctx.Query("end_date")
		if endDateStr == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "end_date is required"})
			return
		}
		endDate, err := parseDate(endDateStr, "end_date")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		capacityStr := ctx.Query("capacity")
		if capacityStr == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "capacity is required"})
			return
		}
		capacity, err := strconv.ParseInt(capacityStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid capacity format"})
			return
		}

		// Create the class object
		class := &models.Class{
			ClassName:  className,
			Start_Date: startDate,
			End_Date:   endDate,
			Capacity:   capacity,
		}

		// Call service to create the class
		services.CreateClass(class)

		ctx.JSON(http.StatusCreated, gin.H{"message": "Class created successfully", "class": class})
	}
}

// GetAllClassesController retrieves all classes
func GetAllClassesController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		classes, err := services.GetAllClasses()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": classes})
	}
}

// CreateBooking handles booking a class
func CreateBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Validate and parse inputs
		name := ctx.Query("name")
		if name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
			return
		}

		dateStr := ctx.Query("date")
		if dateStr == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "date is required"})
			return
		}
		date, err := parseDate(dateStr, "date")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for date"})
			return
		}

		// Create the booking object
		booking := &models.Booking{
			MemberName:  name,
			BookingDate: date,
		}

		// Call service to create the booking
		services.CreateBooking(booking)

		ctx.JSON(http.StatusCreated, gin.H{"message": "Booking successful", "booking": booking})
	}
}
