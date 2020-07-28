package algorithm

import (
	"testing"
	"fmt"
)


func TestGcd(t *testing.T) {

	var g_c_d = gcd(12,5)
	fmt.Println(g_c_d)

	if g_c_d != 1 {
		t.Error()
	}

	g_c_d = gcd(12,6)
	fmt.Println(g_c_d)
	
	if g_c_d != 6 {
		t.Error()
	}
}


func TestLcm(t *testing.T) {
	var l_c_m = lcm(12,5)
	fmt.Println(l_c_m)

	l_c_m = lcm(12,4)
	fmt.Println(l_c_m)
}