package diceware

import (
	"bufio"
	"crypto/rand"
	"embed"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Config struct {
	Lang         string
	Size         int32
	Pbcopy       bool
	Hide         bool
	Lower        bool
	RemoveNumber bool
	Separator    string
}

//go:embed words
var words embed.FS

func (c *Config) Generate() error {
	separator := c.Separator
	if separator == "none" {
		separator = ""
	}

	numberedIndex, err := rand.Int(rand.Reader, big.NewInt(int64(c.Size)))
	if err != nil {
		return err
	}

	var words string
	for i := 0; i < int(c.Size); i++ {
		index, err := findDicewareWordIndex()
		if err != nil {
			return err
		}
		word, err := c.findDicewareWord(index, c.Lang)
		if err != nil {
			return err
		}
		if numberedIndex.Int64() == int64(i) && !c.RemoveNumber {
			randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(10)))
			if err != nil {
				return err
			}
			words = fmt.Sprintf("%s%s%d%s", words, word, randomNumber, separator)
			continue
		}
		words = words + word + separator
	}
	words = words[:len(words)-len(separator)]

	if c.Pbcopy || c.Hide {
		cmd := fmt.Sprintf("echo %s | pbcopy", words)
		if err := exec.Command("sh", "-c", cmd).Run(); err != nil {
			return fmt.Errorf("error copying passphrase: %s", err.Error())
		}
	}

	if c.Hide {
		return nil
	}

	if words == "" {
		return fmt.Errorf("unable to generate passphrase")
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

func (c *Config) findDicewareWord(number string, lang string) (string, error) {
	wordPath := filepath.Join("words", "diceware_words_"+lang, number+".txt")
	word := ""
	wordBytes, err := words.ReadFile(wordPath)
	if err != nil {
		word, err = findCustomDicewareWord(filepath.Join("diceware_words_"+lang, number+".txt"))
		if err != nil {
			return "", fmt.Errorf("unable to find word for index %q. err: %s", number, err.Error())
		}
	} else {
		word = string(wordBytes)
	}

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	transformedWord, _, err := transform.String(t, word)
	if err != nil {
		return "", fmt.Errorf("unable to remove special characters from %q. err: %s", word, err.Error())
	}

	if c.Lower {
		return transformedWord, nil
	}

	return cases.Title(language.Und, cases.NoLower).String(transformedWord), nil
}

func findCustomDicewareWord(wordPath string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, ".diceware-cli.d", wordPath)
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", fmt.Errorf("couldn't read word from custom dictionary")
}
