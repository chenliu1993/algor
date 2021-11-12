/*
 * @lc app=leetcode.cn id=1702 lang=golang
 *
 * [1702] 修改后的最大二进制字符串
 */

// @lc code=start
func maximumBinaryString(binary string) string {
	var (
		zeros     int
		zeroBegin int
	)
	zeros = 0
	zeroBegin = -1
	bytesBinary := []byte(binary)
	for i := 0; i < len(bytesBinary); i++ {
		if bytesBinary[i] == '0' {
			zeros++
			if zeroBegin == -1 {
				zeroBegin = i
			}
			bytesBinary[i] = '1'
		}
	}

	if zeroBegin == -1 {
		return binary
	}
	// Begin Construct
	bytesBinary[zeroBegin+zeros-1] = '0'
	return string(bytesBinary)
}

// @lc code=end

