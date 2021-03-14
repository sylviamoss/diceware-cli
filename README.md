## Diceware CLI
A tool that generates strong passwords based on easily memorable words that are also extremely resistant to attack using the diceware passphrase.

[Click here](http://world.std.com/~reinhold/diceware.html) to know more about the diceware passphrase.

## Usage

### Generate passphrase
```
Usage:
  diceware-cli generate [flags]

Flags:
      --copy               pbcopy password
  -h, --help               help for generate
      --hide               pbcopy and hide password. You WON'T see the password
      --lang string        password language
                            available langs: en, pt (default "en")
      --lower              remove capitalized first letters
      --remove-number      removes the random number we add by default
      --separator string   character that separates the words.
                           use --separator=none to remove reparator (default "/")
      --size int32         the amount words the password will have (default 6)
```

Examples: 

Generate a Portuguese passphrase and copy it automatically: 
```
~> diceware-cli generate --lang=pt --copy
Pungir/Bip4/Quorum/Vau/Vida/Censor
```

Generate an English (default) passphrase with seven words, backslash separator, and copy it automatically: 
```
~> diceware-cli generate --copy --size=7 --separator=\\
Unashamed2\Sublime\Rejoin\Justly\Audition\Glove\Cahoots
```
   

### Add custom language dictionary

```
Usage:
  diceware-cli config [flags]

Flags:
      --add             add new config
  -h, --help            help for config
      --lang            add new language
      --name string     language name
      --source string   dictionary source file
```

Example of adding an Spanish diceware dictionary:
```
diceware-cli config --add-lang --source=/Users/diceware-cli/dictionary_file.txt --name=es
``` 

The `dictionary_file.txt` content must be in the same format as this [example of word list](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt). 
The custom dictionary configuration will be ketp under `$HOME/.diceware-cli.d/diceware_words_{language}`

To further generate Spanish diceware passphrase, you'd do: 
```
~> diceware-cli generate --lang=es
``` 

## Installation Guide

Unzip the zip files to find the binaries inside. To access it from the command-line, you will need to add the folder path somewhere on your **PATH** variable. 

See [this page](https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux-unix) for instructions on setting the **PATH** on Linux and Mac. [This page](https://stackoverflow.com/questions/1618280/where-can-i-set-path-to-make-exe-on-windows) contains instructions for setting the **PATH** on Windows.

The binaries are not notarized to MacOS. Users will have to manually allow the system to execute the binaries.

If you face some difficulty during the installation, please let me know by reporting an issue in this repository.

