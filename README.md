# lottoman
Generating random numbers

##Conditions
* Hoosier Lottery is 6 numbers between 1 and 48 Same number cannot be drawn twice
* Powerball is 5 numbers between 1 and 59 and The 'Powerball' between 1 and 35

##Download the software

[Download](https://github.com/barneshere/lottoman/releases) the latest version of lottoman for all major platforms.

##Compile and run the source

Requires Go 1.4 or newer (earlier versions untested).

```
git clone https://github.com/barneshere/lottoman
cd lottoman
go get https://github.com/mitchellh/gox
go run 6numbers.go
```

##Usage

```NAME:
   lottoman - Application to get the lucky numbers for the lotto.

USAGE:
   lottoman [global options] command [command options] [arguments...]

VERSION:
   0.3

AUTHOR:
  Chubbs Solutions - <lottoman@chubbs.solutions>

COMMANDS:
   hoosierlotto, hl	Get numbers for the Hoosier Lotto
   powerball, pb Get numbers for the powerball
   help, h    	 Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

##About

Crafted with :heart: in Indiana by [Chubbs Solutions] (http://chubbs.solutions).
