package config

const (
	CustomerPaid     = "paid"
	CustomerOccupied = "occupied"
)

var CustomerHistoryStatusTransitionRules = map[string]map[string]bool{
	CustomerPaid: {
		CustomerOccupied: true,
	},
	CustomerOccupied: {
		CustomerPaid: true,
	},
}

var CustomerHistoryStatus = map[string]bool{
	CustomerPaid:     true,
	CustomerOccupied: true,
}
