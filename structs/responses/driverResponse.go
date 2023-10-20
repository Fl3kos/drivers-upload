package responses

// ADriverResponse
//
// # Esta es la estructura usada para responder con frases
//
// swagger:model DriverResponseArray
type ADriverResponse struct {
	DResponse []DriverResponse `json:"drivers"`
}

// DriverResponse
//
// # Esta es la estructura usada para responder con frases
//
// swagger:model DriverResponse
type DriverResponse struct {
	Name     string `json:"name"`
	Dni      string `json:"dni"`
	Phone    string `json:"phonenumber"`
	Username string `json:"username"`
	Password string `json:"password"`
}
