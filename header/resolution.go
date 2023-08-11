package header

import "unicode/utf16"

type ResolutionTable struct {
	TextOffset       uint32
	ResolutionTypeId uint32
}

func (h *HeaderFile) MakeResolutionTable() {
	h.Header.ResolutionTypeTableOffset = h.GetCurrentSize() - 32
	h.Resolutions = []ResolutionTable{
		{
			TextOffset:       0,
			ResolutionTypeId: 16777216,
		},
	}
	h.UpdatePointerTable(h.GetCurrentSize() - 8)
	h.Header.NumberOfResolutionTypes = 1
}

func (h *HeaderFile) MakeResolutionText() {
	coolList := []string{"HDTV"}

	for i := 0; i < len(h.Resolutions); i++ {
		h.Resolutions[i].TextOffset = h.GetCurrentSize() - 32
		h.ResolutionText = append(h.ResolutionText, utf16.Encode([]rune(coolList[i]))...)

		// If current size modulo 4 is 0 we must append 4 null bytes
		if h.GetCurrentSize()%4 == 0 {
			h.ResolutionText = append(h.ResolutionText, []uint16{0, 0}...)
		} else {
			h.ResolutionText = append(h.ResolutionText, 0)
		}
	}
}
