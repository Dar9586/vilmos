package pixel

import "strconv"

type Pixel struct {
	R int
	G int
	B int
	A int
}

func (p *Pixel) String() string {
	return "{R: " + strconv.Itoa(p.R) + " G: " + strconv.Itoa(p.G) + " B: " + strconv.Itoa(p.B) + " A: " + strconv.Itoa(p.A) + "}"
}

func (p *Pixel) Equals(other Pixel) bool {
	return (p.R == other.R) && (p.G == other.G) && (p.B == other.B) && (p.A == other.A)
}
