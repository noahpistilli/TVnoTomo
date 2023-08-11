package header

import "unicode/utf16"

type CopyrightTable RepeatedTable

const (
	Allowed    uint8 = 0
	OneTime    uint8 = 2
	Prohibited uint8 = 3
)

var copyrightTables = []CopyrightTable{
	{
		TextOffset: 0,
		Id:         Allowed,
	},
	{
		TextOffset: 0,
		Id:         OneTime,
	},
	{
		TextOffset: 0,
		Id:         Prohibited,
	},
}

func (h *HeaderFile) MakeCopyrightTable() {
	h.Header.CopyrightTableOffset = h.GetCurrentSize() - 32
	h.Header.NumberOfCopyrightTables = 3

	for _, table := range copyrightTables {
		h.CopyrightTypes = append(h.CopyrightTypes, table)
		h.UpdatePointerTable(h.GetCurrentSize() - 8)
	}
}

func (h *HeaderFile) MakeCopyrightText() {
	coolList := []string{"Allowed", "Once", "Prohibited"}

	for i := 0; i < len(h.CopyrightTypes); i++ {
		h.CopyrightTypes[i].TextOffset = h.GetCurrentSize() - 32
		h.CopyrightText = append(h.CopyrightText, utf16.Encode([]rune(coolList[i]))...)

		// If current size modulo 4 is 0 we must append 4 null bytes
		if h.GetCurrentSize()%4 == 0 {
			h.CopyrightText = append(h.CopyrightText, []uint16{0, 0}...)
		}

		// Apply modulo padding if needed
		for h.GetCurrentSize()%4 != 0 {
			h.CopyrightText = append(h.CopyrightText, 0)
		}
	}
}
