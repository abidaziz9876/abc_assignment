package services

import (
	"demo/models"
	"errors"
	"fmt"
)

//service package where all my main business logic will be defined

var classList []models.Class //in-memory data-store 

func CreateClass(classObj *models.Class) {
	//there can be many checks and error which may occur so we have to handle it also
	classList = append(classList, *classObj)
}

func GetAllClasses() ([]models.Class, error) {

	if len(classList) == 0 {
		return nil, errors.New("no classes available")
	}

	return classList, nil
}

func CreateBooking(booking *models.Booking) {
	//there can be many checks and business logic we can apply here before booking for a class
	fmt.Println("Booking successfully")
}
