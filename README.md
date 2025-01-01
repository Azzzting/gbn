## GBN:椭圆曲线BN族的GO语言库重写
该项目是针对零知识证明软件核心代码部分的**零知识证明算法库**的go语言重写。

原项目使用的是Zkrypt，该库是一个开源的C语言零知识证明算法库，原C++项目链接：https://github.com/guanzhi/zkrypt
### 项目介绍
* 本项目旨在完成对零知识证明算法go语言的扩张，填补BN椭圆曲线族在go语言领域的空白。
* 在本项目中，我们完成了有关BN曲线上的几乎所有运算，包括但不限于加、减、乘、除、求模逆等等。
* 本项目分为include文件和test文件，看两者分别对应着C++的头文件以及测试文件。
* 在使用过程中，include文件夹中的文件可以当作头文件直接引用。
* 在test文件夹中，我们给出了关于这些运算的测试，结果均正确。
### 使用指南
下载文件之后运行指令：
```
go mod tidy
```
因为该库已经上载到了Go包管理网站：https://pkg.go.dev/, 所以mod可以自主下载。

在命令完成之后，可以像C++的include指令一样，采用以下命令即可完成头文件的引入：
```
import https://github.com/Azzzting/gbn/include
```
示例代码：
```
func main() {
	fmt.Printf("test_bn......\n")
	var hex_a_bn string
	var hex_b_bn string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_bn) //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	fmt.Println("b=")
	fmt.Scanln(&hex_b_bn) //hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	test_bn(hex_a_bn, hex_b_bn)

	fmt.Printf("test_bn_mod......\n")
	var hex_a_bn_mod string
	var hex_b_bn_mod string
	fmt.Println("a=")
	fmt.Scanln(&hex_a_bn_mod) //hex_a := "0af03617f5b2f6001695f8442ad230609b9e97edf973850f543305a01448ae2f"
	fmt.Println("b=")
	fmt.Scanln(&hex_b_bn_mod) //hex_b := "10ff7104c71c1ff1a1d9ccf1c7fdc30966466ea4eaaab2e6b0ecccd3ef46586d"
	test_bn_mod(hex_a_bn_mod, hex_b_bn_mod)
  ·······
```

完成上述操作就可以愉快的使用啦！
