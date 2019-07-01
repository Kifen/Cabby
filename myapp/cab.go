package myapp

//Generic interface for a ride-hailing service
type Cab interface {
	CalculateFare() int
	DestinationIsValid(dest string) bool
}
