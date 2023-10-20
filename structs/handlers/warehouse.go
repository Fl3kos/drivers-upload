package handlers

// WarehouseUsers
//
// # elemento de respuesta de warehouse users
//
// swagger:model WarehouseUsers
type WarehouseUsers struct {
	Pkr            int    `json:"pkr"`
	Crd            int    `json:"crd"`
	Adm            int    `json:"adm"`
	WarehousePhone string `json:"phone"`
}
