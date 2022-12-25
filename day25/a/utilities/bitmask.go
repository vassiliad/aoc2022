package utilities

type BitMask uint8

func (b *BitMask) Get(index int) bool {
	return *b&(1<<BitMask(index)) > 0
}

func (b *BitMask) Set(index int) {
	*b |= (1 << BitMask(index))
}

func (b *BitMask) UnSet(index int) {
	*b &= ^(1 << BitMask(index))
}
