## Diceware CLI
A tool that generates strong passwords based on easily memorable words that are also extremely resistant to attack using the diceware passphrase.

[Click here](http://world.std.com/~reinhold/diceware.html) to know more about the diceware passphrase.

## Usage
```
diceware generate [flags]

Flags:
  -c, --copy          pbcopy password
  -h, --help          help for generate
      --hide          pbcopy and hide password. Password WON'T be printed out
  -l, --lang string   password language (default "en")
  -s, --size int32    the amount words the password will have (default 6)
```