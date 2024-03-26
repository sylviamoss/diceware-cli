## Diceware CLI
A tool that generates strong passwords based on easily memorable words that are also extremely resistant to attack using the diceware passphrase.

[Click here](http://world.std.com/~reinhold/diceware.html) to know more about the diceware passphrase.

## Usage

### Installing 
If you have Go installed, simply run: 
```shell
go install github.com/sylviamoss/diceware-cli@latest
```

If that's not your case, check the [Installation Guide](#installation-Guide). 

### Generating passphrase
```shell
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
                           use --separator=none to remove separator (default "/")
      --size int32         the amount words the password will have (default 6)
```

Examples: 

Generate a Portuguese passphrase and copy it automatically: 
```shell
~> diceware-cli generate --lang=pt --copy
Pungir/Bip4/Quorum/Vau/Vida/Censor
```

Generate an English (default) passphrase with seven words, backslash separator, and copy it automatically: 
```shell
~> diceware-cli generate --copy --size=7 --separator=\\
Unashamed2\Sublime\Rejoin\Justly\Audition\Glove\Cahoots
```

#### Overriding flags default 

To avoid typing the same flags over and over again, you can override the default values using a configuration file.    

You can pass the configuration file path as a flag to the `generate` command:
```shell
diceware-clid generate --config=/path/to/config.yaml
```

Or you can create the default configuration file `.diceware-cli.yaml` under your home directory and the CLI will automatically pick it up.

You can generate the configuration file content using the CLI: 
```shell
~> diceware-cli config generate
# diceware-cli config file yaml content
# You can customize the default values of the flags by setting them in this file.
generate:
  lang: en
  separator: /
  size: 6
  copy: false
  hide: false
  lower: false
```

You can use the generated content to write directly to `$HOME/.diceware-cli.yaml`:
```shell
diceware-cli config generate > $HOME/.diceware-cli.yaml
```

Replace the values you would like to override with your own.


### Adding custom language dictionary

```shell
Usage:
  diceware-cli dictionary [flags]

Flags:
      --add-lang        add new config language
  -h, --help            help for dictionary
      --name string     language name
      --source string   dictionary source file
```

Example of adding a Spanish diceware dictionary:
```shell
diceware-cli dictionary --add-lang --source=/Users/diceware-cli/dictionary_file.txt --name=es
``` 

The `dictionary_file.txt` content must be in the same format as this [example of word list](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt). 
The custom dictionary configuration will be ketp under `$HOME/.diceware-cli.d/diceware_words_{language}`

To further generate Spanish diceware passphrase, you'd do: 
```shell
diceware-cli generate --lang=es
``` 

## Installation Guide

Unzip the zip files to find the binaries inside. To access it from the command line, you will need to add the folder path somewhere on your **PATH** variable. 

See [this page](https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux-unix) for instructions on setting the **PATH** on Linux and Mac. [This page](https://stackoverflow.com/questions/1618280/where-can-i-set-path-to-make-exe-on-windows) contains instructions for setting the **PATH** on Windows.

The binaries are not notarized to MacOS. Users will have to manually allow the system to execute the binaries.

If you face some difficulty during the installation, please let me know by reporting an issue in this repository.

