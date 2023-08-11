package header

import (
	"unicode/utf16"
)

type RegionTable struct {
	RegionCode                  uint16
	PrefectureArrayPosition     uint8
	_                           uint8
	RegionNameOffset            uint32
	NumberOfChannels            uint32
	ChannelSelectionTableOffset uint32
}

func (h *HeaderFile) MakeRegionTable() {
	h.Header.RegionTableOffset = h.GetCurrentSize() - 32

	for i := 1; i < 48; i++ {
		region := RegionTable{
			RegionCode:                  uint16(i),
			PrefectureArrayPosition:     uint8(i),
			RegionNameOffset:            0,
			NumberOfChannels:            4,
			ChannelSelectionTableOffset: 0,
		}

		h.Regions = append(h.Regions, region)
		// Text Offset
		h.UpdatePointerTable(h.GetCurrentSize() - 12)
		// ChannelSelectionTable Offset
		h.UpdatePointerTable(h.GetCurrentSize() - 4)
	}

	h.Header.NumberOfRegions = uint32(len(h.Regions))
}

func (h *HeaderFile) MakeRegionText() {
	for i := 0; i < len(h.Regions); i++ {
		h.Regions[i].RegionNameOffset = h.GetCurrentSize() - 32
		h.RegionText = append(h.RegionText, utf16.Encode([]rune("WiiLink City"))...)

		if h.GetCurrentSize()%4 == 0 {
			h.RegionText = append(h.RegionText, []uint16{0, 0}...)
		} else {
			h.RegionText = append(h.RegionText, 0)
		}
	}
}
