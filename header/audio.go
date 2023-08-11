package header

import "unicode/utf16"

type AudioTable struct {
	TextOffset  uint32
	AudioTypeId uint16
	_           uint16
}

var audioTypes = []AudioTable{
	{
		TextOffset:  0,
		AudioTypeId: 256,
	},
	{
		TextOffset:  0,
		AudioTypeId: 512,
	},
	{
		TextOffset:  0,
		AudioTypeId: 768,
	},
	{
		TextOffset:  0,
		AudioTypeId: 2304,
	},
}

func (h *HeaderFile) MakeAudioTable() {
	h.Header.AudioTypeTableOffset = h.GetCurrentSize() - 32

	for _, audio := range audioTypes {
		h.AudioTypes = append(h.AudioTypes, audio)
		h.UpdatePointerTable(h.GetCurrentSize() - 8)
	}

	h.Header.NumberOfAudioTypes = 4
}

func (h *HeaderFile) MakeAudioText() {
	coolList := []string{"Single", "Dual", "Stereo", "5.1"}

	for i := 0; i < len(h.AudioTypes); i++ {
		h.AudioTypes[i].TextOffset = h.GetCurrentSize() - 32
		h.AudioText = append(h.AudioText, utf16.Encode([]rune(coolList[i]))...)

		// If current size modulo 4 is 0 we must append 4 null bytes
		if h.GetCurrentSize()%4 == 0 {
			h.AudioText = append(h.AudioText, []uint16{0, 0}...)
		}

		// Apply modulo padding if needed
		for h.GetCurrentSize()%4 != 0 {
			h.AudioText = append(h.AudioText, 0)
		}
	}
}
