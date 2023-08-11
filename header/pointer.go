package header

/*
Note from Sketch:
The pointer table is something I really don't understand why it exists, but it does so here we go.
For every table but the RegionTable, a pointer table value will point to the TextOffset field in that table.
For the RegionTable, it will point to both the TextOffset and the ChannelSelectionTable offset.
*/

func InitPointerTable() {
	// Hardcoded values. Points to fields in the header.
	tempPointerTable = []uint32{
		56,
		64,
		72,
		80,
		88,
		96,
		104,
		112,
		116,
	}
}

func (h *HeaderFile) UpdatePointerTable(value uint32) {
	tempPointerTable = append(tempPointerTable, value-32)
}
