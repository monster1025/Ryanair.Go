// Ryanair.Go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
)

func main() {
	fromAirport := "RIX"

	url := "https://www.ryanair.com/en/api/2/routes/" + fromAirport + "/"
	request := gorequest.New()
	_, body, err := request.Get(url).End()
	if err != nil {
		panic(err)
	}
	flightDirections := make([]FlightDirection, 0)
	if err := json.Unmarshal([]byte(body), &flightDirections); err != nil {
		panic(err)
	}

	for _, flightDirection := range flightDirections {
		now := time.Now()
		fromDate := formatDate(now)
		toDate := formatDate(now.AddDate(0, 6, 0))

		url := fmt.Sprintf("https://www.ryanair.com/en/api/2/flights/from/%s/to/%s/%s/%s/outbound/cheapest-per-day/", fromAirport, flightDirection.AirportTo, fromDate, toDate)
		//fmt.Println(url)
		request := gorequest.New()
		_, body, err := request.Get(url).End()
		if err != nil {
			panic(err)
		}
		flights := FlightsResponse{}
		if err := json.Unmarshal([]byte(body), &flights); err != nil {
			panic(err)
		}
		flights.FromAirport = fromAirport
		flights.ToAirport = flightDirection.AirportTo

		for _, flight := range flights.Flights {
			if flight.Price.Value == 0 {
				continue
			}
			fmt.Println(fmt.Sprintf("Flight from %s to %s at %s costs %.6f", flights.FromAirport, flights.ToAirport, formatDate(flight.DateFrom), flight.Price.Value))
		}
	}
}
