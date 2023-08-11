package constants

// BroadcastType are the varying types of broadcast formats.
// These include satellite digital + analog and terrestrial digital + analog
type BroadcastType uint16

const (
	SatelliteDigital   BroadcastType = 2
	SatelliteAnalog    BroadcastType = 18
	TerrestrialDigital BroadcastType = 9
	TerrestrialAnalog  BroadcastType = 25
)
