// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__u64toa = 64
)

const (
    _stack__u64toa = 8
)

const (
    _size__u64toa = 1216
)

var (
    _pcsp__u64toa = [][2]uint32{
        {0x1, 0},
        {0xa5, 8},
        {0xa6, 0},
        {0x1cf, 8},
        {0x1d0, 0},
        {0x2f9, 8},
        {0x2fa, 0},
        {0x4b7, 8},
        {0x4c0, 0},
    }
)

var _cfunc_u64toa = []loader.CFunc{
    {"_u64toa_entry", 0,  _entry__u64toa, 0, nil},
    {"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
}
