package handlers

// WarehouseSwagger
// Objeto de warehouseUsers
// swagger:parameters AclPost
type WarehouseSwagger struct {
	//in: body
	Users WarehouseUsers `json:"users"`
}

// WarehouseUsers
// Elementos para publiccar un usuario de un warehouse
// swagger:model WarehouseUsers
type WarehouseUsers struct {
	// example: 5
	Pkr int `json:"pkr"`
	// example: 3
	Crd int `json:"crd"`
	// example: 1
	Adm int `json:"adm"`
	// example: 666777888
	WarehousePhone string `json:"phone"`
}
