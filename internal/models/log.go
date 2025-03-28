package models

type Log struct {
}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) Randomize() *Log {
	return l
}
