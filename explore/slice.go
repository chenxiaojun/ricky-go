package main

import "fmt"

//func main() {
//	// 使用make创建一个切片, 长度是5，容量是10
//	// 容量必须 > 长度， 我们是不能创建长度大于容量的切片的
//	//slice := make([]int, 5, 10)
// 	//fmt.Println(slice)
//
// 	// 还有一种创建切片的方式，是使用字面量
// 	//slice := []int{ 1, 2, 3, 4, 5 }
//
// 	// 数组
// 	//arr := [5]int{4:1}
// 	//切片
// 	//ali := []int{4:1}
// 	//fmt.Println(arr, ali)
//
// 	//slice1 := []int{1, 2, 3, 4, 5}
//	//newSlice := slice1[1:3] // 新的切片和原有的切片公用一个底层数组
// 	//fmt.Println(newSlice)
// 	//newSlice[0] = 8
// 	//fmt.Println(slice1) // 1, 8, 3, 5, 5
//
// 	// 切片是一个动态数组，它可以按需增长，使用内置的append函数即可
//
// 	slice := []int{ 1, 2, 3, 4, 5, 6 }
// 	newSlice := slice[3:6]
//	fmt.Println(newSlice)
//	fmt.Println(slice)
// 	fmt.Println(len(slice), cap(slice))
// 	fmt.Println(len(newSlice), cap(newSlice))
//
// 	newSlice = append(newSlice, 10)
// 	fmt.Println(newSlice)
// 	fmt.Println(slice)
// 	fmt.Println(len(slice), cap(slice))
// 	fmt.Println(len(newSlice), cap(newSlice)) // 3, 3
//
//	newSlice = append(newSlice, 20)
//	fmt.Println(newSlice)
//	fmt.Println(slice)
//	fmt.Println(len(slice), cap(slice))
// 	fmt.Println(len(newSlice), cap(newSlice)) // 4, 6
//
//	newSlice = append(newSlice, 30)
//	fmt.Println(newSlice)
//	fmt.Println(slice)
//	fmt.Println(len(slice), cap(slice))
// 	fmt.Println(len(newSlice), cap(newSlice)) // 5, 6
//
// 	// 尽量让新的slice cap和len相同，这样就不会改变原油数组的值。
//}

//func main() {
//	slice := []int{1, 2, 3, 4, 5}
//	for i, v := range slice {
//		fmt.Printf("索引: %d, 值:%d\n", i, v)
//	}
//}

//func main() {
//	slice := []int{1, 2, 3, 4, 5}
//	fmt.Printf("%p\n", &slice)
//
//	modify(slice) // 切片在函数间的传递是复制的。
//	fmt.Println(slice)
//}
//
//func modify(slice []int) {
//	fmt.Printf("%p\n", &slice) // 切片的地址变了
//	slice[1] = 10 // 影响了main函数里面的数组
//}

func main() {
	slice := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)

	modify(slice) // 数组在函数间的传递是值传递的。
	fmt.Println(slice)
}

func modify(slice [5]int) {
	fmt.Printf("%p\n", &slice) // 切片的地址变了
	slice[1] = 10 // 不改变main里面slice的值
}