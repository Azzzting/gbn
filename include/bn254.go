package bn254_FP

import (
	"encoding/binary"
	"fmt"
	"os"
)

type FP_t [8]uint32

type Poly_T []*FN_t

type Point_t struct {
	X              FP_t
	Y              FP_t
	Z              FP_t
	Is_At_Infinity int
}

type Affine_point_t struct {
	X FP_t
	Y FP_t
}

var BN254_p = [8]uint32{
	0xd87cfd47, 0x3c208c16, 0x6871ca8d, 0x97816a91,
	0x8181585d, 0xb85045b6, 0xe131a029, 0x30644e72,
}
var BN254_n = [8]uint32{
	0xf0000001, 0x43e1f593, 0x79b97091, 0x2833e848,
	0x8181585d, 0xb85045b6, 0xe131a029, 0x30644e72,
}

var BN254_p_mont_one = []uint32{
	0xc58f0d9d, 0xd35d438d, 0xf5c70b3d, 0xa78eb28,
	0x7879462c, 0x666ea36f, 0x9a07df2f, 0xe0a77c1,
}

var BN254_p_mont_one_sqr = []uint32{
	0x538afa89, 0xf32cfc5b, 0xd44501fb, 0xb5e71911,
	0xa417ff6, 0x47ab1eff, 0xcab8351f, 0x6d89f71,
}

var BN254_n_inv_neg = []uint32{
	0xefffffff, 0xc2e1f593, 0x4c6911b3, 0x6586864b,
	0x99062391, 0xe39a9828, 0x0d8341b2, 0x73f82f1d,
}

var BN254_n_mont_one = []uint32{
	0x4ffffffb, 0xac96341c, 0x9f60cd29, 0x36fc7695,
	0x7879462e, 0x666ea36f, 0x9a07df2f, 0x0e0a77c1,
}

var BN254_n_mont_one_sqr = []uint32{
	0xae216da7, 0x1bb8e645, 0xe35c59e3, 0x53fe3ab1,
	0x53bb8085, 0x8c49833d, 0x7f4e44a5, 0x0216d0b1,
}

var BN254_p_inv_neg = []uint32{
	0xe4866389, 0x87d20782, 0x1eca6ac9, 0x9ede7d65,
	0x1833da80, 0xd8afcbd0, 0x91888c6b, 0xf57a22b7,
}

var BN254_p_mu = [9]uint32{
	0x9bf90e51, 0xf3aed8a1, 0x7cd4c086, 0xe965e176,
	0x8073013a, 0xb074a586, 0x23a04a7a, 0x4a474626,
	0x05,
}

var BN254_n_mu = [9]uint32{
	0xe1de9259, 0x20703a6b, 0x9e880ae6, 0x14485200,
	0x80730147, 0xb074a586, 0x23a04a7a, 0x4a474626,
	0x05,
}

var BN254_one = [8]uint32{1, 0, 0, 0, 0, 0, 0, 0}
var BN254_two = [8]uint32{2, 0, 0, 0, 0, 0, 0, 0}

// ues mgml or not
var UseMontgomery = false

// func
func BN_print(str string, a []uint32, k int) {
	fmt.Printf("%s", str)
	for i := k - 1; i >= 0; i-- {
		fmt.Printf("%08x", a[i])
	}
	fmt.Println()
}

func BN_set_word(r []uint32, a uint32, k int) {
	r[0] = a
	for i := k - 1; i >= 1; i-- {
		r[i] = 0
	}
}

func BN_copy(r []uint32, a []uint32, k int) {
	for i := k - 1; i >= 0; i-- {
		r[i] = a[i]
	}
}

