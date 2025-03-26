## GBN:椭圆曲线BN族的GO语言库重写
该项目是针对零知识证明软件核心代码部分的**零知识证明算法库**的go语言重写。

原项目使用的是Zkrypt，该库是一个开源的C语言零知识证明算法库，原C++项目链接：https://github.com/guanzhi/zkrypt
### 项目介绍📚
* 本项目旨在完成对零知识证明算法go语言的扩张，填补BN椭圆曲线族在go语言领域的空白。
* 在本项目中，我们完成了有关BN曲线上的几乎所有运算，包括但不限于加、减、乘、除、求模逆等等。
* 本项目分为include文件和test文件，看两者分别对应着C++的头文件以及测试文件。
* 在使用过程中，**include**文件夹中的文件可以当作头文件直接引用，方便用户集成到自己的项目中。
* 在**test**文件夹中，提供了关于这些运算的完整测试，确保每个操作的正确性和稳定性。
### 使用指南🔍
下载文件之后运行指令：
```
go mod tidy
```
因为该库已经上载到了Go包管理网站：https://pkg.go.dev/, 所以mod可以自主下载。

在命令完成之后，可以像C++的include指令一样，采用以下命令即可完成头文件的引入：
```
import https://github.com/Azzzting/gbn/include
```
完成上述操作就可以愉快的使用啦！🎉

示例代码：
```go
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
}
```
### 测试结果 :white_check_mark:
测试图示：
+ 大数运算：
<img width="600" alt="image" src="https://github.com/user-attachments/assets/10a8729a-0512-47d7-abe5-a2b7c0a21591" />

+ 蒙哥马利模乘：
<img width="600" alt="image" src="https://github.com/user-attachments/assets/81d5affa-3a5c-4677-810c-a48648c3017c" />

+ 标量域运算：
<img width="600" alt="image" src="https://github.com/user-attachments/assets/9eb5deb6-06bd-4485-bb5e-aee2d2964796" />

+ 快速傅里叶变换：
<img width="600" alt="image" src="https://github.com/user-attachments/assets/728dcdff-2c68-47cd-80bb-42cdc489e136" />

### 代码贡献🤝
我们欢迎社区的贡献！如果你发现问题或者有任何优化建议，欢迎提交Issue或者Fork并提交PR。🚀
