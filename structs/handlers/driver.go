package handlers

// Drivers
//
// # Esta es la estructura usada para responder con frases
//
// swagger:model Drivers
// swagger:parameters Drivers
type Drivers struct {
	DriverA []Driver `json:"drivers"`
}

// Driver
//
// # Esta es la estructura usada para responder con frases
//
// swagger:model Driver
type Driver struct {
	Name        string `json:"name"`
	Dni         string `json:"dni"`
	PhoneNumber string `json:"phoneNumber"`
}
