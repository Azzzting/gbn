package main

import (
	"fmt"

	bn254 "github.com/Azzzting/gbn/include"
)

func example_bn(hex_a string, hex_b string) {
	//hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	//hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	r := make([]uint32, 16)
	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)
	bn254.BN_print(" a = ", a, 8)
	bn254.BN_print(" b = ", b, 8)
	//example BN_add
	bn254.BN_add(r, a, b, 8)
	bn254.BN_print(" a + b = ", r, 8)
	//example BN_sub
	bn254.BN_sub(r, a, b, 8)
	bn254.BN_print(" a - b = ", r, 8)
	bn254.BN_sub(r, b, a, 8)
	bn254.BN_print(" b - a = ", r, 8)
	//example mul_lo
	bn254.BN_mul_lo(r, a, b, 8)
	bn254.BN_print(" a * b (mod 2^256) = ", r, 8)
	//mul
	bn254.BN_mul(r, a, b, 8)
	bn254.BN_print(" a * b = ", r, 16)
}

func example_bn_mod(hex_a string, hex_b string) {
	//hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	//hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	hex_n := "30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	n := make([]uint32, 8)
	r := make([]uint32, 16)

	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)
	bn254.BN_from_hex(n, 8, hex_n)

	bn254.BN_print(" a = ", a, 8)
	bn254.BN_print(" b = ", b, 8)
	bn254.BN_print(" n = ", n, 8)

	bn254.BN_mod_add_non_const_time(r, a, b, n, 8)
	bn254.BN_print(" a + b (mod n) = ", r, 8)

	bn254.BN_mod_sub_non_const_time(r, a, b, n, 8)
	bn254.BN_print(" a - b (mod n) = ", r, 8)

	bn254.BN_mod_sub_non_const_time(r, b, a, n, 8)
	bn254.BN_print(" b - a (mod n) = ", r, 8)

	bn254.BN_mod_neg(r, a, n, 8)
	bn254.BN_print(" -a (mod n) = ", r, 8)

}

func example_bn_barrett(hex_a string, hex_b string) {

	//hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	//hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
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

	bn254.BN_print(" a = ", a, 8)
	bn254.BN_print(" b = ", b, 8)
	bn254.BN_print(" p = ", n, 8)
	bn254.BN_print(" mu(p) = ", n_mu, 9)

	bn254.BN_barrett_mod_mul(r, a, b, p, p_mu, 8)
	bn254.BN_print(" a * b (mod p) = ", r, 8)
	bn254.BN_barrett_mod_sqr(r, a, p, p_mu, 8)
	bn254.BN_print(" a^2 (mod p) = ", r, 8)
	bn254.BN_barrett_mod_exp(r, a, b, p, p_mu, 8)
	bn254.BN_print(" a^b (mod p) = ", r, 8)
	bn254.BN_barrett_mod_inv(r, a, p, p_mu, 8)
	bn254.BN_print(" a^-1 (mod p) = ", r, 8)
	bn254.BN_barrett_mod_mul(r, a, b, n, n_mu, 8)
	bn254.BN_print(" a * b (mod n) = ", r, 8)
	bn254.BN_barrett_mod_sqr(r, a, n, n_mu, 8)
	bn254.BN_print(" a^2 (mod n) = ", r, 8)
	bn254.BN_barrett_mod_exp(r, a, b, n, n_mu, 8)
	bn254.BN_print(" a^b (mod n) = ", r, 8)
	bn254.BN_barrett_mod_inv(r, a, n, n_mu, 8)
	bn254.BN_print(" a^-1 (mod n) = ", r, 8)
}

