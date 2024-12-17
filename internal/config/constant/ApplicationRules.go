package constant

var CustomerHistoryStatusTransitionRules = map[string]map[string]bool{
	CustomerPaid: {
		CustomerOccupied: true,
	},
	CustomerOccupied: {
		CustomerPaid: true,
	},
}

var TableInfoStatusTransitionRules = map[string]map[string]bool{
	TableIsAvailable: {
		TableIsOccupied: true,
		TableIsReserved: true,
	},
	TableIsOccupied: {
		TableIsAvailable: true,
	},
	TableIsReserved: {
		TableIsOccupied:  true,
		TableIsAvailable: true,
	},
}
