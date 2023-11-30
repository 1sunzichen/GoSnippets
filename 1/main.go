package One

// go 实现三元表达式、
// 查看伪汇编代码
// go build -gcflags -S main.go 查看伪汇编代码
func If(c bool, a, b interface{}) interface{} {
	if c {
		return a
	}
	return b
}
