package x2j

import (
	"io"
)

var X2jCharsetReader func(charset string, input io.Reader)(io.Reader, error)

type Node struct {
	dup     bool
	attr    bool
	key     string
	val     string
	nodes   []*Node
}
