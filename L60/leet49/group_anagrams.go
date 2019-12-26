package leet49

//An array isn't immutable by nature; you can't make it constant.
//Note the [...] instead of []: it ensures you get a (fixed size) array instead of a slice. So the values aren't fixed but the size is.
var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

func GroupAnagrams(anagrams []string) [][]string {

	n := len(anagrams)
	dict := make(map[int][]string)

	for i := 0; i < n; i++ {
		tmp := []byte(anagrams[i])
		key := 1
		for j := 0; j < len(tmp); j++ {
			key *= primes[tmp[j]-'a']
		}

		if _, ok := dict[key]; !ok {
			dict[key] = []string{}
		}
		dict[key] = append(dict[key], anagrams[i])

	}

	//使用以下方式而不是用make，这样的话不用分配多余元素空间，但耗时一点
	res := [][]string{}

	for _, v := range dict {
		res = append(res, v)
	}
	return res
}
