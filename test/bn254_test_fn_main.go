package main

import (
	bn254_fn "github.com/Azzzting/gbn/include"
)

func test_bn() {
	hex_a := "aaf03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "aaf03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	r := make([]uint32, 16)
	bn254_fn.BN_from_hex(a, 8, hex_a)
	bn254_fn.BN_from_hex(b, 8, hex_b)
	bn254_fn.BN_print(" a = ", a, 8)
	bn254_fn.BN_print(" b = ", b, 8)
	//test BN_add
	bn254_fn.BN_add(r, a, b, 8)
	bn254_fn.BN_print(" r = ", r, 8)
	//test BN_sub
	bn254_fn.BN_sub(r, a, b, 8)
	bn254_fn.BN_print(" r = ", r, 8)
	//test mul_lo
	bn254_fn.BN_mul_lo(r, a, b, 8)
	bn254_fn.BN_print(" r = ", r, 8)
	//mul
	bn254_fn.BN_mul(r, a, b, 8)
	bn254_fn.BN_print(" r = ", r, 16)
}

func test_bn_mod() {
	hex_a := "aaf03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "aaf03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_n := "30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	n := make([]uint32, 8)
	r := make([]uint32, 16)

	bn254_fn.BN_from_hex(a, 8, hex_a)
	bn254_fn.BN_from_hex(b, 8, hex_b)
	bn254_fn.BN_from_hex(n, 8, hex_n)

	bn254_fn.BN_print(" a = ", a, 8)
	bn254_fn.BN_print(" b = ", b, 8)
	bn254_fn.BN_print(" n = ", n, 8)

	bn254_fn.BN_mod_add_non_const_time(r, a, b, n, 8)
	bn254_fn.BN_print(" a + b (mod n) = ", r, 8)

	bn254_fn.BN_mod_sub_non_const_time(r, a, b, n, 8)
	bn254_fn.BN_print(" a - b (mod n) = ", r, 8)

	bn254_fn.BN_mod_neg(r, a, n, 8)
	bn254_fn.BN_print(" -a (mod n) = ", r, 8)
}

func test_bn_barrett() {
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
	bn254_fn.BN_from_hex(a, 8, hex_a)
	bn254_fn.BN_from_hex(b, 8, hex_b)
	bn254_fn.BN_from_hex(p, 8, hex_p)
	bn254_fn.BN_from_hex(p_mu, 9, hex_p_mu)
	bn254_fn.BN_from_hex(n, 8, hex_n)
	bn254_fn.BN_from_hex(n_mu, 9, hex_n_mu)

	bn254_fn.BN_print(" a = ", a, 8)
	bn254_fn.BN_print(" b = ", b, 8)
	bn254_fn.BN_print(" p = ", n, 8)
	bn254_fn.BN_print(" mu(p) = ", n_mu, 9)

	bn254_fn.BN_barrett_mod_mul(r, a, b, p, p_mu, 8)
	bn254_fn.BN_print(" a * b (mod p) = ", r, 8)
	bn254_fn.BN_barrett_mod_sqr(r, a, p, p_mu, 8)
	bn254_fn.BN_print(" a^2 (mod p) = ", r, 8)
	// bn254_fn.BN_barrett_mod_exp(r, a, b, p, p_mu, 8)
	// bn254_fn.BN_print(" a^b (mod p) = ", r, 8)
	// bn254_fn.BN_barrett_mod_inv(r, a, p, p_mu, 8)
	// bn254_fn.BN_print(" a^-1 (mod p) = ", r, 8)
	bn254_fn.BN_barrett_mod_mul(r, a, b, n, n_mu, 8)
	bn254_fn.BN_print(" a * b (mod n) = ", r, 8)
	bn254_fn.BN_barrett_mod_sqr(r, a, n, n_mu, 8)
	bn254_fn.BN_print(" a^2 (mod n) = ", r, 8)
	// bn254_fn.BN_barrett_mod_exp(r, a, b, n, n_mu, 8)
	// bn254_fn.BN_print(" a^b (mod n) = ", r, 8)
}

