package header

import (
	"bytes"
	"encoding/binary"
	"fmt"
	wc24 "github.com/SketchMaster2001/libwc24crypt"
	"github.com/wii-tools/lzx/lz10"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	AESKey = []byte{55, 216, 138, 225, 204, 194, 4, 24, 208, 63, 103, 123, 117, 180, 131, 42}
	IV     = []byte{181, 9, 109, 182, 149, 185, 150, 148, 101, 28, 213, 254, 120, 80, 39, 133}
)

type HeaderFile struct {
	// Data tables
	Header                Header
	Channels              []ChannelTable
	Regions               []RegionTable
	ChannelSelectionTable []ChannelSelectionTable
	Genres                []GenreTable
	AudioTypes            []AudioTable
	Resolutions           []ResolutionTable
	BilingualTable        []BilingualTable
	Unknown               []UnknownTable
	CopyrightTypes        []CopyrightTable

	/* Text Types */
	ChannelText    []uint16
	RegionText     []uint16
	GenreText      []uint16
	AudioText      []uint16
	ResolutionText []uint16
	BilingualText  []uint16
	UnknownText    []uint16
	CopyrightText  []uint16
	SupportText    []uint16

	// Pointer table
	PointerTable []uint32
}

var tempPointerTable []uint32

func MakeHeaderFile() {
	headerFile := HeaderFile{}

	// Init the pointer table
	InitPointerTable()
	fmt.Println(headerFile.GetCurrentSize())
	headerFile.MakeHeader()
	headerFile.MakeChannelTable()
	headerFile.MakeRegionTable()
	headerFile.MakeChannelSelectionTable()
	headerFile.MakeGenreTable()
	headerFile.MakeAudioTable()
	headerFile.MakeResolutionTable()
	headerFile.MakeBilingualTable()
	headerFile.MakeUnknownTable()
	headerFile.MakeCopyrightTable()
	headerFile.MakeChannelText()
	headerFile.MakeRegionText()
	headerFile.MakeGenreText()
	headerFile.MakeAudioText()
	headerFile.MakeResolutionText()
	headerFile.MakeBilingualText()
	headerFile.MakeUnknownText()
	headerFile.MakeCopyrightText()
	headerFile.MakeSupportText()

	// Now we add the pointer table.
	headerFile.Header.PointerTableSize = uint32(len(tempPointerTable))
	headerFile.Header.PointerTableOffset = headerFile.GetCurrentSize() - 32
	headerFile.PointerTable = tempPointerTable

	headerFile.Header.Filesize = headerFile.GetCurrentSize() + 13

	buffer := new(bytes.Buffer)
	headerFile.WriteAll(buffer)

	// Finally write the footer
	buffer.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0})
	buffer.WriteString("main")
	buffer.WriteByte(0)

	// Write and sign the file
	rsaData, err := ioutil.ReadFile("Private.pem")
	checkError(err)

	compressed, err := lz10.Compress(buffer.Bytes())
	checkError(err)

	encrypted, err := wc24.EncryptWC24(compressed, AESKey, IV, rsaData)
	checkError(err)

	err = os.WriteFile("header.dec", buffer.Bytes(), 0666)
	checkError(err)

	err = os.WriteFile("header.bin", encrypted, 0666)
	checkError(err)
}

func (h *HeaderFile) WriteAll(writer io.Writer) {
	// Header
	Write(writer, h.Header)

	// Channels
	Write(writer, h.Channels)

	// Regions + Channel Selection Table
	Write(writer, h.Regions)
	Write(writer, h.ChannelSelectionTable)

	// Genres
	Write(writer, h.Genres)

	// Audio Types
	Write(writer, h.AudioTypes)

	// Program Resolutions
	Write(writer, h.Resolutions)

	// Bilingual Table
	Write(writer, h.BilingualTable)

	// Unknown Table
	Write(writer, h.Unknown)

	// Copyright Table
	Write(writer, h.CopyrightTypes)

	/* Text based tables */
	Write(writer, h.ChannelText)
	Write(writer, h.RegionText)
	Write(writer, h.GenreText)
	Write(writer, h.AudioText)
	Write(writer, h.ResolutionText)
	Write(writer, h.BilingualText)
	Write(writer, h.UnknownText)
	Write(writer, h.CopyrightText)
	Write(writer, h.SupportText)

	// Pointer table
	Write(writer, h.PointerTable)
}

// GetCurrentSize returns the current size of our HeaderFile struct.
// This is useful for calculating the current offset of HeaderFile.
func (h *HeaderFile) GetCurrentSize() uint32 {
	buffer := bytes.NewBuffer([]byte{})
	h.WriteAll(buffer)

	return uint32(buffer.Len())
}

// Write writes the passed data to an io.Writer method.
// It does this in BigEndian mode as this is what the Wii uses.
func Write(writer io.Writer, data any) {
	err := binary.Write(writer, binary.BigEndian, data)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("TV no Tomo header.bin file generator has encountered a fatal error! Reason: %v\n", err)
	}
}
