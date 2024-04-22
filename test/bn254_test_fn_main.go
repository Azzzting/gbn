package main

import (
	bn254_fn "github.com/Azzzting/gbn/include"
)

func test_bn() {
	hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	var a *[]uint32
	var b *[]uint32
	bn254_fn.BN_from_hex(a, 8, hex_a)
	bn254_fn.BN_from_hex(b, 8, hex_b)
	bn254_fn.BN_print(" a = ", a, 8)
	bn254_fn.BN_print(" b = ", b, 8)
}

func main() {
	test_bn()
}
