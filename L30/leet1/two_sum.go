package leet1

func TwoSum(arr []int, target int) (re []int) {
	hm := make(map[int]int, 20)
	re = []int{0, 0}
	for i, v := range arr {
		t := target - arr[i]
		if _, ok := hm[arr[i]]; ok {
			continue
		}
		if _, ok := hm[t]; ok {
			re[0] = hm[t]
			re[1] = i
		}
		hm[v] = i
	}
	return
}

//指针碰撞算法
func TwoSum1(arr []int, target int) (re []int) {

	n := len(arr)
	i, j := 0, n-1
	for i < j {
		sum := arr[i] + arr[j]
		if sum == target {
			return []int{i, j}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return []int{i, j}
}
