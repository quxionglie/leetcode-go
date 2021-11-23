package problem100

/**
28. 实现 strStr()

实现strStr()函数。
给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
说明：
当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。
链接：https://leetcode-cn.com/problems/implement-strstr
*/
func strStr(haystack string, needle string) int {
	//暴力破解
	if needle == "" {
		return 0
	}
	needleLength := len(needle)
	haystackLength := len(haystack)
	for i := 0; i < haystackLength; i++ {
		var j int
		for j = 0; j < needleLength && i+j < haystackLength; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == needleLength {
			return i

		}
	}

	return -1
}
