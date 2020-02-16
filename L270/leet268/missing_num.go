package leet268

//注意题意，这个nums数组并不是已经排序的。所以只能使用等差数列的求和公式
func MissingNum(nums []int) int {

	n := len(nums)
	var sum int
	for _, v := range nums {
		sum += v
	}
	return n*(n+1)/2 - sum
}

//使用bit操作的异或，除一个数是一次外，其他都是两次
func MissingNum1(nums []int) int {

	n := len(nums)
	var res int //本身就已经是0
	for i := 1; i <= n; i++ {
		res ^= i
	}
	for _, v := range nums {
		res ^= v
	}

	return res
}
