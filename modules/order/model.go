package order

type Order struct {
	Id       int    `db:"id" json:"id"`
	Status   string `db:"status" json:"status"`
	DriverId int    `db:"driver_id" json:"driver_id"`
}
