package convert_util

func SliceDeduplicate(src []string) []string {
	srcMap := make(map[string]bool)
	for _, item := range src {
		srcMap[item] = true
	}

	var rst []string

	for k := range srcMap {
		rst = append(rst, k)
	}
	return rst
}

func SliceIsContain(src []string, key string) bool {
	for _, item := range src {
		if key == item {
			return true
		}
	}
	return false
}

// 两个slice不考虑元素顺序的情况下是否相同
func SliceIsSameWithoutSequence(src []string, src2 []string) bool {
	if len(src) != len(src2) {
		return false
	}

	sameFlag := true
	for _, item := range src {
		if SliceIsContain(src2, item) == false {
			sameFlag = false
			break
		}
	}

	return sameFlag
}
