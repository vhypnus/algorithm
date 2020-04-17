package main

// package main

import "fmt"

var caches = []int{0,0,2,5,5,7,10,10,12,15,15}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}


func packM(values []int,capacitys []int,capacity int) int{
	if values == nil || len(values) == 0 {
		return 0
	}

	var max int 

	for i,v := range values {
		var sum int
		if capacitys[i] <= capacity {
				
			if capacity-capacitys[i] < len(caches)  {
				sum = v + caches[capacity-capacitys[i]]
			} else {
				sum = v + pack(values,capacitys,capacity - capacitys[i])
			}
		} else {
			continue
		}

		if sum > max {
			max = sum
		}
		
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

	var max int
	for i:= capacity ; i > 0 ; i++ {
		if caches[i] > 0 {
			return caches[i]
		}else {
			for c,v := range caches {
				if c > capacity {
					break
				}
				var sum int 

				sum = v + pack(values,capacitys,capacity - c) 

				if sum > max {
					max = sum
				}
			}
		}
		 
	}

	return max
}


func main() {
	values := []int{2, 5, 3} //物品大小
	capacitys := []int{2, 3, 4} //物品价值
	// goods := [][]int{
	// 	[]int{7, 2},
	// 	[]int{5, 3},
	// 	[]int{8, 4},
	// }

	

	fmt.Println(packM(values, capacitys, 60))
	fmt.Println(pack(values, capacitys, 60))
	
}
