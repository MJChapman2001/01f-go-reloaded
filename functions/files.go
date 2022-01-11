package functions

import (
	"os"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func ProcessFile(s string, fileName string) {
	newContent := ""

	content, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}

	contentSlice := strings.Fields(string(content))

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

	newContent = strings.Join(contentSlice, " ")

	contentSlice = strings.Fields(newContent)

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
		case 'a':
			vowels := []rune{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}

			nextSlice := []rune(contentSlice[k+1])

			for _, x := range vowels {
				if x == nextSlice[0] {
					contentSlice[k] += "n"
				}
			}
		case 'A':
			vowels := []rune{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}

			nextSlice := []rune(contentSlice[k+1])

			for _, x := range vowels {
				if x == nextSlice[0] {
					contentSlice[k] += "n"
				}
			}
		}
	}

	newContent = strings.Join(contentSlice, " ")

	contentSlice = strings.Fields(newContent)
	newContent = strings.Join(contentSlice, " ")

	os.WriteFile(fileName, []byte(newContent), 0666)
}