func example_bn_montgomery(hex_a string, hex_b string) {

	//hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	//hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
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

	bn254.BN_print(" a = ", a, 8)
	bn254.BN_print(" b = ", b, 8)
	bn254.BN_print(" p = ", p, 8)

	bn254.BN_mod_mul_montgomery(r, a, b, p_one_sqr, p, p_inv_neg, 8)
	bn254.BN_print(" a * b (mod p) = ", r, 8)

	bn254.BN_mod_sqr_montgomery(r, a, p_one_sqr, p, p_inv_neg, 8)
	bn254.BN_print(" a^2 (mod p) = ", r, 8)

	bn254.BN_from_hex(n, 8, hex_n)
	bn254.BN_from_hex(n_inv_neg, 8, hex_n_inv_neg)
	bn254.BN_from_hex(n_one_sqr, 8, hex_n_one_sqr)

	bn254.BN_print(" n = ", n, 8)

	bn254.BN_mod_mul_montgomery(r, a, b, n_one_sqr, n, n_inv_neg, 8)
	bn254.BN_print(" a * b (mod n) = ", r, 8)

	bn254.BN_mod_sqr_montgomery(r, a, n_one_sqr, n, n_inv_neg, 8)
	bn254.BN_print(" a^2 (mod n) = ", r, 8)
}

func example_fn(hex_a string, hex_b string) {
	//hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	//hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	a := make([]uint32, 8)
	b := make([]uint32, 8)
	r := make([]uint32, 8)
	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)
	bn254.BN_print(" a = ", a, 8)
	bn254.BN_print(" b = ", b, 8)
	bn254.BN_print(" n = ", bn254.BN254_n[:], 8)

	bn254.FN_from_hex(a, hex_a)
	bn254.FN_from_hex(b, hex_b)
	bn254.FN_add(r, a, b)
	bn254.FN_get_BN(r, r)
	bn254.BN_print(" a + b (mod n) = ", r, 8)

	bn254.FN_sub(r, a, b)
	bn254.FN_get_BN(r, r)
	bn254.BN_print(" a - b (mod n) = ", r, 8)

	bn254.FN_sub(r, b, a)
	bn254.FN_get_BN(r, r)
	bn254.BN_print(" b - a (mod n) = ", r, 8)

	bn254.FN_neg(r, a)
	bn254.FN_get_BN(r, r)
	bn254.BN_print(" -a (mod p) = ", r, 8)

	bn254.FN_mul(r, a, b)
	bn254.FN_get_BN(r, r)
	bn254.BN_print(" a * b (mod n) = ", r, 8)

	bn254.FN_sqr(r, a)
	bn254.FN_get_BN(r, r)
	bn254.BN_print(" a^2 (mod n) = ", r, 8)

	var a_fn bn254.FN_t
	for i := 0; i < len(a); i++ {
		a_fn[i] = a[i]
	}
	var r_fn bn254.FN_t
	for i := 0; i < len(r); i++ {
		r_fn[i] = r[i]
	}
	//bn254.FN_inv(&r_fn, a_fn)
	//bn254.FN_get_bn(r_fn[:], r_fn[:])

}

