package responses

// ADriverResponse
//
// # Respuesta de los drivers
//
// swagger:model DriverResponseArray
type ADriverResponse struct {
	AResponse []DriverResponse `json:"drivers"`
}

// DriverResponse
//
// # Respuesta con los elementos de un Driver
//
// swagger:model DriverResponse
type DriverResponse struct {
	Name     string `json:"name"`
	Dni      string `json:"dni"`
	Phone    string `json:"phonenumber"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// DriverErrorResponse
//
// swagger:model DriverErrorResponse
type DriverErrorResponse struct {
	Response string `json:"response"`
}
