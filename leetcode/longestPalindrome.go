package leetcode


func longestPalindrome(str string) string {
	if str == "" {
		return ""
	}

	var t = len(str)
	var s,e,s1,e1,ms,me = 0,0,0,0,0,0


	for i,_ := range str {
		
		
		for p := 0 ; p <= i ;p ++{

			if (p == 0 || e > s) && i - p-1 >=0 && i+p < t && str[i-p-1] == str[i+p] {
				s ,e = i-p-1,i+p
			} else if e > s {
				if e -s > me - ms {
					me,ms = e,s
				}

				//reset
				s,e = 0,0
			}

			if (p == 0 || e1 > s1) && i -p -1 >=0 &&  i+p+1 < t && str[i-p-1] == str[i+p+1] {
				s1,e1 = i-p-1,i+p+1
			} else if e1 > s1 {
				if e1-s1 > me - ms {
					me,ms = e1,s1
				}

				//reset
				e1,s1 = 0,0
			}

			//因为这个语句，影响了性能
			// fmt.Printf("%v %v %v %v %v\n",p,s,e,s1,e1)

		}
	
	}
	return str[ms:me+1]
	
}