package fp2

type Cxform struct {
	RedMultTerm   int16 // 8.8
	GreenMultTerm int16 // 8.8
	BlueMultTerm  int16 // 8.8
	AlphaMultTerm int16 // 8.8
	RedAddTerm    uint8 // abs
	GreenAddTerm  uint8 // abs
	BlueAddTerm   uint8 // abs
	AlphaAddTerm  uint8 // abs
}