// func test_bn_montgomery() {
// 	hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
// 	hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
// 	hex_p := "30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"
// 	hex_n := "30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001"
// 	hex_p_inv_neg := "f57a22b791888c6bd8afcbd01833da809ede7d651eca6ac987d20782e4866389"
// 	hex_p_one_sqr := "06d89f71cab8351f47ab1eff0a417ff6b5e71911d44501fbf32cfc5b538afa89"
// 	hex_n_inv_neg := "73f82f1d0d8341b2e39a9828990623916586864b4c6911b3c2e1f593efffffff"
// 	hex_n_one_sqr := "0216d0b17f4e44a58c49833d53bb808553fe3ab1e35c59e31bb8e645ae216da7"
// 	a := make([]uint32, 8)
// 	b := make([]uint32, 8)
// 	n := make([]uint32, 8)
// 	r := make([]uint32, 8)
// 	p := make([]uint32, 8)
// 	p_inv_neg := make([]uint32, 8)
// 	p_one_sqr := make([]uint32, 8)
// 	n_inv_neg := make([]uint32, 8)
// 	n_one_sqr := make([]uint32, 8)

// 	bn254_fn.BN_from_hex(a, 8, hex_a)
// 	bn254_fn.BN_from_hex(b, 8, hex_b)
// 	bn254_fn.BN_from_hex(p, 8, hex_p)
// 	bn254_fn.BN_from_hex(p_inv_neg, 8, hex_p_inv_neg)
// 	bn254_fn.BN_from_hex(p_one_sqr, 8, hex_p_one_sqr)

// 	bn254_fn.BN_print(" a = ", a, 8)
// 	bn254_fn.BN_print(" b = ", b, 8)
// 	bn254_fn.BN_print(" p = ", p, 8)

// 	bn254_fn.BN_mod_sqr_montgomery(r, a, p_one_sqr, p, p_inv_neg, 8)
// 	bn254_fn.BN_print(" a^2 (mod p) = ", r, 8)

// 	bn254_fn.BN_from_hex(n, 8, hex_n)
// 	bn254_fn.BN_from_hex(n_inv_neg, 8, hex_n_inv_neg)
// 	bn254_fn.BN_from_hex(n_one_sqr, 8, hex_n_one_sqr)

// 	bn254_fn.BN_print(" n = ", n, 8)

// 	bn254_fn.BN_mod_mul_montgomery(r, a, b, n_one_sqr, n, n_inv_neg, 8)
// 	bn254_fn.BN_print(" a * b (mod n) = ", r, 8)

// 	bn254_fn.BN_mod_sqr_montgomery(r, a, n_one_sqr, n, n_inv_neg, 8)
// 	bn254_fn.BN_print(" a^2 (mod n) = ", r, 8)
// }

func test_fn() {
	hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	r := make([]uint32, 8)
	bn254_fn.BN_from_hex(a, 8, hex_a)
	bn254_fn.BN_from_hex(b, 8, hex_b)
	bn254_fn.BN_print(" a = ", a, 8)
	bn254_fn.BN_print(" b = ", b, 8)
	bn254_fn.BN_print(" n = ", bn254_fn.BN254_n[:], 8)

	bn254_fn.FN_from_hex(a, hex_a)
	bn254_fn.FN_from_hex(b, hex_b)
	bn254_fn.FN_add(r, a, b)
	bn254_fn.FN_get_BN(r, r)
	bn254_fn.BN_print(" a + b (mod n) = ", r, 8)

	bn254_fn.FN_sub(r, a, b)
	bn254_fn.FN_get_BN(r, r)
	bn254_fn.BN_print(" a - b (mod n) = ", r, 8)

	bn254_fn.FN_sub(r, b, a)
	bn254_fn.FN_get_BN(r, r)
	bn254_fn.BN_print(" b - a (mod n) = ", r, 8)

	bn254_fn.FN_neg(r, a)
	bn254_fn.FN_get_BN(r, r)
	bn254_fn.BN_print(" -a (mod p) = ", r, 8)

	bn254_fn.FN_mul(r, a, b)
	bn254_fn.FN_get_BN(r, r)
	bn254_fn.BN_print(" a * b (mod n) = ", r, 8)

	bn254_fn.FN_sqr(r, a)
	bn254_fn.FN_get_BN(r, r)
	bn254_fn.BN_print(" a^2 (mod n) = ", r, 8)

	var a_fn bn254_fn.FN_t
	for i := 0; i < len(a); i++ {
		a_fn[i] = a[i]
	}
	var r_fn bn254_fn.FN_t
	for i := 0; i < len(r); i++ {
		r_fn[i] = r[i]
	}
	//bn254_fn.FN_inv(r_fn, a_fn)
	bn254_fn.FN_get_bn(r, r)
	bn254_fn.BN_print(" a^-1 (mod n) = ", r, 8)
}

func main() {
	//test_bn()
	//test_bn_mod()
	//test_bn_barrett()
	test_fn()
	//test_bn_montgomery()
}
