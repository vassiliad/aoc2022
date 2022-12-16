package utilities

type BitMask uint64

func (b *BitMask) Get(index int) bool {
	return *b&(1<<uint64(index)) > 0
}

func (b *BitMask) Set(index int) {
	*b |= (1 << uint64(index))
}

func (b *BitMask) UnSet(index int) {
	*b &= ^(1 << uint64(index))
}
