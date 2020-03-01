## Diceware CLI
A tool that generates strong passwords based on easily memorable words that are also extremely resistant to attack using the diceware passphrase.

[Click here](http://world.std.com/~reinhold/diceware.html) to know more about the diceware passphrase.

## Usage

### Generate passphrase
```
Usage:
  diceware generate [flags]

Flags:
  -c, --copy               pbcopy password
  -h, --help               help for generate
      --hide               pbcopy and hide password. Password WON'T be printed out
  -l, --lang string        password language
                            available langs: en, pt (default "en")
      --separator string   character that separates the words,
                            to remove reparator use --separator=none (default " ")
  -s, --size int32         the amount words the password will have (default 6)
```    
   

### Add custom language dictionary

```
Adds new language dictionary.

Usage:
  diceware config [flags]

Flags:
  -a, --add             add new config
  -h, --help            help for config
  -l, --lang            add new language
  -n, --name string     language name
  -s, --source string   dictionary source file
```

Example:
```
diceware config --add --lang --source=/Users/diceware-cli/dictionary_file.txt --name=es
```
or
```
diceware config -a -l -s=/Users/diceware-cli/dictionary_file.txt -n=es
```   


## Installation Guide

Unzip the zip files to find the binaries inside the `pkg` folder.

On Unix systems, place the binaries in your favorite folder, but to access it from the command-line, you will need to add the folder path somewhere on your **PATH** variable. 

See [this page](https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux-unix) for instructions on setting the **PATH** on Linux and Mac. [This page](https://stackoverflow.com/questions/1618280/where-can-i-set-path-to-make-exe-on-windows) contains instructions for setting the **PATH** on Windows.

If you face some difficulty during the installation, please let me know by reporting an issue in this repository.