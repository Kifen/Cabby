package myapp

/*
Struct for (transportation) fare
#costPerMile as the name implies is how much charged for each mile of a ride
#bookingFee is the cost associated with booking the ride for a customer/passenger
#baseFee is a flat fee charged at the beginning of every ride
Struct Fare has accompanying methods to set and get its properties
*/

type Fare struct {
	costPerMile, bookingFee, baseFee float64
}

func (f *Fare) SetCostPerMile(costPerMile float64) {
	f.costPerMile = costPerMile
}

func (f Fare) GetCostPerMile() float64 {
	return f.costPerMile
}

func (f *Fare) SetBookingFee(bookingFee float64) {
	f.bookingFee = bookingFee
}

func (f Fare) GetBookingFee() float64 {
	return f.bookingFee
}

func (f *Fare) SetBaseFee(baseFee float64) {
	f.baseFee = baseFee
}

func (f Fare) GetBaseFee() float64 {
	return f.baseFee
}
