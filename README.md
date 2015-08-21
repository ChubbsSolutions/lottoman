# lottoman

[![Build Status](https://travis-ci.org/ChubbsSolutions/lottoman.png)](https://travis-ci.org/ChubbsSolutions/lottoman)

Hoosier Lottery and Powerball number generator.

##Conditions
* Hoosier Lottery is 6 numbers between 1 and 48 Same number cannot be drawn twice.
* Powerball is 5 numbers between 1 and 59 and The 'Powerball' between 1 and 35.

##Download the software

[Download](https://github.com/ChubbsSolutions/lottoman/releases) the latest version of lottoman for all major platforms.

##Compile and run the source

Requires Go 1.4 or newer (earlier versions untested). Remember to set the GOPATH variable.

```
git clone https://github.com/ChubbsSolutions/lottoman
cd lottoman
go get https://github.com/mitchellh/gox
go run lottoman.go
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

##Email

Lottoman can now email you the lotto numbers. You just need to set a Mailgun API key and domain. Here's how you would do it in Linux:

```
export MAILGUN_PUBLIC_API_KEY=example-key
export MAILGUN_DOMAIN=chubbs.solutions
```
Just add an email address as an argument.


##About

Crafted with :heart: in Indiana by [Chubbs Solutions] (http://chubbs.solutions).
