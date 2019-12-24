package leet43

import "fmt"

func MultiplyString(a, b string) string {

	aLen := len(a)
	bLen := len(b)
	if aLen > 110 || bLen > 110 {
		return ""
	}

	aArr := []byte(a)
	bArr := []byte(b)

	res := make([]byte, aLen+bLen)

	for i, _ := range res {
		res[i] = '0'
	}

	for i := bLen - 1; i >= 0; i-- {
		for j := aLen - 1; j >= 0; j-- {
			product := (aArr[j] - '0') * (bArr[i] - '0')
			p := (res[i+j]-'0')*10 + (res[i+j+1] - '0') + product

			//p1 := product/10 + (res[i+j] - '0')   //高位
			//p2 := product%10 + (res[i+j+1] - '0') //低位
			//p := p1*10 + p2                       //总和
			if p/100 == 1 { //加法进位                     //3位数
				res[i+j-1] = res[i+j-1] + p/100
				p %= 100 //去掉百位
			}
			res[i+j+1] = p%10 + '0'
			res[i+j] = p/10 + '0'
			s := string(res)
			fmt.Printf("%s", s)
		}
	}

	for i := 0; i < len(res); i++ {
		if res[i] != '0' {
			return string(res[i:])
		}
	}
	return string(res[:])
}
