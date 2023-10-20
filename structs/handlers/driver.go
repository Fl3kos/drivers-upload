package handlers

// Drivers
//
// # Driver parametros de entrada
//
// swagger:parameters DriverPost
type Drivers struct {
	//in: body
	DriverA []Driver `json:"drivers"`
}

// Driver
//
// # Elementos de un driver
//
// swagger:model Driver
type Driver struct {
	Name        string `json:"name"`
	Dni         string `json:"dni"`
	PhoneNumber string `json:"phoneNumber"`
}

// WarehouseCode
//
// # codigo del warehouse
//
// swagger:parameters DriverPost
type WarehouseCode struct {
	//in: query
	warehouseCode string `query:"warehouseCode"`
}
