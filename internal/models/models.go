package models

import "time"

type Client struct{
	Count int 
	WindowStart time.Time
}

type ErrorResponse struct{
	Status int `json:"status"`
	Message string `json:"message"`
}