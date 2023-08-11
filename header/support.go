package header

import "unicode/utf16"

func (h *HeaderFile) MakeSupportText() {
	h.Header.SupportTextOffset = h.GetCurrentSize() - 32
	h.SupportText = append(h.SupportText, utf16.Encode([]rune("If you are reading this then Citrus has already been executed. The final phase is done. WiiLink has ended."))...)

	// If current size modulo 4 is 0 we must append 4 null bytes
	if h.GetCurrentSize()%4 == 0 {
		h.SupportText = append(h.SupportText, []uint16{0, 0}...)
	}

	// Apply modulo padding if needed
	for h.GetCurrentSize()%4 != 0 {
		h.SupportText = append(h.SupportText, 0)
	}
}
