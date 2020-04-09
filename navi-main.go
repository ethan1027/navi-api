package main

import (
	"context"
	"log"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func main() {
	c, err := maps.NewClient(maps.WithAPIKey("put your apikey"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	
	r := &maps.DirectionsRequest{
		Origin:      "2502 Babcock Rd, San Antonio, TX 78229",
		Destination: "USAA Gate 6, San Antonio, TX",
		Mode: maps.TravelModeDriving,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	pretty.Println(route)
}
