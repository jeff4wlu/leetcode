package leet137

func SingleNum(in []int) int {

	var a, b int

	for _, v := range in {
		tmp := a&(^b)&(^v) | (^a)&b&v
		b = (^a)&b&(^v) | (^a)&(^b)&v
		a = tmp
	}
	return b
}

//思维最直接的解法。根据题意，除一个数外，其他都是三个即三的倍数。
//一个整数32位，这个数组所有数相加，即32位数的每一bit都是三的倍数，除了要被求出的那个。
//所以，每一bit对3取模后余数就是那个要求的数。
func SingleNum1(nums []int) int {

	var res int
	for i:=31;i>=0;i--{
		var tmp int
		res = res<<1 //整理回原来的位置
		for j:=0;j<len(nums);j++{
			tmp+= (nums[j]>>uint(i))&1 //先求出数组中每个数在此bit中是0或1，然后求和
		}
		res |= tmp%3//求模数,再把结果的每一位拼起来
	}
	return res
}

//思路，两个相同数用异或可以直接消掉。三个相同数时使第三个数出现后消掉为0，怎么实现呢？
//  a b
//0 0 0
//1 x 0
//2 0 x
//3 0 0
//使用两个变量a和b实现上图的状态转换，最后a变量就是只出现一次的那个数。使用到如下位操作
// x^0=x, x&~x=0
func SingleNum2(nums []int) int {

	var a,b int
	for i:=0;i<len(nums);i++{
		a = (a^nums[i])& (^b)
		b=(b^nums[i])&(^a)
	}
	return a
	
}

