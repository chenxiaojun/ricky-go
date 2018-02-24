package test

import "testing"

/**
表组测试，就是一个里面有很多的测试放在一起进行测试
*/
func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}

	sum = Add(4, 5)
	if sum == 9 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}
