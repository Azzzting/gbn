package main

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	bn254 "github.com/Azzzting/gbn/include"
	"github.com/stretchr/testify/assert"
)

func format_bn_montgomery(bn []uint32, length int) string { // 统一格式
	buf := make([]byte, length*4)
	for i := 0; i < length; i++ {
		idx := length - 1 - i
		binary.BigEndian.PutUint32(buf[i*4:], bn[idx])
	}
	return hex.EncodeToString(buf)
}

func Test_bn_montgomery(t *testing.T) {
	hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	hex_p := "30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"
	hex_n := "30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001"
	hex_p_inv_neg := "f57a22b791888c6bd8afcbd01833da809ede7d651eca6ac987d20782e4866389"
	hex_p_one_sqr := "06d89f71cab8351f47ab1eff0a417ff6b5e71911d44501fbf32cfc5b538afa89"
	hex_n_inv_neg := "73f82f1d0d8341b2e39a9828990623916586864b4c6911b3c2e1f593efffffff"
	hex_n_one_sqr := "0216d0b17f4e44a58c49833d53bb808553fe3ab1e35c59e31bb8e645ae216da7"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	n := make([]uint32, 8)
	r := make([]uint32, 8)
	p := make([]uint32, 8)
	p_inv_neg := make([]uint32, 8)
	p_one_sqr := make([]uint32, 8)
	n_inv_neg := make([]uint32, 8)
	n_one_sqr := make([]uint32, 8)

	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)
	bn254.BN_from_hex(p, 8, hex_p)
	bn254.BN_from_hex(p_inv_neg, 8, hex_p_inv_neg)
	bn254.BN_from_hex(p_one_sqr, 8, hex_p_one_sqr)
	var result string

	bn254.BN_mod_mul_montgomery(r, a, b, p_one_sqr, p, p_inv_neg, 8)
	result = format_bn_montgomery(r, 8)
	assert.Equal(t, "26233c6c1df2a1b4fec81ed8a1a4fb2450a06a8adf3576d183a9f2b87b6d2b72", result)

	bn254.BN_mod_sqr_montgomery(r, a, p_one_sqr, p, p_inv_neg, 8)
	result = format_bn_montgomery(r, 8)
	assert.Equal(t, "13b887f45214e1eed372c7b44bd4a7a8d60518114e10c581176cba7fdeccd283", result)

	bn254.BN_from_hex(n, 8, hex_n)
	bn254.BN_from_hex(n_inv_neg, 8, hex_n_inv_neg)
	bn254.BN_from_hex(n_one_sqr, 8, hex_n_one_sqr)

	bn254.BN_mod_mul_montgomery(r, a, b, n_one_sqr, n, n_inv_neg, 8)
	result = format_bn_montgomery(r, 8)
	assert.Equal(t, "0eac5f427f2539e7ea9f0acefb29aa858ebba3f313ded98db15c1a0e74780a45", result)

	bn254.BN_mod_sqr_montgomery(r, a, n_one_sqr, n, n_inv_neg, 8)
	result = format_bn_montgomery(r, 8)
	assert.Equal(t, "08f43c386510239e5dc1ee4f4bf28269c89ab724653e0dcf6a303afa66f21313", result)
}
