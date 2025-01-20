package models

import "time"

//this is my model class for each real entity like here class and booking

type Class struct {
	ClassName       string
	Start_Date time.Time
	End_Date   time.Time
	Capacity   int64
}

type Booking struct{
	MemberName string
	BookingDate time.Time
}