func example_fp(hex_a string, hex_b string) {
	//hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	//hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"

	//hex_fp_add_a_b := "1befa71cbccf15f1b86fc535f2cff36a01e50692e41e37f6051fd274038f069c"
	//hex_fp_sub_a_b := "2a5513860fc876382d0c7108e455c5b4ccd993da773a9cb5df66c4e2fd7f5309"
	//hex_fp_sub_b_a := "060f3aecd16929f18b43d4ad9d2b92a8caa7d6b6f1372dd75cb9c733dafdaa3e"
	//hex_fp_neg_a := "2574185aeb7eaa29a1ba4d7256af27fcfbe2d2a36efe457de7ed8676c4344f18"
	//hex_fp_mul_a_b := "26233c6c1df2a1b4fec81ed8a1a4fb2450a06a8adf3576d183a9f2b87b6d2b72"
	//hex_fp_sqr_a := "13b887f45214e1eed372c7b44bd4a7a8d60518114e10c581176cba7fdeccd283"
	//hex_fp_exp_a_b := "09223b41c088f253f48bcd362e406afe59f49bcff34cb9e1ce5d910c38f06476"
	//hex_fp_inv_a := "1c247c36f20b94c90c886532c309246addd378f27906754263f7af21d12a71a7"

	a := make([]uint32, 8)
	b := make([]uint32, 8)
	r := make([]uint32, 8)

	bn254.BN_from_hex(a, 8, hex_a)
	bn254.BN_from_hex(b, 8, hex_b)

	bn254.BN_print(" a = ", a, 8)
	bn254.BN_print(" b = ", b, 8)

	bn254.FP_from_hex(a, hex_a)
	bn254.FP_from_hex(b, hex_b)

	bn254.FP_add(r, a, b)
	bn254.FP_get_BN(r, r)
	bn254.BN_print(" a + b (mod p) = ", r, 8)
	//fmt.Printf("               = %s\n", hex_fp_add_a_b)

	bn254.FP_sub(r, a, b)
	bn254.FP_get_BN(r, r)
	bn254.BN_print(" a - b (mod p) = ", r, 8)
	//fmt.Printf("               = %s\n", hex_fp_sub_a_b)

	bn254.FP_sub(r, b, a)
	bn254.FP_get_BN(r, r)
	bn254.BN_print(" b - a (mod p) = ", r, 8)
	//fmt.Printf("               = %s\n", hex_fp_sub_b_a)

	bn254.FP_neg(r, a)
	bn254.FP_get_BN(r, r)
	bn254.BN_print(" -a (mod p) = ", r, 8)
	//fmt.Printf("            = %s\n", hex_fp_neg_a)

	bn254.FP_mul(r, a, b)
	bn254.FP_get_BN(r, r)
	bn254.BN_print(" a * b (mod p) = ", r, 8)
	//fmt.Printf("               = %s\n", hex_fp_mul_a_b)

	bn254.FP_sqr(r, a)
	bn254.FP_get_BN(r, r)
	bn254.BN_print(" a^2 (mod p) = ", r, 8)
	//fmt.Printf("             = %s\n", hex_fp_sqr_a)

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
	bn254.BN_print(" a^-1 (mod p) = ", r_fp[:], 8)
	//fmt.Printf("              = %s\n", hex_fp_inv_a)

}

func example_point(hex_a string) {
	//hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	hex_G1 := "0000000000000000000000000000000000000000000000000000000000000001" +
		"0000000000000000000000000000000000000000000000000000000000000002"
	//hex_2G1 := "030644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd3" +
	//	"15ed738c0e0a7c92e7845f96b2ae9c0a68a6a449e3538fc7ff3ebf7a5a18a2c4"
	//hex_3G1 := "0769bf9ac56bea3ff40232bcb1b6bd159315d84715b8e679f2d355961915abf0" +
	//	"2ab799bee0489429554fdb7c8d086475319e63b40b9c5b57cdf1ff3dd9fe2261"
	//hex_aG1 := "06dddf2da5ec0ed18248eb78d2fa3b06d371ddc479ed7d7e0fc932d551d1b054" +
	//	"1c9dfe0f9d67a2c33cb4b4e62e5e2e4ae70a2cc746e4e3d368d59918a0f0de4f"

	var BN254_G1 = bn254.Point_t{
		X: [8]uint32{0xc58f0d9d, 0xd35d438d, 0xf5c70b3d, 0x0a78eb28,
			0x7879462c, 0x666ea36f, 0x9a07df2f, 0x0e0a77c1},
		Y: [8]uint32{0x8b1e1b3a, 0xa6ba871b, 0xeb8e167b, 0x14f1d651,
			0xf0f28c58, 0xccdd46de, 0x340fbe5e, 0x1c14ef83},
		Z: [8]uint32{0xc58f0d9d, 0xd35d438d, 0xf5c70b3d, 0x0a78eb28,
			0x7879462c, 0x666ea36f, 0x9a07df2f, 0x0e0a77c1},
		Is_At_Infinity: 0,
	}

	var G1 bn254.Point_t
	var P bn254.Point_t
	var a bn254.FN_t

	bn254.BN_from_hex(a[:], 8, hex_a)
	bn254.BN_print(" a = ", a[:], 8)

	bn254.Point_from_hex(&G1, hex_G1)
	bn254.Point_print(" G1 = ", &G1)
	bn254.Point_print("    = ", &BN254_G1)
	//fmt.Printf("        = %s\n", hex_G1)

	bn254.Point_dbl(&P, &G1)
	bn254.Point_print(" 2 * G1 = ", &P)
	//fmt.Printf("        = %s\n", hex_2G1)

	bn254.Point_add(&P, &P, &G1)
	bn254.Point_print(" 3 * G1 = ", &P)
	//fmt.Printf("        = %s\n", hex_3G1)

	affineG1 := bn254.Affine_point_t{
		X: G1.X,
		Y: G1.Y,
	}

	bn254.Point_mul_affine_non_const_time(&P, a[:], &affineG1)
	bn254.Point_print(" a * G1 = ", &P)

	bn254.Point_mul_affine_const_time(&P, a[:], &affineG1)
	bn254.Point_print("        = ", &P)

	//fmt.Printf("        = %s\n", hex_aG1)

	bn254.Point_mul_generator(&P, a[:])
	bn254.Point_print("        = ", &P)

}

