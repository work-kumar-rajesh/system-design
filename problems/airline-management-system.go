package main

import (
	"sync"
	"time"
)

// File: aircraft.go
type Aircraft struct {
	TailNumber string
	Model      string
	TotalSeats int
}

func NewAircraft(tailNumber, model string, totalSeats int) *Aircraft {
	return &Aircraft{
		TailNumber: tailNumber,
		Model:      model,
		TotalSeats: totalSeats,
	}
}

// File: airline_management_system.go
type AirlineManagementSystem struct {
	flights          []*Flight
	aircrafts        []*Aircraft
	flightSearch     *FlightSearch
	bookingManager   *BookingManager
	paymentProcessor *PaymentProcessor
	mu               sync.RWMutex
}

func NewAirlineManagementSystem() *AirlineManagementSystem {
	system := &AirlineManagementSystem{
		flights:          make([]*Flight, 0),
		aircrafts:        make([]*Aircraft, 0),
		bookingManager:   GetBookingManager(),
		paymentProcessor: GetPaymentProcessor(),
	}
	system.flightSearch = NewFlightSearch(system.flights)
	return system
}

func (ams *AirlineManagementSystem) AddFlight(flight *Flight) {
	ams.mu.Lock()
	defer ams.mu.Unlock()
	ams.flights = append(ams.flights, flight)
}

func (ams *AirlineManagementSystem) AddAircraft(aircraft *Aircraft) {
	ams.mu.Lock()
	defer ams.mu.Unlock()
	ams.aircrafts = append(ams.aircrafts, aircraft)
}

func (ams *AirlineManagementSystem) SearchFlights(source, destination string, date time.Time) []*Flight {
	return ams.flightSearch.SearchFlights(source, destination, date)
}

// File: booking.go
type Booking struct {
	BookingID   string
	Flight      *Flight
	Passenger   *Passenger
	SeatNumber  int
	BookingTime time.Time
}

func NewBooking(bookingID string, flight *Flight, passenger *Passenger, seatNumber int) *Booking {
	return &Booking{
		BookingID:   bookingID,
		Flight:      flight,
		Passenger:   passenger,
		SeatNumber:  seatNumber,
		BookingTime: time.Now(),
	}
}

// File: booking_manager.go
type BookingManager struct {
	bookings map[string]*Booking
	mu       sync.RWMutex
}

var (
	bookingManagerInstance *BookingManager
	onceBookingManager     sync.Once
)

func GetBookingManager() *BookingManager {
	onceBookingManager.Do(func() {
		bookingManagerInstance = &BookingManager{
			bookings: make(map[string]*Booking),
		}
	})
	return bookingManagerInstance
}

func (bm *BookingManager) AddBooking(booking *Booking) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	bm.bookings[booking.BookingID] = booking
}

func (bm *BookingManager) GetBooking(bookingID string) *Booking {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	return bm.bookings[bookingID]
}

// File: flight.go
type Flight struct {
	FlightNumber string
	Source       string
	Destination  string
	Departure    time.Time
	Arrival      time.Time
	Aircraft     *Aircraft
	Seats        []*Seat
}

func NewFlight(flightNumber, source, destination string, departure, arrival time.Time, aircraft *Aircraft) *Flight {
	seats := make([]*Seat, aircraft.TotalSeats)
	for i := 0; i < aircraft.TotalSeats; i++ {
		seats[i] = &Seat{
			SeatNumber: i + 1,
			IsBooked:   false,
		}
	}
	return &Flight{
		FlightNumber: flightNumber,
		Source:       source,
		Destination:  destination,
		Departure:    departure,
		Arrival:      arrival,
		Aircraft:     aircraft,
		Seats:        seats,
	}
}

func (f *Flight) BookSeat(seatNumber int) bool {
	if seatNumber < 1 || seatNumber > len(f.Seats) {
		return false
	}
	if f.Seats[seatNumber-1].IsBooked {
		return false
	}
	f.Seats[seatNumber-1].IsBooked = true
	return true
}

// File: flight_search.go
type FlightSearch struct {
	flights []*Flight
}

func NewFlightSearch(flights []*Flight) *FlightSearch {
	return &FlightSearch{
		flights: flights,
	}
}

func (fs *FlightSearch) SearchFlights(source, destination string, date time.Time) []*Flight {
	results := make([]*Flight, 0)
	for _, flight := range fs.flights {
		if flight.Source == source && flight.Destination == destination &&
			flight.Departure.Year() == date.Year() &&
			flight.Departure.Month() == date.Month() &&
			flight.Departure.Day() == date.Day() {
			results = append(results, flight)
		}
	}
	return results
}

// File: passenger.go
type Passenger struct {
	PassengerID string
	Name        string
	Email       string
	Phone       string
}

func NewPassenger(passengerID, name, email, phone string) *Passenger {
	return &Passenger{
		PassengerID: passengerID,
		Name:        name,
		Email:       email,
		Phone:       phone,
	}
}

// File: payment.go
type Payment struct {
	PaymentID string
	Amount    float64
	Method    string
	Status    string
}

func NewPayment(paymentID string, amount float64, method, status string) *Payment {
	return &Payment{
		PaymentID: paymentID,
		Amount:    amount,
		Method:    method,
		Status:    status,
	}
}

// File: payment_processor.go
type PaymentProcessor struct {
	payments map[string]*Payment
	mu       sync.RWMutex
}

var (
	paymentProcessorInstance *PaymentProcessor
	oncePaymentProcessor     sync.Once
)

func GetPaymentProcessor() *PaymentProcessor {
	oncePaymentProcessor.Do(func() {
		paymentProcessorInstance = &PaymentProcessor{
			payments: make(map[string]*Payment),
		}
	})
	return paymentProcessorInstance
}

func (pp *PaymentProcessor) ProcessPayment(payment *Payment) {
	pp.mu.Lock()
	defer pp.mu.Unlock()
	pp.payments[payment.PaymentID] = payment
}

// File: seat.go
type Seat struct {
	SeatNumber int
	IsBooked   bool
}
