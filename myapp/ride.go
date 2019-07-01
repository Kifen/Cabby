package myapp

import (
	"fmt"
	"strings"
	"time"

	//Go package implementing the haversine formula (used to get)
	//the great-circle distance between two points
	"github.com/umahmood/haversine"
)

//Struct for ride-hailing service Cabby
//Cabby implements the Cab interface
type Cabby struct {
	startTime, endTime        time.Time
	tFare, tip                float64
	pickUpPoint, dropOffPoint string
	destinations              []string
	fare                      Fare
	locationMap               map[string]map[string]float64
}

//method initializes Cabby properties
func (c *Cabby) InitValues() {
	c.fare.SetCostPerMile(200)
	c.fare.SetBaseFee(500)
	c.fare.SetBookingFee(250)

	c.SetDestinations([]string{"Choba", "Rumuosi", "Mgbuoba", "Alakahia", "Aluu", "Rumokoro", "Rumuola"})
	locationMap := map[string]map[string]float64{
		"choba": {
			"Lat": 4.8941,
			"Lon": 6.9263,
		},
		"rumuosi": {
			"Lat": 4.8807,
			"Lon": 6.9404,
		},
		"mgbuoba": {
			"Lat": 4.8421,
			"Lon": 6.9692,
		},
		"alakahia": {
			"Lat": 4.8851,
			"Lon": 6.9249,
		},
		"aluu": {
			"Lat": 4.9339,
			"Lon": 6.9437,
		},
		"rumokoro": {
			"Lat": 4.8703,
			"Lon": 6.9880,
		},
		"rumuola": {
			"Lat": 4.8356,
			"Lon": 7.0256,
		},
	}
	c.SetLocMap(locationMap)
}

//method to check if a pickUpPoint or dropOffPoint is valid
func (c Cabby) DestinationIsValid(dest string) bool {
	for i := 0; i < len(c.destinations); i++ {
		if strings.ToLower(c.destinations[i]) == strings.ToLower(dest) {
			return true
			//break
		}
	}
	return false
}

/*
#CalculateFare method simply calculates the price of a ride
it gets the latitude and longitude of both the pick up and drop off points
with the haversine formula it calculates the distance between the two points
and with the #Fare sruct as a model it calculates the price of a ride
*/
func (c Cabby) CalculateFare(pickUpPoint, dropOffPoint string) float64 {
	pickUpPointLat := c.locationMap[pickUpPoint]["Lat"]
	pickUpPointLon := c.locationMap[pickUpPoint]["Lon"]

	dropOffPointLat := c.locationMap[dropOffPoint]["Lat"]
	dropOffPointLon := c.locationMap[dropOffPoint]["Lon"]

	pop := haversine.Coord{Lat: pickUpPointLat, Lon: pickUpPointLon}
	dop := haversine.Coord{Lat: dropOffPointLat, Lon: dropOffPointLon}
	mi, _ := haversine.Distance(pop, dop)
	fare := c.calc(mi)
	return fare
}

func (c Cabby) calc(mile float64) float64 {
	costPerMile := c.fare.GetCostPerMile()
	baseFee := c.fare.GetBaseFee()
	bookingFee := c.fare.GetBookingFee()

	return (costPerMile * mile) + baseFee + bookingFee
}

//it checks if the amount paid by the user
//if the user paid less than the supposed fare it returns -1
//if the user paid the exact fare it returns 0
//if the user paid more than the supposed fare it returns 1
func (c Cabby) CheckUserFare(userAmount float64) int {
	if fare := c.tFare; fare < userAmount {
		return 1
	} else if fare > userAmount {
		return -1
	} else {
		return 0
	}
}

//method that gives a summary of the ride
func (c Cabby) TripDetails() {
	fmt.Println("\n<<TRIP DETAILS>>\n")
	fmt.Printf("PickUpPoint: %v\nDropOffPoint: %v\nFare: ₦%.2f\nTip: ₦%.2f\nStartTime: %v\nEndTime: %v",
		c.pickUpPoint, c.dropOffPoint, c.tFare, c.tip, c.startTime, c.endTime)
}

/*
Setters and Getter methods to set and return an instance of Cabby fields
*/
func (c *Cabby) SetStartTime(startTime time.Time) {
	c.startTime = startTime
}

func (c Cabby) GetStartTime() time.Time {
	return c.startTime
}

func (c *Cabby) SetEndTime(endTime time.Time) {
	c.endTime = endTime
}

func (c Cabby) GetEndTime() time.Time {
	return c.endTime
}

func (c *Cabby) SetTfare(tFare float64) {
	c.tFare = tFare
}

func (c Cabby) GetTfare() float64 {
	return c.tFare
}

func (c *Cabby) SetDestinations(destination []string) {
	c.destinations = destination
}

func (c Cabby) GetDestinations() []string {
	return c.destinations
}

func (c *Cabby) SetPickUpPoint(pickUpPoint string) {
	c.pickUpPoint = pickUpPoint
}

func (c Cabby) GetPickUpPoint() string {
	return c.pickUpPoint
}

func (c *Cabby) SetDropOffPoint(dropOffPoint string) {
	c.dropOffPoint = dropOffPoint
}

func (c Cabby) GetDropOffPoint() string {
	return c.dropOffPoint
}

func (c *Cabby) SetTip(tip float64) {
	c.tip = tip
}

func (c Cabby) GetTip() float64 {
	return c.tip
}

func (c *Cabby) SetLocMap(locationMap map[string]map[string]float64) {
	c.locationMap = locationMap
}

func (c Cabby) GetLocMap() map[string]map[string]float64 {
	return c.locationMap
}
