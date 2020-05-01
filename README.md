<p align="center">
  <h3 align="center">jowfuzz</h3>
  <p align="center">Jowfuzz is a web recognition tool.</p>

  <p align="center">
    <a href="https://twitter.com/dtr0x80x">
      <img src="https://img.shields.io/badge/twitter-@dtr0x80-blue.svg">
    </a>
    <a href="https://travis-ci.org/dtr0x80/jowfuzz">
      <img src="https://travis-ci.org/dtr0x80/jowfuzz.svg?branch=master">
    </a>
    <a href="https://www.gnu.org/licenses/gpl-3.0">
      <img src="https://img.shields.io/badge/License-GPLv3-blue.svg">
    </a>
  </p>
</p>
<hr>

## Build

```
$ make or go build main.go -o jowfuzz
```

## Features

```
Web crawler
Subdomain discovery by DNS
Directory discovery
```

## Examples

### Search subdomains
```
./jowfuzz -host google.com -wordlist databases/subdomains/subdomains.txt -subdomain
```

### Search directory's

```
./jowfuzz -host google.com -wordlist databases/directorys/directorys.txt -dir
```

### Web crawler (WIP)

```
./jowfuzz -host https://google.com -crawler -verbose
```

## Developer

```
[+] Gabriel Dutra A.K.A Dtr0x80
[+] dtr0x80@protonmail.com
[+] twitter.com/dtr0x80
```