func example_point_multi_mul(a [8]bn254.FN_t) {
	x := [8]int{1, 2, 4, 8, 16, 32, 64, 128}

	var P [8]*bn254.Point_t
	var T [255]*bn254.Point_t
	var R bn254.Point_t
	//var a [8]bn254.FN_t

	for i := 0; i < 255; i++ {
		T[i] = new(bn254.Point_t)
	}

	for i := 0; i < 8; i++ {
		P[i] = new(bn254.Point_t)
		var a [8]uint32
		bn254.BN_set_word(a[:], uint32(x[i]), 8)
		bn254.Point_mul_generator(P[i], a[:])
	}

	var aConverted [8][8]uint32
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			aConverted[i][j] = uint32(a[i][j])
		}
	}

	bn254.Point_multi_mul_affine_pre_compute(P[:], T[:])

	bn254.Point_multi_mul_affine_with_pre_compute(&R, aConverted, T[:])

	bn254.Point_print("R = ", &R)
}

func example_poly(hex_a string) {
	hex_fa0 := "24beb63b9dac49a8be446e8023e90219bb327ee5e8f49485565703b60a64756b"
	hex_fa1 := "12197fdb7a5989d8d22b33b8605aa067e2bd285bc9e109a309fb71af74715ace"
	hex_fa2 := "0f141a426378638a8968a408f912be2bead6116ca53a59d735792ad68d49ee9b"
	hex_fa3 := "1a50bd98c7a9d6f3bb17750fa2a277623322f13ca863373c358d92d2fbe2dd67"
	hex_fa4 := "063e702321449c8fea1eda7e131387fc54198fb285a79f852e509ddc7f31e243"
	hex_fa5 := "1ed3e7e56b2d347ff6a1a0921c09db0380dd12be9e98a4601cf18e6be58ad9b9"
	hex_fa6 := "11a161912924587f57f7314e1da8c9d71ee952dfc8f8084a9bbfbf98b00a3ebc"
	hex_fa7 := "145d5ebb9c715335c3a97b7457071b05f3fbadbf5aa687d0487e06c87e5ec3d7"
	hex_fa8 := "1489a19338a05a1486f04c2773360d574f71f4545ba48eb35a50590981c02311"
	hex_fa9 := "1fd842bd9a4ce97e9ac5912146d012fee0b10f14265704f3c1e5206e19f99104"
	hex_fa10 := "2058ef90111d3bd23cb156b8332dfa8ed762f9398d92375c2b0b9bd3670c1f6f"
	hex_fa11 := "15404843ddfc9fdac37181b1510d73ad52743346b6787537333aa87ac39d0c47"
	hex_fa12 := "058db597d0d13bf2d8e1a9eefa697c83c3a5c11b1f51da8e827cf11be6c8ddc8"
	hex_fa13 := "19814b23927937f1eb1ae0d66a8a1902ea15eb1b7fbe2582332a47b7eb89497f"
	hex_fa14 := "08ef18ba4a24ffc24630e22242e33445bef64740dd0b5d5f61b880f63ea473e0"
	hex_fa15 := "226be63f280cddb5d3e473b717a67d79d2ebdcb057c68ff7b618da03df575125"

	hex_fb0 := "07dcc8a07e69b5c1cfc1c27d7b86d75193f67dc0421df4ef983223b10fb806d9"
	hex_fb1 := "0d11f84dacaffcb734075b94c816ee840a529398c4a9205a68af72603b0b81e4"
	hex_fb2 := "095809d32b986210a4aa53a6d0861d1b01a482168442464f134814739b3a55f2"
	hex_fb3 := "109c524d2953e124be6743f3fd4c152c40655142126424c6815fc3900771f210"
	hex_fb4 := "1feaac8c6be176e4f44afb51f948094c1246a5de699f3e41f7ff0862eebaecfc"
	hex_fb5 := "20454549504101187276e8996f87bbdd7b10d7e244e7f4319782bf1aeb4dd99d"
	hex_fb6 := "11b09372e71f09e501fa77e1e374bfa8915e7fa7c075b0e7723382ad22160e9a"
	hex_fb7 := "0bc368e47eb5e5595941f3f13f2cb04460e6a36291fc1ee816b4a66b06872eab"
	hex_fb8 := "08047f36b5f5525feb2a44e779002948db927ef3deeae245040f36eb0071fad0"
	hex_fb9 := "1fff65166080ba7a8abfe636fdbefeb659e436f9d94825e922cb98758888cd1f"
	hex_fb10 := "14098ea0f0fdc37d46abd1d9f772c4918c3e08f398c7dc32e4ae898f61102e99"
	hex_fb11 := "132c91b05e311560a8198ce0fab98099ddad6dedfad351e84c6c8e35a9a09455"
	hex_fb12 := "013ad65e3cefeb684dcc020e36fef91f88f2ce0956db4af1c2034cbdf881b90e"
	hex_fb13 := "1476c11de640dfd0f1e7b8b73a80d016289972f7b0f6c3601343f0e0993a6244"
	hex_fb14 := "1d512b84553f12ed786c46668d1006f7ad111ce484bc25435a524fd3b260a697"
	hex_fb15 := "0c6ee93a932f8591c5da3a6d3876040ef8b396c3463b8c0e39ab1632c1d28946"

	//hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"

	fa := bn254.Poly_New(16 + 4)
	fb := bn254.Poly_New(16 + 4)
	fr := bn254.Poly_New(32 + 4)
	fa_len := 16
	fb_len := 16
	fr_len := 0

	var a bn254.FN_t
	//var r bn254.FN_t

	bn254.FN_from_hex((*fa[0])[:], hex_fa0)
	bn254.FN_from_hex((*fa[1])[:], hex_fa1)
	bn254.FN_from_hex((*fa[2])[:], hex_fa2)
	bn254.FN_from_hex((*fa[3])[:], hex_fa3)
	bn254.FN_from_hex((*fa[4])[:], hex_fa4)
	bn254.FN_from_hex((*fa[5])[:], hex_fa5)
	bn254.FN_from_hex((*fa[6])[:], hex_fa6)
	bn254.FN_from_hex((*fa[7])[:], hex_fa7)
	bn254.FN_from_hex((*fa[8])[:], hex_fa8)
	bn254.FN_from_hex((*fa[9])[:], hex_fa9)
	bn254.FN_from_hex((*fa[10])[:], hex_fa10)
	bn254.FN_from_hex((*fa[11])[:], hex_fa11)
	bn254.FN_from_hex((*fa[12])[:], hex_fa12)
	bn254.FN_from_hex((*fa[13])[:], hex_fa13)
	bn254.FN_from_hex((*fa[14])[:], hex_fa14)
	bn254.FN_from_hex((*fa[15])[:], hex_fa15)

	bn254.FN_from_hex((*fb[0])[:], hex_fb0)
	bn254.FN_from_hex((*fb[1])[:], hex_fb1)
	bn254.FN_from_hex((*fb[2])[:], hex_fb2)
	bn254.FN_from_hex((*fb[3])[:], hex_fb3)
	bn254.FN_from_hex((*fb[4])[:], hex_fb4)
	bn254.FN_from_hex((*fb[5])[:], hex_fb5)
	bn254.FN_from_hex((*fb[6])[:], hex_fb6)
	bn254.FN_from_hex((*fb[7])[:], hex_fb7)
	bn254.FN_from_hex((*fb[8])[:], hex_fb8)
	bn254.FN_from_hex((*fb[9])[:], hex_fb9)
	bn254.FN_from_hex((*fb[10])[:], hex_fb10)
	bn254.FN_from_hex((*fb[11])[:], hex_fb11)
	bn254.FN_from_hex((*fb[12])[:], hex_fb12)
	bn254.FN_from_hex((*fb[13])[:], hex_fb13)
	bn254.FN_from_hex((*fb[14])[:], hex_fb14)
	bn254.FN_from_hex((*fb[15])[:], hex_fb15)

	bn254.FN_from_hex(a[:], hex_a)

	bn254.Poly_print("fa = ", fa, 16)
	bn254.Poly_print("fb = ", fb, 16)

	bn254.Poly_add(fr, &fr_len, fa, fa_len, fb, fb_len)
	bn254.Poly_print("fa + fb = ", fr, fr_len)

	bn254.Poly_sub(fr, &fr_len, fa, fa_len, fb, fb_len)
	bn254.Poly_print("fa - fb = ", fr, fr_len)

	bn254.Poly_mul(fr, &fr_len, fa, fa_len, fb, fb_len)
	bn254.Poly_print("fa * fb = ", fr, fr_len)

	bn254.Poly_add_scalar(fr, &fr_len, fa, fa_len, a)
	bn254.Poly_print("fa + a = ", fr, fr_len)

	bn254.Poly_sub_scalar(fr, &fr_len, fa, fa_len, a)
	bn254.Poly_print("fa - a = ", fr, fr_len)

	bn254.Poly_mul_scalar(fr, &fr_len, fa, fa_len, a)
	bn254.Poly_print("fa * a = ", fr, fr_len)

	bn254.Poly_mul_vector(fr, &fr_len, fa, fa_len, fb, fb_len)
	bn254.Poly_print("mul_vec(fa, fb) = ", fr, fr_len)

	bn254.Poly_copy(fr, &fr_len, fa, fa_len)
	bn254.Poly_add_blind(fr, &fr_len, fb, 2)
	bn254.Poly_print("blind2(a) = ", fr, fr_len)

	bn254.Poly_copy(fr, &fr_len, fa, fa_len)
	bn254.Poly_add_blind(fr, &fr_len, fb, 3)
	bn254.Poly_print("blind3(a) = ", fr, fr_len)

	bn254.Poly_copy(fr, &fr_len, fa, fa_len)
	bn254.Poly_div_x_sub_scalar(fr, &fr_len, a)
	bn254.Poly_print("fa/(x-a) = ", fr, fr_len)

	//bn254.Poly_copy(fr, &fr_len, fa, fa_len)
	//bn254.Poly_div_ZH(fr, &fr_len, fa, fa_len, 16)
	//bn254.Poly_print("fa/(x^n-1) = ", fr, fr_len)

	//bn254.Poly_eval(r, fa, fa_len, a)
	//bn254.FN_print("fa(a) = ", r[:])

}

