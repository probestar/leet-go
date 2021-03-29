package l190

func reverseBits(num uint32) uint32 {
	var n uint32 = 0
	for i := 0; i < 32 && num > 0; i++ {
		n |= num & 1 << (31 - i)
		num = num >> 1
	}
	return n
}
