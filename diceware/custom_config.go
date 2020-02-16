package diceware

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CustomConfig struct {
	Lang   bool
	Add    bool
	Source string
	Name   string
}

func (c *CustomConfig) Configure() error {
	if c.Add && c.Lang {
		return c.newLanguage()
	}

	return nil
}

func (c *CustomConfig) newLanguage() error {
	if c.Source == "" || c.Name == "" {
		return fmt.Errorf("Please provide both dictionary source file and language name (--source, --name)")
	}

	file, err := os.Open(c.Source)
	if err != nil {
		return err
	}
	defer file.Close()

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dicewarePath := home + "/.diceware/diceware_words_" + c.Name
	if _, err := os.Stat(dicewarePath); os.IsNotExist(err) {
		os.MkdirAll(dicewarePath, os.ModePerm)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		f, err := os.Create(dicewarePath + "/" + words[0] + ".txt")
		if err != nil {
			return err
		}
		f.WriteString(words[1])
		f.Sync()
		f.Close()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
