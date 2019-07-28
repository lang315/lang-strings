package lstrings

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
	"unicode/utf8"
)

type optionAccent struct {
	destinationCharacters []string
	sourceCharacters      [] string
	lengthSource          int
}

func RemoveAccent(s string) string {
	return removeAccent(s)

	// destinationCharacters, _ := stringToRune(`AAAAEEEIIOOOOUUYaaaaeeeiioooouuyAaDdIiUuOoUuAaAaAaAaAaAaAaAaAaAaAaAaEeEeEeEeEeEeEeEeIiIiOoOoOoOoOoOoOoOoOoOoOoOoUuUuUuUuUuUuUuYyy`)
	// sourceCharacters, lengthSource := stringToRune(`ÀÁÂÃÈÉÊÌÍÒÓÔÕÙÚÝàáâãèéêìíòóôõùúýĂăĐđĨĩŨũƠơƯưẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọỎỏỐốỒồỔổỖỗỘộỚớỜờỞởỠỡỢợỤụỦủỨứỪừỬửỮữỰựỸỹỳ`)
	//
	// option := optionAccent{
	// 	destinationCharacters: destinationCharacters,
	// 	sourceCharacters:      sourceCharacters,
	// 	lengthSource:          lengthSource,
	// }
	//
	// var buffer bytes.Buffer
	//
	// for _, runeValue := range s {
	// 	buffer.WriteString(removeAccentChar(string(runeValue), option))
	//
	// }

}

func removeAccent(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	ss, _, err := transform.String(t, s)
	if err != nil {
		return ""
	}

	return ss
}

func stringToRune(s string) ([]string, int) {
	l := utf8.RuneCountInString(s)

	var texts = make([]string, l+1)

	var index = 0

	for _, runeValue := range s {
		texts[index] = string(runeValue)
		index++
	}
	return texts, l

}

func binarySearch(sortedArray []string, key string, low int, high int) int {
	var middle = (low + high) / 2
	if high < low {
		return -1
	}

	if key == sortedArray[middle] {
		return middle
	} else if key < sortedArray[middle] {
		return binarySearch(sortedArray, key, low, middle-1)
	} else {
		return binarySearch(sortedArray, key, middle+1, high)
	}
}

func removeAccentChar(ch string, option optionAccent) string {
	var index = binarySearch(option.sourceCharacters, ch, 0, option.lengthSource)
	if index >= 0 {
		ch = option.destinationCharacters[index]
	}

	return ch
}
