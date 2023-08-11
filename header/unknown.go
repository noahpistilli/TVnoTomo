package header

import "unicode/utf16"

type UnknownTable RepeatedTable

func (h *HeaderFile) MakeUnknownTable() {
	h.Header.UnknownTableOffset = h.GetCurrentSize() - 32
	h.Header.NumberOfUnknownTables = 1
	h.Unknown = []UnknownTable{
		{
			TextOffset: 0,
			Id:         1,
		},
	}
	h.UpdatePointerTable(h.GetCurrentSize() - 8)
}

func (h *HeaderFile) MakeUnknownText() {
	coolList := []string{"idk"}

	for i := 0; i < len(h.Unknown); i++ {
		h.Unknown[i].TextOffset = h.GetCurrentSize() - 32
		h.UnknownText = append(h.UnknownText, utf16.Encode([]rune(coolList[i]))...)

		// If current size modulo 4 is 0 we must append 4 null bytes
		if h.GetCurrentSize()%4 == 0 {
			h.UnknownText = append(h.UnknownText, []uint16{0, 0}...)
		}

		// Apply modulo padding if needed
		for h.GetCurrentSize()%4 != 0 {
			h.UnknownText = append(h.UnknownText, 0)
		}
	}
}
