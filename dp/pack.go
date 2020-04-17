package main

// package main

import (
	"fmt"
	"time"
)

var caches = make([]int,200,200)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}


func packM(values []int,capacitys []int,capacity int) int{
	if caches[capacity] > 0  {
		return caches[capacity]
	}

	var max int 
	for i,v := range values {
		var sum int
		if capacitys[i] <= capacity {

			sum = v + packM(values,capacitys,capacity - capacitys[i])
		} else {
			continue
		}

		if sum > max {
			max = sum
		}
	}

	if caches[capacity] == 0 {
		caches[capacity] = max
	}
	return max ;
}

func pack(values []int,capacitys []int,capacity int) int{
	if values == nil || len(values) == 0 {
		return 0
	}

	var max int 

	for i,v := range values {
		var sum int
		if capacitys[i] <= capacity {

			sum = v + pack(values,capacitys,capacity - capacitys[i])

		} else {
			continue
		}

		if sum > max {
			max = sum
		}
		
	}

	return max ;
}


func topDown(values []int,capacitys []int,capacity int) int{
	if capacity < len(caches) {
		return caches[capacity]
	} 

	var max int
	for i:= 0 ; i <= capacity ; i++ {
		sum := topDown(values,capacitys,i) + topDown(values,capacitys,capacity - i)
		if sum > max {
			max = sum
		}
	}

	return max
}


func main() {
	values := []int{1,2, 5, 3} //物品大小
	capacitys := []int{1,2, 3, 4} //物品价值
	// goods := [][]int{
	// 	[]int{7, 2},
	// 	[]int{5, 3},
	// 	[]int{8, 4},
	// }

	
	var now = time.Now().Unix()
	fmt.Println(packM(values, capacitys, 199))
	fmt.Println(time.Now().Unix()-now)
	// now = time.Now().Unix()
	// fmt.Println(pack(values, capacitys, 25))
	// fmt.Println(time.Now().Unix()-now)
	
}
