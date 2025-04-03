package main

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	bn254 "github.com/Azzzting/gbn/include"

	"github.com/stretchr/testify/assert"
)

func format_bn(bn []uint32, length int) string { // 统一格式
	buf := make([]byte, length*4)
	for i := 0; i < length; i++ {
		idx := length - 1 - i
		binary.BigEndian.PutUint32(buf[i*4:], bn[idx])
	}
	return hex.EncodeToString(buf)
}

func Test_bn(t *testing.T) {
	hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"

	hex_add_a_b := "1befa71cbccf15f1b86fc535f2cff36a01e50692e41e37f6051fd274038f069c"
	hex_sub_a_b := "f9f0c5132e96d60e74bc2b5262d46d57355829490ec8d228a34638cc250255c2"
	hex_sub_b_a := "060f3aecd16929f18b43d4ad9d2b92a8caa7d6b6f1372dd75cb9c733dafdaa3e"
	hex_mul_lo_a_b := "a08947459e744f16f8daee01f6943c441cd32e084d669ec1e666d75c20ac5203"
	hex_mul_a_b := "00b9ed7b9d5c4626d518f15fb9f4cf1a310f4b03a0191939f1c8324ef8f60c3b" +
		"a08947459e744f16f8daee01f6943c441cd32e084d669ec1e666d75c20ac5203"

	a := make([]uint32, 8)
	b := make([]uint32, 8)
	r := make([]uint32, 16)
	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)
	var result string

	//test BN_add
	bn254.BN_add(r, a, b, 8)
	result = format_bn(r, 8)
	assert.Equal(t, hex_add_a_b, result)

	//test BN_sub
	bn254.BN_sub(r, a, b, 8)
	result = format_bn(r, 8)
	assert.Equal(t, hex_sub_a_b, result)
	bn254.BN_sub(r, b, a, 8)
	result = format_bn(r, 8)
	assert.Equal(t, hex_sub_b_a, result)

	//test mul_lo
	bn254.BN_mul_lo(r, a, b, 8)
	result = format_bn(r, 8)
	assert.Equal(t, hex_mul_lo_a_b, result)

	//test mul
	bn254.BN_mul(r, a, b, 8)
	result = format_bn(r, 16)
	assert.Equal(t, hex_mul_a_b, result)
}
