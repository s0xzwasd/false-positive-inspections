package p

type Iminterface interface {
	u() int
}

type AImpl struct{}

func (a AImpl) u() int { return 1 }
