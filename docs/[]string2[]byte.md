##Go语言实现将[]string转化为[]byte
假设你想实现Go语言的string数组到byte数组的转化。演示函数如下所示：

	func convertStringsToBytes(){
		stringContent := []string{"通知中心","perfect!"}
		byteContent := "\x00"+ strings.Join(stringContent, "\x02\x00")  // x20 = space and x00 = null
		fmt.Println([]byte(byteContent))
		fmt.Println(string([]byte(byteContent)))
	}


完整代码：

	package main
	
	 import (
	 	"fmt"
	 	"strings"
	 )
	
	 func convert(){
	 	stringSlice := []string{"通知中心","perfect!"}
	
	 	stringByte := "\x00" + strings.Join(stringSlice, "\x20\x00") // x20 = space and x00 = null
	
	 	fmt.Println([]byte(stringByte))
	 	
	 	fmt.Println(string([]byte(stringByte)))
	 }
	 func main() {
		convert()
	 }

运行结果：

	[0 233 128 154 231 159 165 228 184 173 229 191 131 2 0 112 101 114 102 101 99 116 33]
	 通知中心  perfect!

上面是最简单的方法，还有另外一种方式可以实现同样的效果。它主要要使用编码机制实现。

	 package main
	
	 import (
	 	"bytes"
	 	"encoding/gob"
	 	"fmt"
	 )
	
	func convert（）{
		stringSlice := []string{"通知中心","perfect!"}
	
	 	buffer := &bytes.Buffer{}
	
	 	gob.NewEncoder(buffer).Encode(stringSlice)
	 	byteSlice := buffer.Bytes()
	 	fmt.Printf("%q\n", byteSlice)
	
	 	fmt.Println("---------------------------")
	
	 	backToStringSlice := []string{}
	 	gob.NewDecoder(buffer).Decode(&backToStringSlice)
	 	fmt.Printf("%v\n", backToStringSlice)
	}
	
	 func main() {
		convert()
	 }




Welcome you!
![weixing publiv account](http://img.blog.csdn.net/20160424104206329)