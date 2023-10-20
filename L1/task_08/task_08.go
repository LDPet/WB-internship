package main

func SetBit(num int64, pos int) int64 {
	return num | (1 << (63 - pos))
}

func ResetBit(num int64, pos int) int64 {
	return num & ^(1 << (63 - pos))
}

func ToggleBit(num int64, pos int) int64 {
	return num ^ (1 << (63 - pos))
}

func main() {
	a := int64(1)
	b := SetBit(a, 60)
	c := ResetBit(a, 63)
	d := ToggleBit(b, 60)

	println(a, b, c, d)
}
