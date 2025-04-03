package main

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	bn254 "github.com/Azzzting/gbn/include"
	"github.com/stretchr/testify/assert"
)

func format_fn(bn []uint32, length int) string { // 统一格式
	buf := make([]byte, length*4)
	for i := 0; i < length; i++ {
		idx := length - 1 - i
		binary.BigEndian.PutUint32(buf[i*4:], bn[idx])
	}
	return hex.EncodeToString(buf)
}

func Test_fn(t *testing.T) {
	hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	r := make([]uint32, 8)
	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)
	var result string

	bn254.FN_from_hex(a, hex_a)
	bn254.FN_from_hex(b, hex_b)
	bn254.FN_add(r, a, b)
	bn254.FN_get_BN(r, r)
	result = format_fn(r, 8)
	assert.Equal(t, "1befa71cbccf15f1b86fc535f2cff36a01e50692e41e37f6051fd274038f069c", result)

	bn254.FN_sub(r, a, b)
	bn254.FN_get_BN(r, r)
	result = format_fn(r, 8)
	assert.Equal(t, "2a5513860fc876382d0c7108e455c5b45d8c1191888242b9e7282e60150255c3", result)

	bn254.FN_sub(r, b, a)
	bn254.FN_get_BN(r, r)
	result = format_fn(r, 8)
	assert.Equal(t, "060f3aecd16929f18b43d4ad9d2b92a8caa7d6b6f1372dd75cb9c733dafdaa3e", result)

	bn254.FN_neg(r, a)
	bn254.FN_get_BN(r, r)
	result = format_fn(r, 8)
	assert.Equal(t, "2574185aeb7eaa29a1ba4d7256af27fc8c95505a8045eb81efaeeff3dbb751d2", result)

	bn254.FN_mul(r, a, b)
	bn254.FN_get_BN(r, r)
	result = format_fn(r, 8)
	assert.Equal(t, "0eac5f427f2539e7ea9f0acefb29aa858ebba3f313ded98db15c1a0e74780a45", result)

	bn254.FN_sqr(r, a)
	bn254.FN_get_BN(r, r)
	result = format_fn(r, 8)
	assert.Equal(t, "08f43c386510239e5dc1ee4f4bf28269c89ab724653e0dcf6a303afa66f21313", result)

	var a_fn bn254.FN_t
	for i := 0; i < len(a); i++ {
		a_fn[i] = a[i]
	}
	var r_fn bn254.FN_t
	for i := 0; i < len(r); i++ {
		r_fn[i] = r[i]
	}
}
