package domain

type Bite struct {
	weight int64
}

func NewBite(weight int64) Bite {
	return Bite{
		weight: weight,
	}
}

func (b *Bite) GetWeight() int64 {
	return b.weight
}
