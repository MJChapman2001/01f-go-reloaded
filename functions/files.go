package functions

import (
	"os"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func Commands(s string) string {
	contentSlice := strings.Fields(s)

	for i, j := range contentSlice {
		switch j {
		case "(hex)":
			contentSlice[i-1] = Hex(contentSlice[i-1])
			contentSlice[i] = ""
		case "(bin)":
			contentSlice[i-1] = Bin(contentSlice[i-1])
			contentSlice[i] = ""
		case "(up)":
			contentSlice[i-1] = Up(contentSlice[i-1])
			contentSlice[i] = ""
		case "(up,":
			val, err := strconv.Atoi(strings.TrimSuffix(contentSlice[i+1], ")"))
			if err != nil {
				log.Fatal(err)
			}

			for val > 0 {
				contentSlice[i-val] = Up(contentSlice[i-val])
				val--
			}
			contentSlice[i], contentSlice[i+1] = "", ""
		case "(low)":
			contentSlice[i-1] = Low(contentSlice[i-1])
			contentSlice[i] = ""
		case "(low,":
			val, err := strconv.Atoi(strings.TrimSuffix(contentSlice[i+1], ")"))
			if err != nil {
				log.Fatal(err)
			}

			for val > 0 {
				contentSlice[i-val] = Low(contentSlice[i-val])
				val--
			}
			contentSlice[i], contentSlice[i+1] = "", ""
		case "(cap)":
			contentSlice[i-1] = Cap(contentSlice[i-1])
			contentSlice[i] = ""
		case "(cap,":
			val, err := strconv.Atoi(strings.TrimSuffix(contentSlice[i+1], ")"))
			if err != nil {
				log.Fatal(err)
			}

			for val > 0 {
				contentSlice[i-val] = Cap(contentSlice[i-val])
				val--
			}
			contentSlice[i], contentSlice[i+1] = "", ""
		}
	}

	newContent := strings.Join(contentSlice, " ")
	contentSlice = strings.Fields(newContent)
	newContent = strings.Join(contentSlice, " ")

	return newContent
}

func Punctuation(s string) string {
	contentSlice := strings.Fields(s)

	for k, l := range contentSlice {
		splitL := []rune(l)

		switch splitL[0] {
		case '.':
			temp := 1

			for contentSlice[k-temp] == "" {
				temp++
			}

			contentSlice[k-temp] += "."
			splitL[0] = ' '
			contentSlice[k] = string(splitL)
		case ',':
			temp := 1

			for contentSlice[k-temp] == "" {
				temp++
			}

			contentSlice[k-temp] += ","
			splitL[0] = ' '
			contentSlice[k] = string(splitL)
		case '!':
			temp := 1

			for contentSlice[k-temp] == "" {
				temp++
			}

			contentSlice[k-temp] += "!"
			splitL[0] = ' '
			contentSlice[k] = string(splitL)
		case '?':
			temp := 1

			for contentSlice[k-temp] == "" {
				temp++
			}

			contentSlice[k-temp] += "?"
			splitL[0] = ' '
			contentSlice[k] = string(splitL)
		case ';':
			temp := 1

			for contentSlice[k-temp] == "" {
				temp++
			}

			contentSlice[k-temp] += ";"
			splitL[0] = ' '
			contentSlice[k] = string(splitL)
		case ':':
			temp := 1

			for contentSlice[k-temp] == "" {
				temp++
			}

			contentSlice[k-temp] += ":"
			splitL[0] = ' '
			contentSlice[k] = string(splitL)
		}
	}

	newContent := strings.Join(contentSlice, " ")
	contentSlice = strings.Fields(newContent)
	newContent = strings.Join(contentSlice, " ")

	return newContent
}

func LetterA(s string) string {
	contentSlice := strings.Fields(s)

	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}
	
	for m := 0; m < len(contentSlice)-1; m++ {
		if contentSlice[m] == "a" || contentSlice[m] == "A" {
			tempSlice := []rune(contentSlice[m+1])

			for _, n := range vowels {
				if tempSlice[0] == n {
					contentSlice[m] += "n"
				}
			}
		}
	}

	newContent := strings.Join(contentSlice, " ")

	return newContent
}

func Quote(s string) string {
	contentSlice := strings.Fields(s)
	quoteCount := 1

	for index, data := range contentSlice {
		if data == "'" {
			if quoteCount%2 == 0 {
				temp := 1

				for contentSlice[index-temp] == "" {
					temp++
				}

				contentSlice[index-temp] += "'"
				contentSlice[index] = ""
			} else {
				if index+1 < len(contentSlice)-1 {
					contentSlice[index+1] = "'" + contentSlice[index+1]
					contentSlice[index] = ""
				}
			}

			quoteCount++
		}
	}

	newContent := strings.Join(contentSlice, " ")
	contentSlice = strings.Fields(newContent)
	newContent = strings.Join(contentSlice, " ")

	return newContent
}

func ProcessFile(s string, fileName string) {

	content, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)

	data = Commands(data)
	data = Punctuation(data)
	data = LetterA(data)
	data = Quote(data)

	os.WriteFile(fileName, []byte(data), 0666)
}