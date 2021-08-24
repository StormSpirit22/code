package main

func main() {
	
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letterMap := make(map[byte]string)
	letterMap['2'] = "abc"
	letterMap['3'] = "def"
	letterMap['4'] = "ghi"
	letterMap['5'] = "jkl"
	letterMap['6'] = "mno"
	letterMap['7'] = "pqrs"
	letterMap['8'] = "tuv"
	letterMap['9'] = "wxyz"

	var res []string
	db := []byte(digits)
	var letters []string
	for _, d := range db {
		letters = append(letters, letterMap[d])
	}

	var backtrack func([]byte, int)
	backtrack = func(track []byte, start int) {
		if len(track) == len(digits) {
			tmp := make([]byte, len(digits))
			copy(tmp, track)
			res = append(res, string(tmp))
			return
		}

		for i := start; i < len(letters); i++ {
			for j := 0; j < len(letters[i]); j++ {
				track = append(track, letters[i][j])
				backtrack(track, i+1)
				track = track[:len(track)-1]
			}
		}
	}
	backtrack([]byte{}, 0)
	return res
}

