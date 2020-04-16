/*

*/

package leetcode


func convert(str string, numRows int) string {

	if numRows == 0 || numRows == 1{
		return str
	}
	var t = len(str)
	// fmt.Println(t)

	var data = make([]byte,0,t)
	for i := 0 ;i < numRows ; i ++{
		
		var s,e int
		for p := 0; ;p++{

			s,e = e,e+numRows-1	

			if i == 0 {
				if p & 1 == 0 && s+i < t{
					data = append(data,str[s+i])
				}
			} else if i == numRows - 1 {
				if p & 1 == 0 && s +i < t {
					data = append(data,str[s+i])	
				}
			} else {
				if p & 1 == 0 && s+i < t{
					data = append(data,str[s+i])
				} else if (p << 1) & 2 == 2 && s+numRows - i -1 < t {
					data = append(data,str[s+numRows - i -1])
				}
				
			}

			if e >= t {
				break 
			}

		}
		
	}

	var b strings.Builder
	b.Write(data)
	return b.String()
	
}