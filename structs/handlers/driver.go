package handlers

type Drivers struct {
	DriverA []Driver `json:"drivers"`

}

type Driver struct {
	Name string `json:"name"`
	Dni  string `json:"dni"`
	PhoneNumber string `json:"phoneNumber"`
}