func BN_cmp(a []uint32, b []uint32, k int) int {
	for i := k - 1; i >= 0; i-- {
		if a[i] > b[i] {
			return 1
		}
		if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

func BN_is_zero(a []uint32, k int) (result int) {
	for i := k - 1; i >= 0; i-- {
		if a[i] != 0 {
			return 0
		}
	}
	return 1
}

func BN_is_one(a []uint32, k int) (result int) {
	if a[0] != 1 {
		return 0
	}
	for i := k - 1; i >= 1; i-- {
		if a[i] == 1 {
			return 0
		}
	}
	return 1
}

func BN_add(r []uint32, a []uint32, b []uint32, k int) int {
	var w uint64 = 0
	for i := 0; i < k; i++ {
		w += uint64(a[i]) + uint64(b[i])
		r[i] = uint32(w & 0xffffffff)
		w >>= 32
	}
	return int(w)
}

func BN_sub(r []uint32, a []uint32, b []uint32, k int) int {
	var w int64 = 0
	for i := 0; i < k; i++ {
		w += int64(a[i]) - int64(b[i])
		r[i] = uint32(w & 0xffffffff)
		w >>= 32
	}
	return int(w)
}

func BN_mul(r []uint32, a []uint32, b []uint32, k int) {

	for i := 0; i < k; i++ {
		r[i] = 0
	}
	for i := 0; i < k; i++ {
		var w uint64 = 0
		for j := 0; j < k; j++ {
			w += uint64((r)[i+j]) + uint64(a[i])*uint64(b[j])
			r[i+j] = uint32(w & 0xffffffff)
			w >>= 32
		}
		r[i+k] = uint32(w)
	}
}

func BN_mul_lo(r []uint32, a []uint32, b []uint32, k int) {
	for i := 0; i < k; i++ {
		r[i] = 0
	}
	for i := 0; i < k; i++ {
		var w uint64 = 0
		for j := 0; j < k-i; j++ {
			w += uint64((r)[i+j]) + uint64(a[i])*uint64(b[j])
			r[i+j] = uint32(w & 0xffffffff)
			w >>= 32
		}
	}
}

func Uint32_to_bytes(a uint32, out []uint8) {
	out[0] = uint8((a >> 24) & 0xff)
	out[1] = uint8((a >> 16) & 0xff)
	out[2] = uint8((a >> 8) & 0xff)
	out[3] = uint8(a & 0xff)
}

func Uint32_from_bytes(in []uint8) uint32 {
	return (uint32(in[0]) << 24) |
		(uint32(in[1]) << 16) |
		(uint32(in[2]) << 8) |
		uint32(in[3])
}

func BN_to_bytes(a []uint32, k int, out []uint8) {
	j := 0
	for i := k - 1; i >= 0; i-- {
		Uint32_to_bytes(a[i], out[j:j+4])
		j = j + 4
	}
}

func BN_from_bytes(a []uint32, k int, in []uint8) {
	offset := 0
	for i := k - 1; i >= 0; i-- {
		in0 := in[offset : offset+4]
		a[i] = Uint32_from_bytes(in0)
		offset = offset + 4
	}
}

func BN_to_hex(a []uint32, k int, out []string) {
	for i := k - 1; i >= 0; i-- {
		out[i] = fmt.Sprintf("%08x", a[i])
	}
}

func Hexchar2int(c rune) int {
	switch {
	case '0' <= c && c <= '9':
		return int(c - '0')
	case 'a' <= c && c <= 'f':
		return int(c - 'a' + 10)
	case 'A' <= c && c <= 'F':
		return int(c - 'A' + 10)
	default:
		return -1
	}
}

func Hex2bin(in string, inlen int, out []uint8) int {
	if inlen%2 != 0 {
		return -1
	}
	for i := 0; i < inlen; i += 2 {
		c1 := Hexchar2int(rune(in[i]))
		c2 := Hexchar2int(rune(in[i+1]))
		if c1 < 0 || c2 < 0 {
			return -1
		}

		out[i/2] = uint8(c1<<4 | c2)
	}
	return 1
}

func BN_from_hex(r []uint32, k int, hex string) int {
	buf := make([]uint8, 4*k)
	if Hex2bin(hex, k*8, buf) < 0 {
		return -1
	}
	BN_from_bytes(r, k, buf)
	return 1
}
func BN_equ_hex(a []uint32, k int, hex string) bool {
	a_ := make([]uint32, k)
	BN_from_hex(a_, k, hex)
	return BN_cmp(a, a_, k) == 0
}

func BN_rand(r []uint32, k int) int {
	file, err := os.Open("/dev/urandom")
	if err != nil {
		return -1
	}
	defer file.Close()

	data := make([]byte, 4*k)
	_, err = file.Read(data)
	if err != nil {
		return -1
	}

	for i := 0; i < k; i++ {
		(r)[i] = binary.BigEndian.Uint32(data[i*4 : (i+1)*4])
	}

	return 1
}

func BN_rand_range(r []uint32, rng []uint32, k int) int {
	for {
		BN_rand(r, k)
		if BN_cmp(r, rng, k) < 0 {
			break
		}
	}
	return 1
}

func FN_exp(r FN_t, a FN_t, e []uint32, elen int) {
	t := make([]uint32, 8)
	var w uint32
	FN_set_one(t)
	for i := elen - 1; i >= 0; i-- {
		w = e[i]
		for j := 0; j < 32; j++ {
			FN_sqr(t, t)
			if (w & 0x80000000) != 0 {
				FN_mul(t, t, a[:])
			}
			w <<= 1
		}
	}
	FN_copy(r[:], t)
}
func FP_exp(r *FP_t, a FP_t, e []uint32, elen int) {
	t := make([]uint32, 8)
	var w uint32
	FP_set_one(t)
	for i := elen - 1; i >= 0; i-- {
		w = e[i]
		for j := 0; j < 32; j++ {
			FP_sqr(t, t)
			if (w & 0x80000000) != 0 {
				FP_mul(t, t, a[:])
			}
			w <<= 1
		}
	}
	FP_copy(r[:], t)
}
func BN_mod_add_non_const_time(r []uint32, a []uint32, b []uint32, p []uint32, k int) {
	BN_add(r, a, b, k)
	if BN_cmp(r, p, k) >= 0 {
		BN_sub(r, r, p, k)
	}
}

func BN_mod_sub_const_time(r []uint32, a []uint32, b []uint32, p []uint32, k int) {
	r2 := make([]uint32, 2*k)
	r3 := r2[k:]
	w := BN_sub(r2, a, b, k)
	BN_sub(r3, b, a, k)
	copy(r2[k:], r3)
	BN_sub(r2, p, r2, k)
	r4 := r2[k+k*w:]
	r = make([]uint32, k)
	copy(r, r4)
}
func BN_mod_sub_non_const_time(r []uint32, a []uint32, b []uint32, p []uint32, k int) {
	if BN_cmp(a, b, k) >= 0 {
		BN_sub(r, a, b, k)
	} else {
		BN_sub(r, b, a, k)
		BN_sub(r, p, r, k)
	}
}

//func BN_mod_add__time(r *uint32, a *uint32, b *uint32, p *uint32, k int)
//func BN_mod_sub__time(r *uint32, a *uint32, b *uint32, p *uint32, k int)

func BN_mod_neg(r []uint32, a []uint32, p []uint32, k int) {
	BN_sub(r, p, a, k)
}

func BN_reduce_once(r []uint32, a []uint32, b []uint32, k int) {
	r2 := make([]uint32, 2*k)

	// 将 a 复制到 r2 的前半部分
	copy(r2[:k], a)
	r3 := r2[k:]
	// 计算 r1 = a - b
	w := BN_sub(r3, a, b, k)

	// 将 r1 复制到 r
	r4 := r2[k+k*w:]
	r = make([]uint32, k)
	copy(r, r4)
}

func BN_barrett_mod_mul(r []uint32, a []uint32, b []uint32, p []uint32, u []uint32, k int) {
	z := make([]uint32, 2*k)
	q := make([]uint32, 2*(k+1))
	p_ := make([]uint32, k+1)
	t_ := make([]uint32, k+1)
	r_ := make([]uint32, k+1)

	// 复制 p 的内容到 p_
	for i := 0; i < k; i++ {
		p_[i] = (p)[i]
	}
	p_[k] = 0

	// 计算 z = a * b
	BN_mul(z, a, b, k)

	// 计算 q = z * u
	z1 := z[k-1:]
	BN_mul(q, z1, u, k+1)

	// 计算 t_ = q * p_
	q1 := q[k+1:]
	BN_mul_lo(t_, q1, p_, k+1)

	// 计算 r_ = z - t_
	BN_sub(r_, z, t_, k+1)

	// 通过 BN_reduce_once 函数进行约简操作
	BN_reduce_once(r_, r_, p_, k+1)
	BN_reduce_once(r_, r_, p_, k)

	// 将 r_ 的内容复制到 r 中
	for i := 0; i < k; i++ {
		(r)[i] = r_[i]
	}
}

func BN_barrett_mod_sqr(r []uint32, a []uint32, p []uint32, u []uint32, k int) {
	BN_barrett_mod_mul(r, a, a, p, u, k)
}

func BN_barrett_mod_exp(r []uint32, a []uint32, e []uint32, p []uint32, u []uint32, k int) {
	var t = make([]uint32, k)
	w := uint32(0)

	// t = 1
	BN_set_one(t, k)

	for i := k - 1; i >= 0; i-- {
		w = e[i]
		for j := 0; j < 32; j++ {
			BN_barrett_mod_sqr(t, t, p, u, k)
			if w&0x80000000 != 0 {
				BN_barrett_mod_mul(t, t, a, p, u, k)
			}
			w <<= 1
		}
	}
	BN_copy(r, t, k)
}

func BN_barrett_mod_inv(r []uint32, a []uint32, p []uint32, u []uint32, k int) {
	e := make([]uint32, k)
	e[0] = 2
	for i := 1; i < k; i++ {
		e[i] = 0
	}
	BN_sub(e, p, e, k)
	BN_barrett_mod_exp(r, a, e, p, u, k)
}

func BN_mont_mod_mul(r []uint32, a []uint32, b []uint32, p []uint32, p_inv_neg []uint32, k int) {
	z := make([]uint32, k*2)
	c := make([]uint32, k*2)
	t := make([]uint32, k)
	BN_mul(z, a, b, k)
	BN_mul_lo(t, z, p_inv_neg, k)
	BN_mul(c, t, p, k)
	BN_add(c, c, z, k*2)
	c1 := c[k:]
	if BN_cmp(c1, p, k) >= 0 {
		BN_sub(c1, c1, p, k)
	}
	copy(c[k:], c1)
	for i := 0; i < k; i++ {
		(r)[i] = c[k+i]
	}
}

func BN_mont_mod_sqr(r []uint32, a []uint32, p []uint32, p_inv_neg []uint32, k int) {
	BN_mont_mod_mul(r, a, a, p, p_inv_neg, k)
}

func BN_mont_mod_exp(r []uint32, a []uint32, e []uint32, p []uint32, p_inv_neg []uint32, k int) {
	var t []uint32
	BN_set_one(t, k)

	for i := k - 1; i >= 0; i-- {
		w := (e)[i]
		for j := 0; j < 32; j++ {
			BN_mont_mod_sqr(t, t, p, p_inv_neg, k)
			if w&0x80000000 != 0 {
				BN_mont_mod_mul(t, t, a, p, p_inv_neg, k)
			}
			w <<= 1
		}
	}
	BN_copy(r, t, k)
}

func BN_mont_mod_inv(r []uint32, a []uint32, p []uint32, p_inv_neg []uint32, k int) {
	e := make([]uint32, k)
	e[0] = 2
	for i := 1; i < k; i++ {
		e[i] = 0
	}
	BN_sub(e, p, e, k)
	BN_mont_mod_exp(r, a, e, p, p_inv_neg, k)
}

func BN_mont_set(r []uint32, a []uint32, R_sqr []uint32, p []uint32, p_inv_neg []uint32, k int) {
	BN_mont_mod_mul(r, a, R_sqr, p, p_inv_neg, k)
}

func BN_mont_get(r []uint32, a []uint32, p []uint32, p_inv_neg []uint32, k int) {
	one := BN254_one[:]
	BN_set_one(one, k)
	BN_mont_mod_mul(r, a, one, p, p_inv_neg, k)
}

// for testing
func BN_mod_mul_montgomery(r []uint32, a []uint32, b []uint32, R_sqr []uint32, p []uint32, p_inv_neg []uint32, k int) {
	a_ := make([]uint32, k)
	b_ := make([]uint32, k)
	r_ := make([]uint32, k)
	BN_mont_set(a_, a, R_sqr, p, p_inv_neg, k)
	BN_mont_set(b_, b, R_sqr, p, p_inv_neg, k)
	BN_mont_mod_mul(r_, a_, b_, p, p_inv_neg, k)
	BN_mont_get(r, r_, p, p_inv_neg, k)
}

// also testing
func BN_mod_sqr_montgomery(r []uint32, a []uint32, R_sqr []uint32, p []uint32, p_inv_neg []uint32, k int) {
	a_ := make([]uint32, k)
	r_ := make([]uint32, k)
	BN_mont_set(a_, a, R_sqr, p, p_inv_neg, k)
	BN_mont_mod_sqr(r_, a_, p, p_inv_neg, k)
	BN_mont_get(r, r_, p, p_inv_neg, k)
}

type FN_t [8]uint32

// define
func BN_set_zero(r []uint32, k int) {
	if k > 0 {
		BN_set_word(r, 0, k)
	}
}

func BN_set_one(r []uint32, k int) {
	if k > 0 {
		BN_set_word(r, 1, k)
	}
}
func BN_mod_add(r []uint32, a []uint32, b []uint32, p []uint32, k int) {
	//BN_mod_add__time(r,a,b,p,k);
	BN_mod_add_non_const_time(r, a, b, p, k)
}
func BN_mod_sub(r []uint32, a []uint32, b []uint32, p []uint32, k int) {
	//BN_mod_sub__time(r,a,b,p,k);
	BN_mod_sub_non_const_time(r, a, b, p, k)
}
func FN_print(str string, a []uint32) {
	BN_print(str, a, 8)
}
func FP_print(str string, a []uint32) {
	BN_print(str, a, 8)
}
func FN_rand(r []uint32, a []uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	BN_rand_range(r, rng, 8)
}
func FP_rand(r []uint32, a []uint32) {
	rng := make([]uint32, len(BN254_p))
	copy(rng, BN254_p[:])
	BN_rand_range(r, rng, 8)
}
func FN_copy(r []uint32, a []uint32) {
	BN_copy(r, a, 8)
}
func FP_copy(r []uint32, a []uint32) {
	BN_copy(r, a, 8)
}
func FN_set_word(r []uint32, a uint32) {
	BN_set_word(r, a, 8)
}
func FN_set_zero(r []uint32, k int) {
	BN_set_zero(r, 8)
}
func FP_set_zero(r []uint32) {
	BN_set_zero(r, 8)
}
func FN_is_zero(a []uint32) {
	BN_is_zero(a, 8)
}
func FP_is_zero(a []uint32) bool {
	return BN_is_zero(a, 8) == 1
}
func FN_add(r []uint32, a []uint32, b []uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	BN_mod_add(r, a, b, rng, 8)
}
func FP_add(r []uint32, a []uint32, b []uint32) {
	rng := make([]uint32, len(BN254_p))
	copy(rng, BN254_p[:])
	BN_mod_add(r, a, b, rng, 8)
}
func FN_sub(r []uint32, a []uint32, b []uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	BN_mod_sub(r, a, b, rng, 8)
}
func FP_sub(r []uint32, a []uint32, b []uint32) {
	rng := make([]uint32, len(BN254_p))
	copy(rng, BN254_p[:])
	BN_mod_sub(r, a, b, rng, 8)
}
func FN_neg(r []uint32, a []uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	BN_mod_neg(r, a, rng, 8)
}
func FP_neg(r []uint32, a []uint32) {
	rng := make([]uint32, len(BN254_p))
	copy(rng, BN254_p[:])
	BN_mod_neg(r, a, rng, 8)
}
func FN_dbl(r []uint32, a []uint32) {
	FN_add(r, a, a)
}
func FP_dbl(r []uint32, a []uint32) {
	FP_add(r, a, a)
}
func FN_to_bytes(a []uint32, out []uint8) {
	BN_to_bytes(a, 8, out)
}
func FP_to_bytes(a []uint32, out []uint8) {
	BN_to_bytes(a, 8, out)
}
func FN_from_bytes(a []uint32, in []uint8) {
	BN_from_bytes(a, 8, in)
}
func FP_from_bytes(a []uint32, in []uint8) {
	BN_from_bytes(a, 8, in)
}
func FN_set_BN(r []uint32, a []uint32) {
	FN_copy(r, a)
}
func FP_set_BN(r []uint32, a []uint32) {
	rng := make([]uint32, len(BN254_p))
	copy(rng, BN254_p[:])
	rng1 := make([]uint32, len(BN254_p_mont_one_sqr))
	copy(rng1, BN254_p_mont_one_sqr[:])
	rng2 := make([]uint32, len(BN254_p_inv_neg))
	copy(rng2, BN254_p_inv_neg[:])
	BN_mont_mod_mul(r, a, rng1, rng, rng2, 8)
}
func FN_set_one(r []uint32) {
	BN_set_one(r, 8)
}
func FP_set_one(r []uint32) {
	rng := make([]uint32, len(BN254_p_mont_one))
	copy(rng, BN254_p_mont_one[:])
	BN_copy(r, rng, 8)
}
func FN_to_hex(a []uint32, out []string) {
	BN_to_hex(a, 8, out)
}
func FN_from_hex(r []uint32, hex string) {
	BN_from_hex(r, 8, hex)
}
func FP_from_hex(r []uint32, hex string) {
	BN_from_hex(r, 8, hex)
	FP_set_BN(r, r)
}
func FN_get_BN(r []uint32, a []uint32) {
	FN_copy(a, r)
}
func FP_get_BN(a []uint32, r []uint32) {
	rng := make([]uint32, len(BN254_p))
	copy(rng, BN254_p[:])
	rng1 := make([]uint32, len(BN254_p_inv_neg))
	copy(rng1, BN254_p_inv_neg[:])
	BN_mont_get(r, a, rng, rng1, 8)
}
func FN_is_one(a []uint32) {
	BN_is_one(a, 8)
}
func FP_is_one(a []uint32) bool {
	rng := make([]uint32, len(BN254_p_mont_one))
	copy(rng, BN254_p_mont_one[:])
	return BN_cmp(a, rng, 8) == 0
}
func FN_mul(r []uint32, a []uint32, b []uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	rng1 := make([]uint32, len(BN254_n_mu))
	copy(rng1, BN254_n_mu[:])
	BN_barrett_mod_mul(r, a, b, rng, rng1, 8)
}
func FP_mul(r []uint32, a []uint32, b []uint32) {
	rng := make([]uint32, len(BN254_p))
	copy(rng, BN254_p[:])
	rng1 := make([]uint32, len(BN254_p_inv_neg))
	copy(rng1, BN254_p_inv_neg[:])
	BN_mont_mod_mul(r, a, b, rng, rng1, 8)
}
func FN_sqr(r []uint32, a []uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	rng1 := make([]uint32, len(BN254_n_mu))
	copy(rng1, BN254_n_mu[:])
	BN_barrett_mod_sqr(r, a, rng, rng1, 8)
}
func FP_sqr(r []uint32, a []uint32) {
	rng := make([]uint32, len(BN254_p))
	copy(rng, BN254_p[:])
	rng1 := make([]uint32, len(BN254_p_inv_neg))
	copy(rng1, BN254_p_inv_neg[:])
	BN_mont_mod_sqr(r, a, rng, rng1, 8)
}

func FN_inv(r FN_t, a FN_t) {
	e := make([]uint32, 8)
	e[0] = 2
	for i := 1; i < 8; i++ {
		e[i] = 0
	}

	BN_sub(e, BN254_n[:], e, 8)
	FN_exp(r, a, e, 8)
}

func FP_inv(r *FP_t, a FP_t) {
	e := make([]uint32, 8)
	e[0] = 2
	for i := 1; i < 8; i++ {
		e[i] = 0
	}

	BN_sub(e, BN254_p[:], e, 8)
	FP_exp(r, a, e, 8)
}

func FN_get_bn(a []uint32, r []uint32) {
	FN_copy(r, a)
}

func FP_tri(r *FP_t, a *FP_t) {
	var t FP_t
	FP_dbl(t[:], a[:])
	FP_add(r[:], t[:], a[:])
}

func Point_from_hex(R *Point_t, hex string) {
	FP_from_hex(R.X[:], hex)
	FP_from_hex(R.Y[:], hex[64:])
	FP_set_one(R.Z[:])
	R.Is_At_Infinity = 0
}

func Point_get_xy(P *Point_t, x []uint32, y []uint32) {
	var Z_inv FP_t

	FP_inv(&Z_inv, P.Z)
	FP_mul(y, P.Y[:], Z_inv[:])
	FP_sqr(Z_inv[:], Z_inv[:])
	FP_mul(x, P.X[:], Z_inv[:])
	FP_get_BN(x, x)
	FP_mul(y, y, Z_inv[:])
	FP_get_BN(y, y)
}

func Point_print(str string, P *Point_t) {
	var x [8]uint32
	var y [8]uint32

	Point_get_xy(P, x[:], y[:])

	fmt.Printf("%s", str)
	for i := 7; i >= 0; i-- {
		fmt.Printf("%08x", x[i])
	}
	fmt.Print(" ")
	for i := 7; i >= 0; i-- {
		fmt.Printf("%08x", y[i])
	}
	fmt.Println()
}

func Point_dbl(R *Point_t, P *Point_t) {
	var T_0 FP_t
	var T_1 FP_t
	var T_2 FP_t
	var T_3 FP_t
	var T_4 FP_t

	FP_sqr(T_0[:], P.X[:])
	FP_tri(&T_0, &T_0)
	FP_sqr(T_1[:], T_0[:])
	FP_sqr(T_2[:], P.Y[:])
	FP_mul(T_3[:], P.X[:], T_2[:])
	FP_dbl(T_3[:], T_3[:])
	FP_dbl(T_3[:], T_3[:])
	FP_dbl(T_4[:], T_3[:])
	FP_sub(T_1[:], T_1[:], T_4[:])
	FP_sub(T_3[:], T_3[:], T_1[:])
	FP_mul(T_0[:], T_0[:], T_3[:])
	FP_dbl(T_2[:], T_2[:])
	FP_sqr(T_2[:], T_2[:])
	FP_dbl(T_2[:], T_2[:])
	FP_sub(T_0[:], T_0[:], T_2[:])
	FP_mul(T_2[:], P.Y[:], P.Z[:])
	FP_dbl(T_2[:], T_2[:])

	FP_copy(R.X[:], T_1[:])
	FP_copy(R.Y[:], T_0[:])
	FP_copy(R.Z[:], T_2[:])
	R.Is_At_Infinity = P.Is_At_Infinity
}

func Point_copy(R *Point_t, P *Point_t) {
	FP_copy(R.X[:], P.X[:])
	FP_copy(R.Y[:], P.Y[:])
	FP_copy(R.Z[:], P.Z[:])
	R.Is_At_Infinity = P.Is_At_Infinity
}

func Point_set_infinity(R *Point_t) {
	FP_set_one(R.X[:])
	FP_set_one(R.Y[:])
	FP_set_zero(R.Z[:])
	R.Is_At_Infinity = 1
}

func Point_add_jacobian(R *Point_t, P *Point_t, Q *Point_t) {
	var T_1 FP_t
	var T_2 FP_t
	var T_3 FP_t
	var T_4 FP_t
	var T_5 FP_t
	var T_6 FP_t
	var T_7 FP_t
	var T_8 FP_t

	if P.Is_At_Infinity != 0 {
		Point_copy(R, Q)
		return
	}

	FP_sqr(T_1[:], P.Z[:])         // T_1 = Z_1^2
	FP_sqr(T_2[:], Q.Z[:])         // T_2 = Z_2^2
	FP_mul(T_3[:], Q.X[:], T_1[:]) // T_3 = X_2 * Z_1^2
	FP_mul(T_4[:], P.X[:], T_2[:]) // T_4 = X_1 * Z_2^2
	FP_add(T_5[:], T_3[:], T_4[:]) // T_5 = X_2 * Z_1^2 + X_1 * Z_2^2 = C
	FP_sub(T_3[:], T_3[:], T_4[:]) // T_3 = X_2 * Z_1^2 - X_1 * Z_2^2 = B
	FP_mul(T_1[:], T_1[:], P.Z[:]) // T_1 = Z_1^3
	FP_mul(T_1[:], T_1[:], Q.Y[:]) // T_1 = Y_2 * Z_1^3
	FP_mul(T_2[:], T_2[:], Q.Z[:]) // T_2 = Z_2^3
	FP_mul(T_2[:], T_2[:], P.Y[:]) // T_2 = Y_1 * Z_2^3
	FP_add(T_6[:], T_1[:], T_2[:]) // T_6 = Y_2 * Z_1^3 + Y_1 * Z_2^3 = D
	FP_sub(T_1[:], T_1[:], T_2[:]) // T_1 = Y_2 * Z_1^3 - Y_1 * Z_2^3 = A

	if FP_is_zero(T_1[:]) && FP_is_zero(T_3[:]) {
		Point_dbl(R, P)
		return
	}

	if FP_is_one(T_1[:]) && FP_is_zero(T_6[:]) {
		Point_set_infinity(R)
		return
	}

	FP_sqr(T_6[:], T_1[:])         // T_6 = A^2
	FP_mul(T_7[:], T_3[:], P.Z[:]) // T_7 = B * Z_1
	FP_mul(T_7[:], T_7[:], Q.Z[:]) // T_7 = B * Z_1 * Z_2 = Z_3
	FP_sqr(T_8[:], T_3[:])         // T_8 = B^2
	FP_mul(T_5[:], T_5[:], T_8[:]) // T_5 = B^2 * C
	FP_mul(T_3[:], T_3[:], T_8[:]) // T_3 = B^3
	FP_mul(T_4[:], T_4[:], T_8[:]) // T_4 = B^2 * X_1 * Z_2^2
	FP_sub(T_6[:], T_6[:], T_5[:]) // T_6 = A^2 - B^2 * C = X_3
	FP_sub(T_4[:], T_4[:], T_6[:]) // T_4 = B^2 * X_1 * Z_2^2 - X_3
	FP_mul(T_1[:], T_1[:], T_4[:]) // T_1 = A * (B^2 * X_1 * Z_2^2 - X_3)
	FP_mul(T_2[:], T_2[:], T_3[:]) // T_2 = B^3 * Y_1 * Z_1^3
	FP_sub(T_1[:], T_1[:], T_2[:]) // T_1 = A * (B^2 * X_1 * Z_2^2 - X_3) - B^3 * Y_1 * Z_1^3 = Y_3

	FP_copy(R.X[:], T_6[:])
	FP_copy(R.Y[:], T_1[:])
	FP_copy(R.Z[:], T_7[:])
	R.Is_At_Infinity = 0
}

func Point_add(R *Point_t, P *Point_t, Q *Point_t) {
	Point_add_jacobian(R, P, Q)
}

func Point_add_affine(R *Point_t, P *Point_t, Q *Affine_point_t) {
	var T_0 FP_t
	var T_1 FP_t
	var T_2 FP_t
	var T_3 FP_t
	var T_4 FP_t
	var T_5 FP_t

	if P.Is_At_Infinity != 0 {
		FP_copy(R.X[:], Q.X[:])
		FP_copy(R.Y[:], Q.Y[:])
		FP_set_one(R.Z[:])
		R.Is_At_Infinity = 0
		return
	}

	FP_sqr(T_0[:], P.Z[:])
	FP_mul(T_1[:], T_0[:], P.Z[:])
	FP_mul(T_1[:], Q.Y[:], T_1[:])
	FP_sub(T_1[:], T_1[:], P.Y[:])
	FP_mul(T_0[:], Q.X[:], T_0[:])
	FP_sub(T_2[:], T_0[:], P.X[:])
	FP_add(T_0[:], T_0[:], P.X[:])

	FP_sqr(T_3[:], T_2[:])
	FP_mul(T_4[:], T_3[:], T_0[:])
	FP_sqr(T_5[:], T_1[:])
	FP_sub(T_5[:], T_5[:], T_4[:])
	FP_mul(T_4[:], P.X[:], T_3[:])
	FP_sub(T_4[:], T_4[:], T_5[:])
	FP_mul(T_4[:], T_1[:], T_4[:])
	FP_mul(T_3[:], T_3[:], T_2[:])
	FP_mul(T_3[:], P.Y[:], T_3[:])
	FP_sub(T_4[:], T_4[:], T_3[:])
	FP_mul(T_3[:], T_2[:], P.Z[:])

	FP_copy(R.X[:], T_5[:])
	FP_copy(R.Y[:], T_4[:])
	FP_copy(R.Z[:], T_3[:])
	R.Is_At_Infinity = 0

	if FP_is_zero(T_3[:]) {
		Point_set_infinity(R)
	}
}

func Point_mul_affine_non_const_time(R *Point_t, a []uint32, P *Affine_point_t) {
	var bits uint32
	Point_set_infinity(R)

	bits = a[7] << 2

	for nbits := 30; nbits > 0; nbits-- {
		Point_dbl(R, R)
		if bits&0x80000000 != 0 {
			Point_add_affine(R, R, P)
		}
		bits <<= 1
	}

	for k := 6; k >= 0; k-- {
		bits = a[k]
		for nbits := 32; nbits > 0; nbits-- {
			Point_dbl(R, R)
			if bits&0x80000000 != 0 {
				Point_add_affine(R, R, P)
			}
			bits <<= 1
		}
	}
}

func Point_mul_affine_const_time(R *Point_t, a []uint32, P *Affine_point_t) {
	var bits uint32

	Point_set_infinity(R)

	bits = a[7] << 2

	for nbits := 30; nbits > 0; nbits-- {
		var Rs [2]Point_t
		do_add := (bits & 0x80000000) >> 31
		Point_dbl(&Rs[0], R)
		Point_add_affine(&Rs[1], &Rs[0], P)
		Point_copy(R, &Rs[do_add])
		bits <<= 1
	}

	for k := 6; k >= 0; k-- {
		bits = a[k]
		for nbits := 32; nbits > 0; nbits-- {
			var Rs [2]Point_t
			do_add := (bits & 0x80000000) >> 31
			Point_dbl(&Rs[0], R)
			Point_add_affine(&Rs[1], &Rs[0], P)
			Point_copy(R, &Rs[do_add])
			bits <<= 1
		}
	}
}

func Point_mul_affine(R *Point_t, a []uint32, P *Affine_point_t) {
	Point_mul_affine_non_const_time(R, a, P)
}

func Point_mul_generator(R *Point_t, a []uint32) {
	var BN254_G1 = Point_t{
		X: [8]uint32{0xc58f0d9d, 0xd35d438d, 0xf5c70b3d, 0x0a78eb28,
			0x7879462c, 0x666ea36f, 0x9a07df2f, 0x0e0a77c1},
		Y: [8]uint32{0x8b1e1b3a, 0xa6ba871b, 0xeb8e167b, 0x14f1d651,
			0xf0f28c58, 0xccdd46de, 0x340fbe5e, 0x1c14ef83},
		Z: [8]uint32{0xc58f0d9d, 0xd35d438d, 0xf5c70b3d, 0x0a78eb28,
			0x7879462c, 0x666ea36f, 0x9a07df2f, 0x0e0a77c1},
		Is_At_Infinity: 0,
	}

	affineBN254_G1 := Affine_point_t{
		X: BN254_G1.X,
		Y: BN254_G1.Y,
	}

	Point_mul_affine(R, a[:], &affineBN254_G1)
}

func Point_multi_mul_affine_pre_compute(P []*Point_t, T []*Point_t) {
	for i := 0; i < 255; i++ {
		iv := uint8(i + 1)
		Point_set_infinity(T[i])
		for k := 0; k < 8; k++ {
			if iv&0x01 != 0 {
				Point_add(T[i], T[i], P[k])
			}
			iv >>= 1
		}
	}
}

func Point_multi_mul_affine_with_pre_compute(R *Point_t, a [8][8]uint32, T []*Point_t) {
	var bits [8]uint32

	Point_set_infinity(R)

	for i := 0; i < 8; i++ {
		bits[i] = (a[i][7]) << 2
	}
	for nbits := 30; nbits > 0; nbits-- {
		u := uint8(0)
		for i := 7; i >= 0; i-- {
			u = u | (uint8((uint32(bits[i])&0x80000000)>>31) << uint(i))
		}
		Point_dbl(R, R)
		if u != 0 {
			Point_add(R, R, T[u-1])
		}
		for i := 0; i < 8; i++ {
			bits[i] <<= 1
		}
	}

	for nwords := 6; nwords >= 0; nwords-- {
		for i := 0; i < 8; i++ {
			bits[i] = a[i][nwords]
		}
		for nbits := 32; nbits > 0; nbits-- {
			u := uint8(0)
			for i := 7; i >= 0; i-- {
				u = u | (uint8((uint32(bits[i])&0x80000000)>>31) << uint(i))
			}
			Point_dbl(R, R)
			if u != 0 {
				Point_add(R, R, T[u-1])
			}
			for i := 0; i < 8; i++ {
				bits[i] <<= 1
			}

		}
	}
}

func Poly_New(len int) Poly_T {
	ret := make(Poly_T, len)
	for i := range ret {
		ret[i] = new(FN_t)
	}
	return ret
}

func Poly_print(str string, a Poly_T, alen int) {
	fmt.Printf("%s", str)
	for i := 0; i < alen; i++ {
		fmt.Printf("  %d", i)
		BN_print(" ", (*a[i])[:], 8)
	}
	fmt.Println()
}

func Poly_add(r Poly_T, rlen *int, a Poly_T, alen int, b Poly_T, blen int) {
	fa := a
	fb := b
	falen := alen
	fblen := blen

	if alen > blen {
		fa = b
		fb = a
		falen = blen
		fblen = alen
	}

	for i := 0; i < falen; i++ {
		FN_add((*r[i])[:], (*fa[i])[:], (*fb[i])[:])
	}
	for i := falen; i < fblen; i++ {
		FN_copy((*r[i])[:], (*fb[i])[:])
	}

	*rlen = fblen
}

func Poly_sub(r Poly_T, rlen *int, a Poly_T, alen int, b Poly_T, blen int) {

	if alen >= blen {
		for i := 0; i < blen; i++ {
			FN_sub((*r[i])[:], (*a[i])[:], (*b[i])[:])
		}
		for i := blen; i < alen; i++ {
			FN_copy((*r[i])[:], (*a[i])[:])
		}
		*rlen = alen
	} else {
		for i := 0; i < alen; i++ {
			FN_sub((*r[i])[:], (*a[i])[:], (*b[i])[:])
		}
		for i := alen; i < blen; i++ {
			FN_neg((*r[i])[:], (*b[i])[:])
		}
		*rlen = blen
	}
}

func Poly_copy(r Poly_T, rlen *int, a Poly_T, alen int) {
	for i := 0; i < alen; i++ {
		FN_copy((*r[i])[:], (*a[i])[:])
	}
	*rlen = alen
}

func Poly_mul(r Poly_T, rlen *int, a Poly_T, alen int, b Poly_T, blen int) {
	var j int
	var t FN_t
	tmp := Poly_New(alen + blen - 1)

	for j := 0; j < blen; j++ {
		FN_mul((*tmp[j])[:], (*a[0])[:], (*b[j])[:])
	}

	for i := 1; i < alen; i++ {
		for j := 0; j < blen-1; j++ {
			FN_mul(t[:], (*a[i])[:], (*b[j])[:])
			FN_add((*tmp[i+j])[:], (*tmp[i+j])[:], t[:])
		}
		FN_mul((*tmp[i+j])[:], (*a[i])[:], (*b[j])[:])
	}
	*rlen = alen + blen - 1
	Poly_copy(r, rlen, tmp, alen+blen-1)
}

func Poly_add_scalar(r Poly_T, rlen *int, a Poly_T, alen int, scalar FN_t) {
	FN_add((*r[0])[:], (*a[0])[:], scalar[:])
	for i := 1; i < alen; i++ {
		FN_copy((*r[i])[:], (*a[i])[:])
	}
	*rlen = alen
}

func Poly_sub_scalar(r Poly_T, rlen *int, a Poly_T, alen int, scalar FN_t) {
	FN_sub((*r[0])[:], (*a[0])[:], scalar[:])
	for i := 1; i < alen; i++ {
		FN_copy((*r[i])[:], (*a[i])[:])
	}
	*rlen = alen
}

func Poly_mul_scalar(r Poly_T, rlen *int, a Poly_T, alen int, scalar FN_t) {
	for i := 0; i < alen; i++ {
		FN_mul((*r[i])[:], (*a[i])[:], scalar[:])
	}
	*rlen = alen
}

func Poly_mul_vector(r Poly_T, rlen *int, a Poly_T, alen int, vec []*FN_t, veclen int) {
	for i := 0; i < alen; i++ {
		FN_mul((*r[i])[:], (*a[i])[:], vec[i%veclen][:])
	}
	*rlen = alen
}

func Poly_add_blind(a Poly_T, rlen *int, b Poly_T, blen int) {
	for i := 0; i < blen; i++ {
		FN_sub((*a[i])[:], (*a[i])[:], (*b[i])[:])
		FN_copy((*a[*rlen+i])[:], (*b[i])[:])
	}
	(*rlen) += blen
}

func Poly_div_x_sub_scalar(a Poly_T, alen *int, scalar FN_t) {
	var t FN_t
	for i := *alen - 2; i >= 0; i-- {
		FN_mul(t[:], (*a[i+1])[:], scalar[:])
		FN_add((*a[i])[:], (*a[i])[:], t[:])
	}
	for i := 0; i < *alen-1; i++ {
		FN_copy((*a[i])[:], (*a[i+1])[:])
	}
	(*alen)--
}

func Poly_div_ZH(r Poly_T, rlen *int, a Poly_T, alen int, n int) {
	for i := 0; i < n; i++ {
		FN_neg((*r[i])[:], (*a[i])[:])
	}
	for i := n; i < alen; i++ {
		FN_sub((*r[i])[:], (*r[i-n])[:], (*a[i])[:])
	}
	*rlen = alen - n
}

func Poly_eval(r FN_t, a []*FN_t, alen int, x FN_t) {
	FN_copy(r[:], a[alen-1][:])

	for alen -= 2; alen >= 0; alen-- {
		FN_mul(r[:], r[:], x[:])
		FN_add(r[:], r[:], a[alen][:])
	}
}

func reverse_bits(i int, nbits int) int {
	r := 0
	for nbits > 0 {
		r = (r << 1) | (i & 1)
		i >>= 1
		nbits--
	}

	return r
}

func FFt(vals []*FN_t, H []*FN_t, nbits int) {
	n := 1 << nbits
	Hoffset := n / 2

	for i := 0; i < n; i++ {
		var tmp FN_t
		i_rev := reverse_bits(i, nbits)
		if i < i_rev {
			FN_copy(tmp[:], (*vals[i])[:])
			FN_copy((*vals[i])[:], (*vals[i_rev])[:])
			FN_copy((*vals[i_rev])[:], tmp[:])
		}
	}

	for i := 0; i < nbits; i++ {

		half_len := 1 << i
		full_len := half_len << 1
		count := (1 << nbits) / full_len

		for j := 0; j < count; j++ {
			Li := full_len * j
			Ri := Li + half_len
			Hi := 0

			for k := 0; k < half_len; k++ {
				var x FN_t
				var y FN_t

				FN_copy(x[:], (*vals[Li])[:])
				FN_copy(y[:], (*vals[Ri])[:])
				FN_mul(y[:], y[:], (*H[Hi])[:])
				FN_add((*vals[Li])[:], x[:], y[:])
				FN_sub((*vals[Ri])[:], x[:], y[:])

				Li++
				Ri++
				Hi += Hoffset
			}
		}
		Hoffset /= 2
	}
}

func Poly_interpolate(vals []*FN_t, H []*FN_t, nbits int) {
	var n_inv FN_t

	n := 1 << nbits

	FFt(vals, H, nbits)

	BN_set_word(n_inv[:], uint32(n), 8)
	FN_inv(n_inv, n_inv)

	Poly_mul_scalar(vals, &n, vals, n, n_inv)

	for i := 1; i < (n+1)/2; i++ {
		FN_copy(n_inv[:], (*vals[i])[:])
		FN_copy((*vals[i])[:], (*vals[n-i])[:])
		FN_copy((*vals[n-i])[:], n_inv[:])
	}
}
