package header

import (
	"TVnoTomo/constants"
	"unicode/utf16"
)

// ChannelTable contains all metadata for a channel.
type ChannelTable struct {
	BroadcastType      constants.BroadcastType
	ChannelId          uint16
	TextOffset         uint32
	BroadcastTypeAgain uint8
	_                  [3]byte
}

// ChannelSelectionTable contains the metadata for a channel within a specific region.
type ChannelSelectionTable struct {
	BroadcastType  constants.BroadcastType
	ChannelId      uint16
	Position       uint16
	ChannelNumber  uint16
	IsAutoSelected bool
	_              [3]byte
}

// We will have 1 channel per type
var tempChannelArray = []ChannelTable{
	{
		BroadcastType:      constants.TerrestrialDigital,
		ChannelId:          1024,
		TextOffset:         0,
		BroadcastTypeAgain: uint8(constants.TerrestrialDigital),
	},
	{
		BroadcastType:      constants.SatelliteDigital,
		ChannelId:          101,
		TextOffset:         0,
		BroadcastTypeAgain: uint8(constants.SatelliteDigital),
	},
	{
		BroadcastType:      constants.TerrestrialAnalog,
		ChannelId:          257,
		TextOffset:         0,
		BroadcastTypeAgain: uint8(constants.TerrestrialAnalog),
	},
	{
		BroadcastType:      constants.SatelliteAnalog,
		ChannelId:          73,
		TextOffset:         0,
		BroadcastTypeAgain: uint8(constants.SatelliteAnalog),
	},
}

var tempChannelSelectionArray = []ChannelSelectionTable{
	{
		BroadcastType:  constants.TerrestrialDigital,
		ChannelId:      1024,
		Position:       1,
		ChannelNumber:  101,
		IsAutoSelected: false,
	},
	{
		BroadcastType:  constants.SatelliteDigital,
		ChannelId:      101,
		Position:       0,
		ChannelNumber:  101,
		IsAutoSelected: false,
	},
	{
		BroadcastType:  constants.TerrestrialAnalog,
		ChannelId:      257,
		Position:       1,
		ChannelNumber:  0,
		IsAutoSelected: false,
	},
	{
		BroadcastType:  constants.SatelliteAnalog,
		ChannelId:      73,
		Position:       0,
		ChannelNumber:  101,
		IsAutoSelected: false,
	},
}

func (h *HeaderFile) MakeChannelTable() {
	for _, table := range tempChannelArray {
		h.Channels = append(h.Channels, table)
		h.UpdatePointerTable(h.GetCurrentSize() - 8)
	}
}

func (h *HeaderFile) MakeChannelSelectionTable() {
	for i := 0; i < len(h.Regions); i++ {
		h.Regions[i].ChannelSelectionTableOffset = h.GetCurrentSize() - 32

		for _, table := range tempChannelSelectionArray {
			h.ChannelSelectionTable = append(h.ChannelSelectionTable, table)
		}
	}
}

func (h *HeaderFile) MakeChannelText() {
	for i := 0; i < len(h.Channels); i++ {
		h.Channels[i].TextOffset = h.GetCurrentSize() - 32
		h.ChannelText = append(h.ChannelText, utf16.Encode([]rune("Sketch TV"))...)

		// Pad the text.
		if h.GetCurrentSize()%4 == 0 {
			h.ChannelText = append(h.ChannelText, []uint16{0, 0}...)
		} else {
			// Only real possibility is that it is equal to 2.
			h.ChannelText = append(h.ChannelText, 0)
		}
	}
}
