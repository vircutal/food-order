package config

const (
	CustomerPaid     = "paid"
	CustomerOccupied = "occupied"
)

var CustomerStatusTransitionRules = map[string]map[string]bool{
	CustomerPaid: {
		CustomerOccupied: true,
	},
	CustomerOccupied: {
		CustomerPaid: true,
	},
}

var CustomerStatus = map[string]bool{
	CustomerPaid:     true,
	CustomerOccupied: true,
}
