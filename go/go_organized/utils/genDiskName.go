package utils

// GetDiskName get the name of the disk given the number of it
func GetDiskName(curNum int) string {
	if curNum < 0 {
		curNum = -curNum
	}
	var diskName string
	for {
		diskName = string(rune('a'+curNum%26)) + diskName
		curNum = curNum / 26
		if curNum == 0 {
			break
		}
		curNum--
	}

	return "sd" + diskName + "1"
}
