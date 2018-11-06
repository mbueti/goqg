package model

type kernel struct {
	q            [3][]float64
	qh           [3][]complex64
	Qh           [3][]complex64
	ph           [3][]complex64
	Ph           [3][]complex64
	u            [3][]float64
	v            [3][]float64
	uh           [3][]complex64
	vh           [3][]complex64
	uq           [3][]float64
	vq           [3][]float64
	uqh          [3][]complex64
	vqh          [3][]complex64
	dqhdt        [3][]complex64
	dqhdtP       [3][]complex64
	dqhdtPP      [3][]complex64
	dummyFFTIn   []float64
	dummyFFTOut  []float64
	dummyIFFTIn  []float64
	dummyIFFTOut []float64
	kk           []float64
	ll           []float64
	a            [4][]complex64
	ik           []complex64
	il           []complex64
	k2l2         [2][]float64
	Ubg          [2][]float64
	Vbg          [2][]float64
	ikQy         [3][]complex64
	filter       [2][]float64
	rek          float64
	rbg          float64
}
