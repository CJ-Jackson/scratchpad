package maps

import (
	"testing"

	//
	"reflect"
	"unsafe"
)

func BenchmarkMap1S(b *testing.B) {
	m := make(map[string]int)

	k_tmp := [12]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	k := string(k_tmp[:])
	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		counter += m[k]
	}
}

func BenchmarkMap1S3U(b *testing.B) {
	m := make(map[string]int)

	k_tmp := [3]uint32{0, 1, 2}
	h := reflect.StringHeader{Data: uintptr(unsafe.Pointer(&k_tmp[0])), Len: len(k_tmp) * int(unsafe.Sizeof(k_tmp[0]))}
	k := *(*string)(unsafe.Pointer(&h))
	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		hh := reflect.StringHeader{Data: uintptr(unsafe.Pointer(&k_tmp[0])), Len: 12}
		kk := *(*string)(unsafe.Pointer(&hh))
		counter += m[kk]
	}
}
func BenchmarkMap1Str3U(b *testing.B) {
	m := make(map[string]int)

	k_tmp := [3]uint32{0, 1, 2}
	h := (*[12]byte)(unsafe.Pointer(&k_tmp[0]))
	k := string(h[:])
	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		hh := (*[12]byte)(unsafe.Pointer(&k_tmp[0]))
		kk := string(hh[:])
		counter += m[kk]
	}
}

func BenchmarkMap12C(b *testing.B) {
	m := make(map[[12]byte]int)

	k := [12]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		counter += m[k]
	}
}

func BenchmarkMap3U(b *testing.B) {
	m := make(map[[3]uint32]int)

	k := [3]uint32{0, 1, 2}
	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		counter += m[k]
	}
}

func BenchmarkMap2L(b *testing.B) {
	m := make(map[[2]uint64]int)

	k := [2]uint64{1234, 456}
	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		counter += m[k]
	}
}

func BenchmarkMap1L(b *testing.B) {
	m := make(map[uint64]int)

	k := uint64(1234)
	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		counter += m[k]
	}
}

func BenchmarkMap1K(b *testing.B) {
	m := make(map[interface{}]int)

	k := &struct {
		k interface{}
	}{k: 1234}

	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		counter += m[k]
	}
}

type M struct {
	k uint64
}

func BenchmarkMap1M(b *testing.B) {
	m := make(map[M]int)

	k := M{1234}

	m[k] = 1

	b.ResetTimer()

	counter := 0
	for i := 0; i < b.N; i++ {
		counter += m[k]
	}
}
