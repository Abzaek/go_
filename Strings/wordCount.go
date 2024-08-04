package Str

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func removePunctuation(str *string) string {
	var builder strings.Builder
	for _, s := range *str {
		if unicode.IsLetter(s) || unicode.IsSpace(s) {
			builder.WriteRune(s)
		}
	}
	return builder.String()
}

func WordCounter(shouldPrint bool) []string {

	reader := bufio.NewReader(os.Stdin)
	var CountLookup map[string]int = make(map[string]int)

	fmt.Printf("enter your string: ")
	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("error found: %s\n", err)
		return make([]string, 0)
	}

	str2 := removePunctuation(&str)
	strArr := strings.Fields(str2)

	if !shouldPrint {
		return strArr
	}

	for _, str := range strArr {
		CountLookup[strings.ToLower(str)] += 1
	}

	fmt.Println(CountLookup)
	return make([]string, 0)
}
