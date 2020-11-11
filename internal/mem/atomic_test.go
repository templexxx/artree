package mem

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

var (
	magic128 = make([]byte, 16)
)

func init() {
	binary.LittleEndian.PutUint64(magic128[:8], 0xdeddeadbeefbeef)
	binary.LittleEndian.PutUint64(magic128[8:], 0xdeddeadbeefbeef)
}

func TestAtomicCAS16B(t *testing.T) {

	var x struct {
		before []uint8
		i      []byte
		after  []uint8
	}
	x.before = magic128
	x.after = magic128

	old := make([]byte, 16)
	newV := make([]byte, 16)
	nnv := make([]byte, 16)
	x.i = MakeAlignedBlock(16, 16)

	for val := uint64(1); val+val > val; val += val {

		binary.LittleEndian.PutUint64(old[:8], val)
		binary.LittleEndian.PutUint64(old[8:], ^val)

		binary.LittleEndian.PutUint64(newV[:8], val+1)
		binary.LittleEndian.PutUint64(newV[8:], ^(val + 1))

		copy(x.i, old)

		if !AtomicCAS16B(&x.i[0], &old[0], &newV[0]) {
			t.Fatal("should have swapped")
		}
		if !bytes.Equal(x.i, newV) {
			t.Fatalf("wrong x.i after swap: x.i=%#x exp=%#x", x.i, newV)
		}

		copy(x.i, newV)

		binary.LittleEndian.PutUint64(nnv[:8], val+2)
		binary.LittleEndian.PutUint64(nnv[8:], ^(val + 2))

		if AtomicCAS16B(&x.i[0], &old[0], &nnv[0]) {
			t.Fatal("should not have swapped")
		}
		if !bytes.Equal(x.i, newV) {
			t.Fatalf("wrong x.i after swap: x.i=%#x val+1=%#x", x.i, newV)
		}
	}
	if !bytes.Equal(x.before, magic128) || !bytes.Equal(x.after, magic128) {
		t.Fatal("wrong magic")
	}
}

func BenchmarkCAS16B(b *testing.B) {

	dst := MakeAlignedBlock(16, 16)
	o := make([]byte, 16)
	n := make([]byte, 16)
	rand.Read(o)
	rand.Read(n)

	b.SetBytes(16)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = AtomicCAS16B(&dst[0], &o[0], &n[0])
	}
}