/*
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

解题思路:
1.

时间复杂度:
空间复杂度:
*/
package leetcode


func lengthOfLongestSubstring(s string) int {
	//可以优化为128个bit
    var arr []int = make([]int,127,127)
    // fmt.Println(arr)
    var len,max ,mark int =0,0,0
    for i,c := range s {
        if arr[int(c)] > mark {
        	len = i - arr[int(c)]+1

            mark = arr[int(c)]
            arr[int(c)] = i+1
        }else {
            arr[int(c)] = i + 1
            len += 1
        }

        if len > max {
       		max = len
        }

        // fmt.Printf("%v %v %v %v \n",max,len,mark, i)
    }

    return max 
}