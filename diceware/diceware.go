package diceware

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os/exec"
	"strconv"
	"unicode"

	"github.com/gobuffalo/packr"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type GenerateConfig struct {
	Lang      string
	Size      int32
	Pbcopy    bool
	Hide      bool
	Separator string
}

var wordsBox packr.Box

func Generate(generateConfig GenerateConfig, box packr.Box) error {
	wordsBox = box

	separator := generateConfig.Separator
	if separator == "none" {
		separator = ""
	}

	var words = ""
	for i := 1; i <= int(generateConfig.Size); i++ {
		index := findDicewareWordIndex()
		word := findDicewareWord(index, generateConfig.Lang)
		words = words + word + separator
	}
	words = words[:len(words)-len(separator)]

	if generateConfig.Pbcopy || generateConfig.Hide {
		cmd := fmt.Sprintf("echo %s | pbcopy", words)
		if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
			return err
		}
		fmt.Println("Password copied!")
	}

	if generateConfig.Hide {
		return nil
	}

	if words == "" {
		fmt.Println("Unable to generate passphrase.")
		return nil
	}

	fmt.Println("-------------------")
	fmt.Println(words)
	fmt.Println("-------------------")
	return nil
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
			panic(err) // TODO don't panic!
		}
		number = nBig.Int64()
	}

	return number
}

func findDicewareWord(number string, lang string) string {
	word, err := wordsBox.FindString("diceware_words_" + lang + "/" + number + ".txt")
	if err != nil {
		return ""
	}

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	transformedWord, _, _ := transform.String(t, word)
	return transformedWord
}
