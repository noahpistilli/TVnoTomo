package header

import "unicode/utf16"

// RepeatedTable represents the structure of a table which is used many times in the file.
type RepeatedTable struct {
	TextOffset uint32
	Id         uint8
	_          [3]byte
}

type BilingualTable RepeatedTable

func (h *HeaderFile) MakeBilingualTable() {
	h.Header.BilingualTableOffset = h.GetCurrentSize() - 32
	h.Header.NumberOfBilingualTables = 1
	h.BilingualTable = []BilingualTable{
		{
			TextOffset: 0,
			Id:         1,
		},
	}
	h.UpdatePointerTable(h.GetCurrentSize() - 8)
}

func (h *HeaderFile) MakeBilingualText() {
	coolList := []string{"Yes"}

	for i := 0; i < len(h.BilingualTable); i++ {
		h.BilingualTable[i].TextOffset = h.GetCurrentSize() - 32
		h.BilingualText = append(h.BilingualText, utf16.Encode([]rune(coolList[i]))...)

		// If current size modulo 4 is 0 we must append 4 null bytes
		if h.GetCurrentSize()%4 == 0 {
			h.BilingualText = append(h.BilingualText, []uint16{0, 0}...)
		}

		// Apply modulo padding if needed
		for h.GetCurrentSize()%4 != 0 {
			h.BilingualText = append(h.BilingualText, 0)
		}
	}
}
