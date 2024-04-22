package main

import(
	"fmt"
	"github.com/Azzzting/gbn"
)

func test_bn(){
	var a[8] uint32;
	var b[8] uint32;
	bn_from_hex(a, 8, hex_a);
	bn_from_hex(b, 8, hex_b);
	bn_print(" a = ", a, 8);
	bn_print(" b = ", b, 8);
}

func main(){
	test_bn();
}
