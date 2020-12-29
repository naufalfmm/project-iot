package word

import (
	"github.com/naufalfmm/project-iot/common/consts"
)

func FirstLetterWord(str string) string {
	result := ""

	v := true
	for i := 0; i < len(str); i++ {
		if str[i] == consts.AsciiCodeSpace {
			v = false
			continue
		}

		if v == true && str[i] != consts.AsciiCodeSpace {
			result += string(ToLowerCase(str[i]))
			v = false
		}
	}

	return result
}
