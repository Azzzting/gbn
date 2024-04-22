package bn254_fn

import (
	"encoding/binary"
	"fmt"
	"os"
)

var BN254_p [8]uint32
var BN254_n [8]uint32
var BN254_p_mu [9]uint32
var BN254_n_mu [9]uint32
var BN254_p_inv_neg [8]uint32
var BN254_p_mont_one [8]uint32
var BN254_p_mont_one_sqr [8]uint32
var BN254_n_inv_neg [8]uint32
var BN254_n_mont_one [8]uint32
var BN254_n_mont_one_sqr [8]uint32
var BN254_one = [8]uint32{1, 0, 0, 0, 0, 0, 0, 0}
var BN254_two = [8]uint32{2, 0, 0, 0, 0, 0, 0, 0}

// ues mgml or not
var UseMontgomery = false

// func
func BN_print(str *string, a *[]uint32, k int) {
	fmt.Printf("%s", *str)
	for i := k - 1; i >= 0; i-- {
		fmt.Printf("%08x", (*a)[i])
	}
	fmt.Println()
}

func BN_set_word(r *[]uint32, a uint32, k int) {
	(*r)[0] = a
	for i := k - 1; i >= 1; i-- {
		(*r)[i] = 0
	}
}

func BN_copy(r *[]uint32, a *[]uint32, k int) {
	for i := k - 1; i >= 0; i-- {
		(*r)[k] = (*a)[k]
	}
}

func BN_cmp(a *[]uint32, b *[]uint32, k int) int {
	for i := k - 1; i >= 0; i-- {
		if (*a)[k] > (*b)[k] {
			return 1
		}
		if (*a)[k] < (*b)[k] {
			return -1
		}
	}
	return 0
}

func BN_is_zero(a *[]uint32, k int) (result int) {
	for i := k - 1; i >= 0; i-- {
		if (*a)[i] == 1 {
			return 0
		}
	}
	return 1
}

func BN_is_one(a *[]uint32, k int) (result int) {
	if (*a)[0] != 1 {
		return 0
	}
	for i := k - 1; i >= 1; i-- {
		if (*a)[i] == 1 {
			return 0
		}
	}
	return 1
}

func BN_add(r *[]uint32, a *[]uint32, b *[]uint32, k int) int {
	var w uint64 = 0
	for i := 0; i < k; i++ {
		w += uint64((*a)[i]) + uint64((*b)[i])
		(*r)[i] = uint32(w & 0xffffffff)
		w >>= 32
	}
	return int(w)
}

func BN_sub(r *[]uint32, a *[]uint32, b *[]uint32, k int) int {
	var w uint64 = 0
	for i := 0; i < k; i++ {
		w += uint64((*a)[i]) - uint64((*b)[i])
		(*r)[i] = uint32(w & 0xffffffff)
		w >>= 32
	}
	return int(w)
}

func BN_mul(r *[]uint32, a *[]uint32, b *[]uint32, k int) {
	for i := 0; i < k; i++ {
		(*r)[i] = 0
	}
	for i := 0; i < k; i++ {
		var w uint64 = 0
		for j := 0; j < k; j++ {
			w += uint64((*r)[i+j]) + uint64((*a)[i])*uint64((*b)[j])
			(*r)[i+j] = uint32(w & 0xffffffff)
			w >>= 32
		}
		(*r)[i+k] = uint32(w)
	}
}

