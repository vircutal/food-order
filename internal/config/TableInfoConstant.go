package config

const (
	TableIsAvailable = "available"
	TableIsOccupied  = "occupied"
	TableIsReserved  = "reserved"
)

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

var TableInfoStatus = map[string]bool{
	TableIsAvailable: true,
	TableIsOccupied:  true,
	TableIsReserved:  true,
}
