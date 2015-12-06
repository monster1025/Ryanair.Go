package main

import (
	"time"
)

// direction response
type FlightDirection struct {
	AirportFrom string `json:"airportFrom"`
	AirportTo   string `json:"airportTo"`
}

//response

type FlightsResponse struct {
	Flights     []Flight `json:"flights"`
	MinPrice    Price    `json:"minPrice"`
	MaxPrice    Price    `json:"maxPrice"`
	FromAirport string
	ToAirport   string
}

type Flight struct {
	Price      Price     `json:"price"`
	AdvertText Advertise `json:"advertText"`
	DateFrom   time.Time `json:"dateFrom"`
	DateTo     time.Time `json:"dateTo"`
}
type Price struct {
	MinPrice            string  `json:"minPrice"`
	Value               float32 `json:"value"`
	ValueMainUnit       string  `json:"valueMainUnit"`
	ValueFractionalUnit string  `json:"valueFractionalUnit"`
	CurrencySymbol      string  `json:"currencySymbol"`
}
type Advertise struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