func BN_mul_lo(r *[]uint32, a *[]uint32, b *[]uint32, k int) {
	for i := 0; i < k; i++ {
		(*r)[i] = 0
	}
	for i := 0; i < k; i++ {
		var w uint64 = 0
		for j := 0; j < k-i; j++ {
			w += uint64((*r)[i+j]) + uint64((*a)[i])*uint64((*b)[j])
			(*r)[i+j] = uint32(w & 0xffffffff)
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

func BN_to_bytes(a *[]uint32, k int, out *[]uint8) {
	for i := k - 1; i >= 0; i-- {
		Uint32_to_bytes((*a)[k], (*out)[i*4:])
	}
}

func BN_from_bytes(a *[]uint32, k int, in *[]uint8) {
	for i := k - 1; i >= 0; i-- {
		(*a)[k] = Uint32_from_bytes((*in)[i*4:])
	}
}

func BN_to_hex(a *[]uint32, k int, out *[]string) {
	for i := k - 1; i >= 0; i-- {
		(*out)[i] = fmt.Sprintf("%08x", (*a)[i])
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

func Hex2bin(in *[]rune, inlen int, out *[]uint8) int {
	if inlen%2 != 0 {
		return -1
	}
	for i := 0; i < inlen; i += 2 {
		c1 := Hexchar2int((*in)[i])
		c2 := Hexchar2int((*in)[i+1])
		if c1 < 0 || c2 < 0 {
			return -1
		}

		(*out)[i/2] = uint8(c1<<4 | c2)
	}
	return 1
}

func BN_from_hex(r *[]uint32, k int, hex string) int {
	buf := make([]uint8, k*4)
	runes := []rune(hex)
	if Hex2bin(&runes, len(hex), &buf) < 0 {
		return -1
	}
	BN_from_bytes(r, k, &buf)
	return 1
}
func BN_equ_hex(a *[]uint32, k int, hex string) bool {
	a_ := make([]uint32, k)
	BN_from_hex(&a_, k, hex)
	return BN_cmp(a, &a_, k) == 0
}

func BN_rand(r *[]uint32, k int) int {
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
		(*r)[i] = binary.BigEndian.Uint32(data[i*4 : (i+1)*4])
	}

	return 1
}

func BN_rand_range(r *[]uint32, rng *[]uint32, k int) int {
	for {
		BN_rand(r, k)
		if BN_cmp(r, rng, k) < 0 {
			break
		}
	}
	return 1
}

func BN_mod_add_non_const_time(r *[]uint32, a *[]uint32, b *[]uint32, p *[]uint32, k int) {
	BN_add(r, a, b, k)
	if BN_cmp(r, p, k) >= 0 {
		BN_sub(r, r, p, k)
	}
}

func BN_mod_sub_const_time(r *[]uint32, a *[]uint32, b *[]uint32, p *[]uint32, k int) {
	r2 := make([]uint32, 2*k)
	r3 := r2[k:]
	w := BN_sub(&r2, a, b, k)
	BN_sub(&r3, b, a, k)
	copy(r2[k:], r3)
	BN_sub(&r2, p, &r2, k)
	r4 := r2[k+k*w:]
	*r = make([]uint32, k)
	copy(*r, r4)
}
func BN_mod_sub_non_const_time(r *[]uint32, a *[]uint32, b *[]uint32, p *[]uint32, k int) {
	if BN_cmp(a, b, k) >= 0 {
		BN_sub(r, a, b, k)
	} else {
		BN_sub(r, b, a, k)
		BN_sub(r, p, r, k)
	}
}

func BN_mod_add__time(r *uint32, a *uint32, b *uint32, p *uint32, k int)
func BN_mod_sub__time(r *uint32, a *uint32, b *uint32, p *uint32, k int)

func BN_mod_neg(r *[]uint32, a *[]uint32, p *[]uint32, k int) {
	BN_sub(r, p, a, k)
}

func BN_reduce_once(r *[]uint32, a *[]uint32, b *[]uint32, k int) {
	r2 := make([]uint32, 2*k)

	// 将 a 复制到 r2 的前半部分
	copy(r2[:k], *a)
	r3 := r2[k:]
	// 计算 r1 = a - b
	w := BN_sub(&r3, a, b, k)

	// 将 r1 复制到 r
	r4 := r2[k+k*w:]
	*r = make([]uint32, k)
	copy(*r, r4)
}

func BN_barrett_mod_mul(r *[]uint32, a *[]uint32, b *[]uint32, p *[]uint32, u *[]uint32, k int) {
	z := make([]uint32, 2*k)
	q := make([]uint32, 2*(k+1))
	p_ := make([]uint32, k+1)
	t_ := make([]uint32, k+1)
	r_ := make([]uint32, k+1)

	// 复制 p 的内容到 p_
	for i := 0; i < k; i++ {
		p_[i] = (*p)[i]
	}
	p_[k] = 0

	// 计算 z = a * b
	BN_mul(&z, a, b, k)

	// 计算 q = z * u
	z1 := z[k-1:]
	BN_mul(&q, &z1, u, k+1)

	// 计算 t_ = q * p_
	q1 := q[k+1:]
	BN_mul_lo(&t_, &q1, &p_, k+1)

	// 计算 r_ = z - t_
	BN_sub(&r_, &z, &t_, k+1)

	// 通过 BN_reduce_once 函数进行约简操作
	BN_reduce_once(&r_, &r_, &p_, k+1)
	BN_reduce_once(&r_, &r_, &p_, k)

	// 将 r_ 的内容复制到 r 中
	for i := 0; i < k; i++ {
		(*r)[i] = r_[i]
	}
}

func BN_barrett_mod_sqr(r *[]uint32, a *[]uint32, p *[]uint32, u *[]uint32, k int) {
	BN_barrett_mod_mul(r, a, a, p, u, k)
}

func BN_barrett_mod_exp(r *[]uint32, a *[]uint32, e *[]uint32, p *[]uint32, u *[]uint32, k int) {
	var t *[]uint32
	w := uint32(0)

	// t = 1
	BN_set_one(t, k)

	for i := k - 1; i >= 0; i-- {
		w = (*e)[i]
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

func BN_barrett_mod_inv(r *[]uint32, a *[]uint32, p *[]uint32, u *[]uint32, k int) {
	e := make([]uint32, k)
	e[0] = 2
	for i := 1; i < k; i++ {
		e[i] = 0
	}
	BN_sub(&e, p, &e, k)
	BN_barrett_mod_exp(r, a, &e, p, u, k)
}

func BN_mont_mod_mul(r *[]uint32, a *[]uint32, b *[]uint32, p *[]uint32, p_inv_neg *[]uint32, k int) {
	z := make([]uint32, k*2)
	c := make([]uint32, k*2)
	t := make([]uint32, k)
	BN_mul(&z, a, b, k)
	BN_mul_lo(&t, &z, p_inv_neg, k)
	BN_mul(&c, &t, p, k)
	BN_add(&c, &c, &z, k*2)
	c1 := c[k:]
	if BN_cmp(&c1, p, k) >= 0 {
		BN_sub(&c1, &c1, p, k)
	}
	copy(c[k:], c1)
	for i := 0; i < k; i++ {
		(*r)[i] = c[k+i]
	}
}

func BN_mont_mod_sqr(r *[]uint32, a *[]uint32, p *[]uint32, p_inv_neg *[]uint32, k int) {
	BN_mont_mod_mul(r, a, a, p, p_inv_neg, k)
}

func BN_mont_mod_exp(r *[]uint32, a *[]uint32, e *[]uint32, p *[]uint32, p_inv_neg *[]uint32, k int) {
	var t *[]uint32
	BN_set_one(t, k)

	for i := k - 1; i >= 0; i-- {
		w := (*e)[i]
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

func BN_mont_mod_inv(r *[]uint32, a *[]uint32, p *[]uint32, p_inv_neg *[]uint32, k int) {
	e := make([]uint32, k)
	e[0] = 2
	for i := 1; i < k; i++ {
		e[i] = 0
	}
	BN_sub(&e, p, &e, k)
	BN_mont_mod_exp(r, a, &e, p, p_inv_neg, k)
}

func BN_mont_set(r *[]uint32, a *[]uint32, R_sqr *[]uint32, p *[]uint32, p_inv_neg *[]uint32, k int) {
	BN_mont_mod_mul(r, a, R_sqr, p, p_inv_neg, k)
}

func BN_mont_get(r *[]uint32, a *[]uint32, p *[]uint32, p_inv_neg *[]uint32, k int) {
	var one *[]uint32
	BN_set_one(one, k)
	BN_mont_mod_mul(r, a, one, p, p_inv_neg, k)
}

// for testing
func BN_mod_mul_montgomery(r *[]uint32, a *[]uint32, b *[]uint32, R_sqr *[]uint32, p *[]uint32, p_inv_neg *[]uint32, k int) {
	a_ := make([]uint32, k)
	b_ := make([]uint32, k)
	r_ := make([]uint32, k)
	BN_mont_set(&a_, a, R_sqr, p, p_inv_neg, k)
	BN_mont_set(&b_, b, R_sqr, p, p_inv_neg, k)
	BN_mont_mod_mul(&r_, &a_, &b_, p, p_inv_neg, k)
	BN_mont_get(r, &r_, p, p_inv_neg, k)
}

// also testing
func BN_mod_sqr_montgomery(r *[]uint32, a *[]uint32, R_sqr *[]uint32, p *[]uint32, p_inv_neg *[]uint32, k int) {
	a_ := make([]uint32, k)
	r_ := make([]uint32, k)
	BN_mont_set(&a_, a, R_sqr, p, p_inv_neg, k)
	BN_mont_mod_sqr(&r_, &a_, p, p_inv_neg, k)
	BN_mont_get(r, &r_, p, p_inv_neg, k)
}

type Fn_t [8]uint32

// define
func BN_set_zero(r *[]uint32, k int) {
	BN_set_word(r, 0, k)
}
func BN_set_one(r *[]uint32, k int) {
	BN_set_word(r, 1, k)
}
func BN_mod_add(r *[]uint32, a *[]uint32, b *[]uint32, p *[]uint32, k int) {
	//BN_mod_add__time(r,a,b,p,k);
	BN_mod_add_non_const_time(r, a, b, p, k)
}
func BN_mod_sub(r *[]uint32, a *[]uint32, b *[]uint32, p *[]uint32, k int) {
	//BN_mod_sub__time(r,a,b,p,k);
	BN_mod_sub_non_const_time(r, a, b, p, k)
}
func FN_print(str *string, a *[]uint32) {
	BN_print(str, a, 8)
}
func FN_rand(r *[]uint32, a *[]uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	BN_rand_range(r, &rng, 8)
}
func FN_copy(r *[]uint32, a *[]uint32) {
	BN_copy(r, a, 8)
}
func FN_set_word(r *[]uint32, a uint32) {
	BN_set_word(r, a, 8)
}
func FN_set_zero(r *[]uint32, k int) {
	BN_set_zero(r, 8)
}
func FN_is_zero(a *[]uint32) {
	BN_is_zero(a, 8)
}
func FN_add(r *[]uint32, a *[]uint32, b *[]uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	BN_mod_add(r, a, b, &rng, 8)
}
func FN_sub(r *[]uint32, a *[]uint32, b *[]uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	BN_mod_sub(r, a, b, &rng, 8)
}
func FN_neg(r *[]uint32, a *[]uint32, p *[]uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	BN_mod_neg(r, a, &rng, 8)
}
func FN_dbl(r *[]uint32, a *[]uint32) {
	FN_add(r, a, a)
}
func FN_to_bytes(a *[]uint32, out *[]uint8) {
	BN_to_bytes(a, 8, out)
}
func FN_from_bytes(a *[]uint32, in *[]uint8) {
	BN_from_bytes(a, 8, in)
}
func FN_set_BN(r *[]uint32, a *[]uint32) {
	FN_copy(r, a)
}
func FN_set_one(r *[]uint32) {
	BN_set_one(r, 8)
}
func FN_to_hex(a *[]uint32, out *[]string) {
	BN_to_hex(a, 8, out)
}
func FN_from_hex(r *[]uint32, hex string) {
	BN_from_hex(r, 8, hex)
}
func FN_get_BN(r *[]uint32, a *[]uint32) {
	FN_copy(a, r)
}
func FN_is_one(a *[]uint32) {
	BN_is_one(a, 8)
}
func FN_mul(r *[]uint32, a *[]uint32, b *[]uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	rng1 := make([]uint32, len(BN254_n_mu))
	copy(rng1, BN254_n_mu[:])
	BN_barrett_mod_mul(r, a, b, &rng, &rng1, 8)
}
func FN_sqr(r *[]uint32, a *[]uint32, p *[]uint32, u *[]uint32) {
	rng := make([]uint32, len(BN254_n))
	copy(rng, BN254_n[:])
	rng1 := make([]uint32, len(BN254_n_mu))
	copy(rng1, BN254_n_mu[:])
	BN_barrett_mod_sqr(r, a, &rng, &rng1, 8)
}
