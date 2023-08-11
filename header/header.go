package header

type Header struct {
	Magic                     [4]byte
	Type                      [4]byte
	Filesize                  uint32
	PointerTableOffset        uint32
	PointerTableSize          uint32
	Unknown                   uint32
	FooterSize                uint32
	_                         uint32
	Unknown2                  uint32
	Unknown3                  uint32
	_                         uint32
	Unknown4                  uint32
	Unknown5                  uint32
	Unknown6                  uint32
	Unknown7                  uint32
	Unknown8                  uint32
	HeaderSize                uint32
	Unknown9                  uint32
	_                         [12]byte
	NumberOfChannels          uint32
	ChannelTableOffset        uint32
	NumberOfRegions           uint32
	RegionTableOffset         uint32
	NumberOfGenres            uint32
	GenreTableOffset          uint32
	NumberOfAudioTypes        uint32
	AudioTypeTableOffset      uint32
	NumberOfResolutionTypes   uint32
	ResolutionTypeTableOffset uint32
	NumberOfBilingualTables   uint32
	BilingualTableOffset      uint32
	NumberOfUnknownTables     uint32
	UnknownTableOffset        uint32
	NumberOfCopyrightTables   uint32
	CopyrightTableOffset      uint32
	SupportTextOffset         uint32
}

var (
	Magic = [4]byte{'H', 'D', 'P', 'K'}
	Type  = [4]byte{'0', '0', '1', 'B'}
)

func (h *HeaderFile) MakeHeader() {
	h.Header = Header{
		Magic:              Magic,
		Type:               Type,
		Filesize:           0,
		PointerTableOffset: 0,
		PointerTableSize:   0,
		Unknown:            1,
		FooterSize:         5,
		Unknown2:           1,
		Unknown3:           18,
		Unknown4:           720,
		Unknown5:           10,
		Unknown6:           200,
		Unknown7:           200,
		Unknown8:           12,
		HeaderSize:         120,
		Unknown9:           30,
		// TODO: Make NumberOfChannels and Offset dynamic. We currently force some channels for testing
		NumberOfChannels:          4,
		ChannelTableOffset:        120,
		NumberOfRegions:           0,
		RegionTableOffset:         0,
		NumberOfGenres:            0,
		GenreTableOffset:          0,
		NumberOfAudioTypes:        0,
		AudioTypeTableOffset:      0,
		NumberOfResolutionTypes:   0,
		ResolutionTypeTableOffset: 0,
		NumberOfBilingualTables:   0,
		BilingualTableOffset:      0,
		NumberOfUnknownTables:     0,
		UnknownTableOffset:        0,
		NumberOfCopyrightTables:   0,
		CopyrightTableOffset:      0,
		SupportTextOffset:         0,
	}
}
