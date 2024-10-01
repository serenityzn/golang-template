package types

import (
	"net/netip"
	"strconv"
)

const (
	Stdout LogOutput = iota
	Fileout
)

type LogOutput uint32

type ConfHttpPort int

func (c ConfHttpPort) String() string {
	return strconv.Itoa(int(c))
}
func (c ConfHttpPort) Int() int {
	return int(c)
}

type LogLevel string

func (l LogLevel) Validate() bool {
	switch l {
	case "debug":
		return true
	case "info":
		return true
	case "error":
		return true
	default:
		return false
	}
}

type HttpHost string

func (h HttpHost) Validate() error {
	_, err := netip.ParseAddr(string(h))
	return err
}

func (h HttpHost) String() string {
	return string(h)
}

type LogName string

func (l LogName) Validate() bool {
	if len(l) < 1 {
		return false
	}
	return true
}

func (l LogName) String() string {
	return string(l)
}
