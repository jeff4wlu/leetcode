package leet1

func TwoSum(arr []int, target int) (re [2]int) {
	hm := make(map[int]int, 20)
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
