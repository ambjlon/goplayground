# goplayground
## goplaygroud安装
go get github.com/ambjlon/goplayground

## goplayground使用
1. go run main.go
待测试的逻辑写在app目录下, 在main函数调用待测的函数. 注释main函数中的其他测试函数.
2. go test  
书写规则

+ 文件名必须是\*_test.go. **必须!!**
+ 必须import testing这个包.
+ 有的(测试)用例函数必须是Test开头.
+ 测试函数TestXxx(t *testing.T)的参数是testing.T, Xxx的首字母必须是大写. 比如Testintdiv是不对的.
+ 通过t.Log或者fmt来输出信息.

使用运行  

+ go test -run='TestGetCurrentTime' (不用指定文件)
+ go test 依次运行当前目录下所有*_test文件中的所有TestXxx(t *testing.T)函数.
