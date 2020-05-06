package geek

import(
	"log"
)
/*
 * https://www.geeksforgeeks.org/array-rotation/
 * l 奇偶不同
 */
func Rotation(ar[]int,d int) {

 	var l = len(ar)
 	d = d % l
 	var max = gcd(d,l) 	

 	for  i := 0 ; i < max ; i ++ {

 		temp,j,k := ar[i],i,0
 		for {
 			k = j + d 
 			if k >= l {
 				k = k -l 
 			}

 			if k == i {
 				break
 			}

 			ar[j] = ar[k]
 			j = k
 		}

 		ar[j] = temp
 	}
}

func gcd(a int ,b int) int{
	if b == 0 {
		return a 
	} else {
		return gcd(b,a %b)
	}
}


func Largest(ar []int ,s,e int) int {
	log.Printf("s e %v %v",s,e)
	if s > e {
		// 
		return -1
	} 
	if s == e {
		return s 
	}

	if ar[e] > ar[s] {
		return len(ar)-1
	}
	
	var middle = (s+e)/2
	log.Printf("middle %v",middle)

	if middle < e && ar[middle] > ar[middle+1] {
		return middle
	}

	if middle > s && ar[middle] < ar[middle-1] {
		return middle -1
	}

	if ar[middle] > ar[e] {	
		return Largest(ar,middle+1,e)
	} else{
		return Largest(ar,s,middle-1)
	}
}



func Binary_search(ar []int,v,s ,e  int) int{
	var max = len(ar)-1
	var middle = middle(max,s,e)

	if ar[middle] == v{
		return middle
	}else {
		if ar[middle] > v {
			e = pre(max,middle)
		} else {
			
			s = next(max,middle)
		}
	}

	log.Printf("s e   middle  %v %v %v" ,s,e,middle)
	return Binary_search(ar,v,s,e)
	// return -1

}

func pre(max,index int) int {
	if index -1 < 0 {
		return max
	} else {
		return index -1
	}
}

func next(max,index int) int{
	if index + 1 > max {
		return 0
	} else {
		return index +1
	}
}

func middle(max ,s ,e int) int{

	var m ,l int
	
	if s > e {
		l = max -s +e
		log.Printf("l %v %v",l,l/2)
		if s + l/2 > max {
			m = l/2 - (max -s )
		} else{
			m = s + l/2
		}
	} else {
		l = e - s
		m = s +l/2 
	}
	log.Printf("middle %v",m)
	return m
}


func Quicksort(ar []int,s,e int) {

}

func Partition(ar []int,s,e int) {
	var j int = s

	var pivot = ar[e]
	for k := s ;k < e ; k++ {
		if ar[k] < pivot {
			Swap(ar,j,k)
			j = j+1
		} 
	}

	Swap(ar,j,e)
}

func Swap(ar []int,i,j int) {
	temp := ar[i]
	ar[i] = ar[j]
	ar[j] = temp
}