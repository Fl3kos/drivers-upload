package responses

// ADriverResponse
//
// # Array de drivers de respuesta
//
// swagger:model DriverResponseArray
type ADriverResponse struct {
	DResponse []DriverResponse `json:"drivers"`
}

// DriverResponse
//
// # elemento de respuesta de drivers
//
// swagger:model DriverResponse
type DriverResponse struct {
	Name     string `json:"name"`
	Dni      string `json:"dni"`
	Phone    string `json:"phonenumber"`
	Username string `json:"username"`
	Password string `json:"password"`
}
