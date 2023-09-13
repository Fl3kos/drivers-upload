package responses

type ADriverResponse struct {
	DResponse []DriverResponse `json:"drivers"`

}

type DriverResponse struct {
	Name string `json:"name"`
	Dni string `json:"dni"`
	Phone string `json:"phonenumber"`
	Username string `json:"username"`
	Password string `json:"password"`
}