func example_fft() {

	hex_fa0 := "24beb63b9dac49a8be446e8023e90219bb327ee5e8f49485565703b60a64756b"
	hex_fa1 := "12197fdb7a5989d8d22b33b8605aa067e2bd285bc9e109a309fb71af74715ace"
	hex_fa2 := "0f141a426378638a8968a408f912be2bead6116ca53a59d735792ad68d49ee9b"
	hex_fa3 := "1a50bd98c7a9d6f3bb17750fa2a277623322f13ca863373c358d92d2fbe2dd67"
	hex_fa4 := "063e702321449c8fea1eda7e131387fc54198fb285a79f852e509ddc7f31e243"
	hex_fa5 := "1ed3e7e56b2d347ff6a1a0921c09db0380dd12be9e98a4601cf18e6be58ad9b9"
	hex_fa6 := "11a161912924587f57f7314e1da8c9d71ee952dfc8f8084a9bbfbf98b00a3ebc"
	hex_fa7 := "145d5ebb9c715335c3a97b7457071b05f3fbadbf5aa687d0487e06c87e5ec3d7"
	hex_fa8 := "1489a19338a05a1486f04c2773360d574f71f4545ba48eb35a50590981c02311"
	hex_fa9 := "1fd842bd9a4ce97e9ac5912146d012fee0b10f14265704f3c1e5206e19f99104"
	hex_fa10 := "2058ef90111d3bd23cb156b8332dfa8ed762f9398d92375c2b0b9bd3670c1f6f"
	hex_fa11 := "15404843ddfc9fdac37181b1510d73ad52743346b6787537333aa87ac39d0c47"
	hex_fa12 := "058db597d0d13bf2d8e1a9eefa697c83c3a5c11b1f51da8e827cf11be6c8ddc8"
	hex_fa13 := "19814b23927937f1eb1ae0d66a8a1902ea15eb1b7fbe2582332a47b7eb89497f"
	hex_fa14 := "08ef18ba4a24ffc24630e22242e33445bef64740dd0b5d5f61b880f63ea473e0"
	hex_fa15 := "226be63f280cddb5d3e473b717a67d79d2ebdcb057c68ff7b618da03df575125"

	hex_w0 := "0000000000000000000000000000000000000000000000000000000000000001"
	hex_w1 := "21082ca216cbbf4e1c6e4f4594dd508c996dfbe1174efb98b11509c6e306460b"
	hex_w2 := "2b337de1c8c14f22ec9b9e2f96afef3652627366f8170a0a948dad4ac1bd5e80"
	hex_w3 := "107aab49e65a67f9da9cd2abf78be38bd9dc1d5db39f81de36bcfa5b4b039043"
	hex_w4 := "30644e72e131a029048b6e193fd841045cea24f6fd736bec231204708f703636"
	hex_w5 := "2290ee31c482cf92b79b1944db1c0147635e9004db8c3b9d13644bef31ec3bd3"
	hex_w6 := "1d59376149b959ccbd157ac850893a6f07c2d99b3852513ab8d01be8e846a566"
	hex_w7 := "2d8040c3a09c49698c53bfcb514d55a5b39e9b17cb093d128b8783adb8cbd723"
	hex_w8 := "30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000000"
	hex_w9 := "0f5c21d0ca65e0db9be1f670eca407d08ec5ec67626a74f892ccebcd0cf9b9f6"
	hex_w10 := "0530d09118705106cbb4a786ead16926d5d174e181a26686af5448492e42a181"
	hex_w11 := "1fe9a328fad7382fddb3730a89f574d14e57caeac619eeb30d24fb38a4fc6fbe"
	hex_w12 := "0000000000000000b3c4d79d41a91758cb49c3517c4604a520cff123608fc9cb"
	hex_w13 := "0dd360411caed09700b52c71a6655715c4d558439e2d34f4307da9a4be13c42e"
	hex_w14 := "130b17119778465cfb3acaee30f81dee20710ead41671f568b11d9ab07b95a9b"
	hex_w15 := "02e40daf409556c02bfc85eb303402b774954d30aeb0337eb85a71e6373428de"

	a := bn254.Poly_New(16)
	H := bn254.Poly_New(16)

	bn254.FN_from_hex((*a[0])[:], hex_fa0)
	bn254.FN_from_hex((*a[1])[:], hex_fa1)
	bn254.FN_from_hex((*a[2])[:], hex_fa2)
	bn254.FN_from_hex((*a[3])[:], hex_fa3)
	bn254.FN_from_hex((*a[4])[:], hex_fa4)
	bn254.FN_from_hex((*a[5])[:], hex_fa5)
	bn254.FN_from_hex((*a[6])[:], hex_fa6)
	bn254.FN_from_hex((*a[7])[:], hex_fa7)
	bn254.FN_from_hex((*a[8])[:], hex_fa8)
	bn254.FN_from_hex((*a[9])[:], hex_fa9)
	bn254.FN_from_hex((*a[10])[:], hex_fa10)
	bn254.FN_from_hex((*a[11])[:], hex_fa11)
	bn254.FN_from_hex((*a[12])[:], hex_fa12)
	bn254.FN_from_hex((*a[13])[:], hex_fa13)
	bn254.FN_from_hex((*a[14])[:], hex_fa14)
	bn254.FN_from_hex((*a[15])[:], hex_fa15)

	bn254.FN_from_hex((*H[0])[:], hex_w0)
	bn254.FN_from_hex((*H[1])[:], hex_w1)
	bn254.FN_from_hex((*H[2])[:], hex_w2)
	bn254.FN_from_hex((*H[3])[:], hex_w3)
	bn254.FN_from_hex((*H[4])[:], hex_w4)
	bn254.FN_from_hex((*H[5])[:], hex_w5)
	bn254.FN_from_hex((*H[6])[:], hex_w6)
	bn254.FN_from_hex((*H[7])[:], hex_w7)
	bn254.FN_from_hex((*H[8])[:], hex_w8)
	bn254.FN_from_hex((*H[9])[:], hex_w9)
	bn254.FN_from_hex((*H[10])[:], hex_w10)
	bn254.FN_from_hex((*H[11])[:], hex_w11)
	bn254.FN_from_hex((*H[12])[:], hex_w12)
	bn254.FN_from_hex((*H[13])[:], hex_w13)
	bn254.FN_from_hex((*H[14])[:], hex_w14)
	bn254.FN_from_hex((*H[15])[:], hex_w15)

	bn254.FFt(a, H, 4)

	bn254.Poly_print("a(H) = ", a, 16)

	bn254.Poly_interpolate(a, H, 4)
	bn254.Poly_print("a = ", a, 16)
}

