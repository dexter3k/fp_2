package fp2

type SwfHeader struct {
	Compression uint8 // Unused in FP2
	Version     uint8
	FileLength  uint32
	FrameSize   Rect
	FrameRate   uint16 // 8.8
	FrameCount  uint16
}

type SwfFile struct {
	SwfStream
	SwfHeader

	headerDone bool

	buffer [8]byte
}

func (s *SwfFile) processNextTag() {
	idLength := s.ReadU16()
	code := idLength >> 6

	length := uint32(idLength & 0x3f)
	if length == 0x3f {
		length = s.ReadU32()
	}
}
