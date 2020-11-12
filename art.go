package artree

import (
	"unsafe"

	"github.com/templexxx/cpu"
)

func init() {
	if !cpu.X86.HasCMPXCHG16B {
		panic("art need CMPXCHG16B feature, but not supported in this machine")
	}
}

// ART implements The Adaptive Radix Tree.
type ART struct {
	root unsafe.Pointer
	n    uint64
}

// Insert is used to add a newentry or update
// an existing entry. Returns if updated.
func (t *ART) Insert(key []byte, value unsafe.Pointer) bool {
	if t.root == nil {

	}
	return false
}

func (t *ART) Search() {

}

func (t *ART) Delete() {

}

type Item interface {
	Less(than Item) bool
}

type ItemIterator func(i Item) bool

func (t *ART) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {

}
