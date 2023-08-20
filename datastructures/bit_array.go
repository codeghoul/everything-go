package datastructures

type BitArray struct {
	data []byte
	size uint32
}

func NewBitArray(size uint32) *BitArray {
	byteSize := (size + 7) / 8
	data := make([]byte, byteSize)
	return &BitArray{data: data, size: size}
}

func (ba *BitArray) Set(index uint32) {
	byteIndex := index / 8
	bitIndex := index % 8
	ba.data[byteIndex] |= 1 << (bitIndex)
}

func (ba *BitArray) Unset(index uint32) {
	byteIndex := index / 8
	bitIndex := index % 8
	ba.data[byteIndex] &= ^(1 << (bitIndex))
}

func (ba *BitArray) Get(index uint32) bool {
	byteIndex := index / 8
	bitIndex := index % 8
	return (ba.data[byteIndex] & (1 << (bitIndex))) != 0
}
