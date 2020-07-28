package algorithm

//最大公约数
func gcd(a int,b int) int{
	if a < b {
		a,b = b,a
	}

	if b == 0 {
		return a
	}

	return gcd(b, a %b )
}


//最小公位数
func lcm(a int,b int) int{

	var g_c_d = gcd(a,b)
	return a*b / g_c_d
}