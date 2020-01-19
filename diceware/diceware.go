package diceware

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Generate(lang string, size int32) {
	var words = ""

	for i := 1; i <= int(size); i++ {
		index := findDicewareWordIndex()
		word := findDicewareWord(index, lang)
		words = words + word + " "
	}

	fmt.Println(words)
}

func findDicewareWordIndex() string {
	var number = ""
	for j := 1; j <= 5; j++ {
		number = number + strconv.FormatInt(throwDice(), 10)
	}
	return number
}

func throwDice() int64 {
	var number int64 = 0

	for number == 0 {
		nBig, err := rand.Int(rand.Reader, big.NewInt(7))
		if err != nil {
			panic(err)
		}
		number = nBig.Int64()
	}

	return number
}

func findDicewareWord(number string, lang string) string {
	file, err := os.Open("diceware_words_" + lang + "/" + number + ".txt")

	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
		transformedWord, _, _ := transform.String(t, word)
		return transformedWord
	}
	return ""
}
