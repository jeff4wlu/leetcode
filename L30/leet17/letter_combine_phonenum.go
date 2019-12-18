package leet17

func LetterComPhonenumDFS(digits string, dict map[string]string, lvl int, word string, res *[]string) {
	if lvl == len(digits) {
		*res = append(*res, word)
		return
	}

	s := dict[digits[lvl:lvl+1]]

	for i := 0; i < len(s); i++ {
		LetterComPhonenumDFS(digits, dict, lvl+1, word+s[i:i+1], res)
	}

}

func LetterComPhonenum(digits string) []string {
	if digits == "" {
		return nil
	}

	dict := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var res []string

	LetterComPhonenumDFS(digits, dict, 0, "", &res)
	return res
}
