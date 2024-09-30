package types

const (
	Stdout LogOutput = iota
	Fileout
)

type LogOutput uint32
