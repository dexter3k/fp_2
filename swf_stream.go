package fp2

type SwfStream struct {
	Data   []byte
	Offset int

	bits     uint8
	bitsLeft uint8
}

func (s *SwfStream) ResetBitStream() {
	s.bitsLeft = 0
}
