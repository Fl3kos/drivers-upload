package handlers

// DriverSwagger
//
// # Driver parametros de entrada
//
// swagger:parameters DriverPost
type DriverSwagger struct {
	//in: body
	// required: true
	DriverA Drivers `json:"drivers"`
}

// Drivers
//
// # Array de Driver para su post en auth
//
// swagger:model Drivers
type Drivers struct {
	//in: body
	// required: true
	DriverA []Driver `json:"drivers"`
}

// Driver
//
// # Elementos de un driver
//
// swagger:model Driver
type Driver struct {
	// example: pepe
	Name string `json:"name"`
	// example: 12345678Z
	Dni string `json:"dni"`
	// example: 666777888
	PhoneNumber string `json:"phoneNumber"`
}

// warehouseCode
//
// # Codigo del warehouse
//
// swagger:parameters DriverPost AclPost
type WarehouseCode struct {
	// warehouse code to path
	//
	// in: path
	// required: true
	// example: 12345
	WarehouseCode string `path:"warehouseCode"`
}
