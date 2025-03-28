package models

type SInt uint64

func (i SInt) Int() int {
	return int(i)
}

func (i SInt) Uint() uint {
	return uint(i)
}
