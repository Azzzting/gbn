package main

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	bn254 "github.com/Azzzting/gbn/include"
	"github.com/stretchr/testify/assert"
)

func format_bn_barrett(bn []uint32, length int) string { // 统一格式
	buf := make([]byte, length*4)
	for i := 0; i < length; i++ {
		idx := length - 1 - i
		binary.BigEndian.PutUint32(buf[i*4:], bn[idx])
	}
	return hex.EncodeToString(buf)
}

func Test_bn_barrett(t *testing.T) {
	hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	hex_p := "30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"
	hex_n := "30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001"
	hex_p_mu := "000000054a47462623a04a7ab074a5868073013ae965e1767cd4c086f3aed8a19bf90e51"
	hex_n_mu := "000000054a47462623a04a7ab074a58680730147144852009e880ae620703a6be1de9259"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	n := make([]uint32, 8)
	r := make([]uint32, 8)
	p := make([]uint32, 8)
	p_mu := make([]uint32, 9)
	n_mu := make([]uint32, 9)
	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)
	bn254.BN_from_hex(p, 8, hex_p)
	bn254.BN_from_hex(p_mu, 9, hex_p_mu)
	bn254.BN_from_hex(n, 8, hex_n)
	bn254.BN_from_hex(n_mu, 9, hex_n_mu)
	var result string

	bn254.BN_barrett_mod_mul(r, a, b, p, p_mu, 8)
	result = format_bn_barrett(r, 8)
	assert.Equal(t, "26233c6c1df2a1b4fec81ed8a1a4fb2450a06a8adf3576d183a9f2b87b6d2b72", result)

	bn254.BN_barrett_mod_sqr(r, a, p, p_mu, 8)
	result = format_bn_barrett(r, 8)
	assert.Equal(t, "13b887f45214e1eed372c7b44bd4a7a8d60518114e10c581176cba7fdeccd283", result)

	bn254.BN_barrett_mod_exp(r, a, b, p, p_mu, 8)
	result = format_bn_barrett(r, 8)
	assert.Equal(t, "09223b41c088f253f48bcd362e406afe59f49bcff34cb9e1ce5d910c38f06476", result)

	bn254.BN_barrett_mod_inv(r, a, p, p_mu, 8)
	result = format_bn_barrett(r, 8)
	assert.Equal(t, "1c247c36f20b94c90c886532c309246addd378f27906754263f7af21d12a71a7", result)

	bn254.BN_barrett_mod_mul(r, a, b, n, n_mu, 8)
	result = format_bn_barrett(r, 8)
	assert.Equal(t, "0eac5f427f2539e7ea9f0acefb29aa858ebba3f313ded98db15c1a0e74780a45", result)

	bn254.BN_barrett_mod_sqr(r, a, n, n_mu, 8)
	result = format_bn_barrett(r, 8)
	assert.Equal(t, "08f43c386510239e5dc1ee4f4bf28269c89ab724653e0dcf6a303afa66f21313", result)

	bn254.BN_barrett_mod_exp(r, a, b, n, n_mu, 8)
	result = format_bn_barrett(r, 8)
	assert.Equal(t, "05fc3467db36cc7bda4d6258abdf4a8d440e8799a393cbcabb171b07239e0a9d", result)

	bn254.BN_barrett_mod_inv(r, a, n, n_mu, 8)
	result = format_bn_barrett(r, 8)
	assert.Equal(t, "22606952067292f23ae4bf1e3a8e6bd7b47b6b38fc9e91b11dfcb0f2e73692c2", result)
}