func main() {
	fmt.Printf("example begin......\n")
	fmt.Printf("example_bn......\n")
	var hex_a_bn string
	var hex_b_bn string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_bn) //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	fmt.Println("b=")
	fmt.Scanln(&hex_b_bn) //hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	example_bn(hex_a_bn, hex_b_bn)

	fmt.Printf("example_bn_mod......\n")
	var hex_a_bn_mod string
	var hex_b_bn_mod string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_bn_mod) //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	fmt.Println("b=")
	fmt.Scanln(&hex_b_bn_mod) //hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	example_bn_mod(hex_a_bn_mod, hex_b_bn_mod)

	fmt.Printf("example_bn_barrett......\n")
	var hex_a_bn_barrett string
	var hex_b_bn_barrett string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_bn_barrett) //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	fmt.Println("b=")
	fmt.Scanln(&hex_b_bn_barrett) //hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	example_bn_barrett(hex_a_bn_barrett, hex_b_bn_barrett)

	fmt.Printf("example_bn_montgomery......\n")
	var hex_a_bn_montgomery string
	var hex_b_bn_montgomery string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_bn_montgomery) //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	fmt.Println("b=")
	fmt.Scanln(&hex_b_bn_montgomery) //hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	example_bn_montgomery(hex_a_bn_montgomery, hex_b_bn_montgomery)

	fmt.Printf("example_fn......\n")
	var hex_a_fn string
	var hex_b_fn string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_fn) //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	fmt.Println("b=")
	fmt.Scanln(&hex_b_fn) //hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	example_fn(hex_a_fn, hex_b_fn)

	fmt.Printf("example_fp......\n")
	var hex_a_fp string
	var hex_b_fp string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_fp) //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	fmt.Println("b=")
	fmt.Scanln(&hex_b_fp)          //hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	example_fp(hex_a_fp, hex_b_fp) //√

	fmt.Printf("example_point......\n")
	var hex_a_point string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_point)   //hex_a_point := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	example_point(hex_a_point) //√

	fmt.Printf("example_point_multi_mul......\n")
	var a_point_multi_mul [8]bn254.FN_t //1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64
	fmt.Printf("请输入......\n")           //需要一个一个打进去
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Scanln(&a_point_multi_mul[i][j])
		}
	}
	example_point_multi_mul(a_point_multi_mul) //√

	fmt.Printf("example_poly......\n")
	var hex_a_poly string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_poly)  //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	example_poly(hex_a_poly) //√

	fmt.Printf("example_fft......\n")
	example_fft() //√ //无输入
}
