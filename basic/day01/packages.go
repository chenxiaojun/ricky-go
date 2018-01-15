// 程序的入口文件是包 main
package main

// 使用了打包导入的方式
import (
	"fmt"
	"math/rand"
	"math"
)

// 由于这段代码的运行环境是固定的，所以rand.Intn总是会返回相同的数字
func main () {
	// 返回一个随机的整数
	fmt.Println("rand int number", rand.Intn(10))
	fmt.Println("rand int number", rand.Intn(10))

	// 返回一个64位浮点数 0.0 < f < 1.0
	fmt.Println("rand float number", rand.Float64())
	// 利用这个技巧生成其它范围的随机数
	fmt.Println((rand.Float64()*5)+5)

	//要让伪随机数生成器有确定性，可以给它一个明确的种子
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100))

	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}

