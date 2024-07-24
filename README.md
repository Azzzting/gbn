## GBN:椭圆曲线BN族的GO语言库重写
该库是对Zkrypt是一个开源的C语言零知识证明算法库的go语言重写，原C++项目链接：https://github.com/guanzhi/zkrypt
### 项目介绍
本项目旨在完成对零知识证明算法go语言的扩张，填补BN椭圆曲线族在go语言领域的空白。
在本项目中，我们完成了有关BN曲线上的几乎所有运算，包括但不限于加、减、乘、除、求模逆等等。
本项目分为include文件和test文件，看两者分别对应着C++的头文件以及测试文件。
在使用过程中，include文件夹中的文件可以当作头文件直接引用。
在test文件夹中，我们给出了关于这些运算的测试，结果均正确。
### 使用指南
下载文件之后运行指令：
```
go mod tidy
```
之后可以像C++的include文件一样，采用以下命令即可完成头文件的引入：
```
import https://github.com/Azzzting/gbn/include
```
完成上述操作就可以愉快的使用啦！
