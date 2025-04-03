package main

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	bn254 "github.com/Azzzting/gbn/include"
	"github.com/stretchr/testify/assert"
)

func format_fp(bn []uint32, length int) string { // 统一格式
	buf := make([]byte, length*4)
	for i := 0; i < length; i++ {
		idx := length - 1 - i
		binary.BigEndian.PutUint32(buf[i*4:], bn[idx])
	}
	return hex.EncodeToString(buf)
}

func Test_fp(t *testing.T) {
	hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"

	hex_fp_add_a_b := "1befa71cbccf15f1b86fc535f2cff36a01e50692e41e37f6051fd274038f069c"
	hex_fp_sub_a_b := "2a5513860fc876382d0c7108e455c5b4ccd993da773a9cb5df66c4e2fd7f5309"
	hex_fp_sub_b_a := "060f3aecd16929f18b43d4ad9d2b92a8caa7d6b6f1372dd75cb9c733dafdaa3e"
	hex_fp_neg_a := "2574185aeb7eaa29a1ba4d7256af27fcfbe2d2a36efe457de7ed8676c4344f18"
	hex_fp_mul_a_b := "26233c6c1df2a1b4fec81ed8a1a4fb2450a06a8adf3576d183a9f2b87b6d2b72"
	hex_fp_sqr_a := "13b887f45214e1eed372c7b44bd4a7a8d60518114e10c581176cba7fdeccd283"

	a := make([]uint32, 8)
	b := make([]uint32, 8)
	r := make([]uint32, 8)

	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)

	bn254.FP_from_hex(a, hex_a)
	bn254.FP_from_hex(b, hex_b)
	var result string

	bn254.FP_add(r, a, b)
	bn254.FP_get_BN(r, r)
	result = format_fp(r, 8)
	assert.Equal(t, hex_fp_add_a_b, result)

	bn254.FP_sub(r, a, b)
	bn254.FP_get_BN(r, r)
	result = format_fp(r, 8)
	assert.Equal(t, hex_fp_sub_a_b, result)

	bn254.FP_sub(r, b, a)
	bn254.FP_get_BN(r, r)
	result = format_fp(r, 8)
	assert.Equal(t, hex_fp_sub_b_a, result)

	bn254.FP_neg(r, a)
	bn254.FP_get_BN(r, r)
	result = format_fp(r, 8)
	assert.Equal(t, hex_fp_neg_a, result)

	bn254.FP_mul(r, a, b)
	bn254.FP_get_BN(r, r)
	result = format_fp(r, 8)
	assert.Equal(t, hex_fp_mul_a_b, result)

	bn254.FP_sqr(r, a)
	bn254.FP_get_BN(r, r)
	result = format_fp(r, 8)
	assert.Equal(t, hex_fp_sqr_a, result)

	var a_fp bn254.FP_t
	for i := 0; i < len(a); i++ {
		a_fp[i] = a[i]
	}
	var r_fp bn254.FP_t
	for i := 0; i < len(r); i++ {
		r_fp[i] = r[i]
	}

	bn254.FP_inv(&r_fp, a_fp)
	bn254.FP_get_BN(r_fp[:], r_fp[:])
	result = format_fp(r, 8)
}
// 1
