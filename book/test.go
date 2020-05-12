package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type world struct {
	radius float64
}
type location struct {
	name string  `json:"name"`
	lat  float64 `json:"latitude"`
	long float64 `json:"longitude"`
}

func (l location) description() string {
	return fmt.Sprint("%v()%.lf,%.lf", l.name, l.lat, l.long)
}

type gps struct {
	world       world
	current     location
	destination location
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}
func (g gps) message() string {
	return fmt.Sprint("%.lf km to %v", g.distance(), g.destination.description())
}
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
}
func main() {
	type location struct {
		Name string  `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	//curiosity := location{-4.5895, 137.4417}
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734}}
	bytes, err := json.MarshalIndent(locations, "", "")
	exitOnError(err)

	fmt.Println(string(bytes))
}

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
