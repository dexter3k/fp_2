package fp2

type Matrix struct {
	ScaleX      int32 // 16.16
	ScaleY      int32 // 16.16
	RotateSkew0 int32 // 16.16
	RotateSkew1 int32 // 16.16
	TranslateX  int32 // twips
	TranslateY  int32 // twips
}

func (m *Matrix) Transform(p *Point, out *Point) *Point {
	if out == nil {
		out = &Point{}
	}

	out.X = p.X * m.ScaleX + p.Y * m.RotateSkew1 + m.TranslateX
	out.Y = p.Y * m.ScaleY + p.Y * m.RotateSkew0 + m.TranslateY

	return out
}
