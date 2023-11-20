package One

// go 实现三元表达式、

func If(c bool, a, b interface{}) interface{} {
	if c {
		return a
	}
	return b
}
