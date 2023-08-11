package header

import "unicode/utf16"

type GenreTable struct {
	MainGenrePos    uint8
	SubGenrePos     uint8
	_               uint16
	GenreNameOffset uint32
}

var tempGenres = []GenreTable{
	{
		MainGenrePos:    10,
		SubGenrePos:     0,
		GenreNameOffset: 0,
	},
	{
		MainGenrePos:    10,
		SubGenrePos:     11,
		GenreNameOffset: 0,
	},
}

func (h *HeaderFile) MakeGenreTable() {
	h.Header.GenreTableOffset = h.GetCurrentSize() - 32

	for _, genre := range tempGenres {
		h.Genres = append(h.Genres, genre)
		h.UpdatePointerTable(h.GetCurrentSize() - 4)
	}

	h.Header.NumberOfGenres = uint32(len(tempGenres))
}

func (h *HeaderFile) MakeGenreText() {
	coolList := []string{"News", "Citrus Execution"}

	for i := 0; i < len(h.Genres); i++ {
		h.Genres[i].GenreNameOffset = h.GetCurrentSize() - 32
		h.GenreText = append(h.GenreText, utf16.Encode([]rune(coolList[i]))...)
		
		if h.GetCurrentSize()%4 == 0 {
			h.GenreText = append(h.GenreText, []uint16{0, 0}...)
		} else {
			h.GenreText = append(h.GenreText, 0)
		}
	}
}
