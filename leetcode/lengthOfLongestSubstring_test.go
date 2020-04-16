package leetcode

import(
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	 var str = "pwwekw"
	 var max = lengthOfLongestSubstring(str)
	 fmt.Println(max)


	 str = " "
	 max = lengthOfLongestSubstring(str)
	 fmt.Println(max)


	 str = "abcdefgddbb"
	 max = lengthOfLongestSubstring(str)
	 fmt.Println(max)

	 str = "abcdefgh"
	 max = lengthOfLongestSubstring(str)
	 fmt.Println(max)
	
}