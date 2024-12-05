package entity

type Item struct {
	Name        string `json:"name"`
	Size        string `json:"size"`
	UniqueID    string `json:"unique_id"`
	Quantity    int    `json:"quantity"`
	WarehouseID int    `json:"warehouse_id"`
}
