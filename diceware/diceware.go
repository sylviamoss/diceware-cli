package diceware

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
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

	var words string
	for i := 1; i <= int(generateConfig.Size); i++ {
		index, err := findDicewareWordIndex()
		if err != nil {
			return err
		}
		word, err := findDicewareWord(index, generateConfig.Lang)
		if err != nil {
			return err
		}
		words = words + word + separator
	}
	words = words[:len(words)-len(separator)]

	if generateConfig.Pbcopy || generateConfig.Hide {
		cmd := fmt.Sprintf("echo %s | pbcopy", words)
		if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
			return fmt.Errorf("error copying passphrase: %s", err.Error())
		}
		fmt.Println("Password copied!")
	}

	if generateConfig.Hide {
		return nil
	}

	if words == "" {
		return fmt.Errorf("unable to generate passphrase.")
	}

	fmt.Println(words)
	return nil
}

func findDicewareWordIndex() (string, error) {
	var number string
	for j := 1; j <= 5; j++ {
		dice, err := throwDice()
		if err != nil {
			return number, err
		}
		number = number + strconv.FormatInt(dice, 10)
	}
	return number, nil
}

func throwDice() (int64, error) {
	var number int64
	for number == 0 {
		nBig, err := rand.Int(rand.Reader, big.NewInt(7))
		if err != nil {
			return number, fmt.Errorf("error while throwing the dice: %s", err.Error())
		}
		number = nBig.Int64()
	}

	return number, nil
}

func findDicewareWord(number string, lang string) (string, error) {
	wordPath := "diceware_words_" + lang + "/" + number + ".txt"
	word, err := wordsBox.FindString(wordPath)
	if err != nil {
		word, err = findCustomDicewareWord(wordPath)
		if err != nil {
			return "", fmt.Errorf("unable to find word for index %s. err: %s", number, err.Error())
		}
	}

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	transformedWord, _, err := transform.String(t, word)
	if err != nil {
		return "", fmt.Errorf("unable to remove special characters from %s. err: %s", word, err.Error())
	}
	return transformedWord, nil
}

func findCustomDicewareWord(wordPath string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	file, err := os.Open(home + "/.diceware/" + wordPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", fmt.Errorf("couldn't read word from custom dictionary.")
}
