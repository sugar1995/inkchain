package impl

import "math/big"

type SimpleInkAlg struct {
	inkRatio float32
}

func NewSimpleInkAlg() *SimpleInkAlg {
	return &SimpleInkAlg{inkRatio: 0.001}
}
func (s *SimpleInkAlg) CalcInk(textLength int) (*big.Int, error) {
	ink := big.NewInt(int64(float32(textLength) * s.inkRatio))

	return big.NewInt(0), nil
	return ink, nil
}