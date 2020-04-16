
package sort

func quicksort(data []int,s int,e int){

	if e > s {
		c := partition(data,s,e)
		quicksort(data,s,c-1)
		quicksort(data,c+1,e)
	}


}

func partition(data []int,s int,e int) int {
	p := data[e]

	//l(<),m (>)
	var c = s
	for i := s ;i<= e-1 ;i++{
		if data[i] <= p {

			swap(data,c,i)
			c ++
		} 
	}

	swap(data,c,e)
	return c
}

