<p align="center">
  <h3 align="center">jowfuzz</h3>
  <p align="center">Jowfuzz is a web recognition tool.</p>

  <p align="center">
    <a href="https://twitter.com/dtr0x80">
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

<p>
Jowfuzz is a web application fuzzer. It is possible to analyze subdomains, directories and web crawler on targets.
</p>

![gif](https://i.imgur.com/kacluJH.gif)

## Build

```
$ make or go build main.go -o jowfuzz
```

<hr>

## Examples

### Searching subdomains
```
./jowfuzz -host google.com -wordlist databases/subdomains/subdomains.txt -subdomain
```

### Searching directories

```
./jowfuzz -host google.com -wordlist databases/directorys/directorys.txt -dir
```

### Web crawler (WIP)

```
./jowfuzz -host https://google.com -crawler -verbose
```

<hr>

## Docker environment

```
docker build -t jowfuzz .
```

## Example

```
docker run -it jowfuzz -host clubedostrinta.com.br -wordlist databases/subdomains/subdomains.txt -subdomain
```

## Developer

```
[+] Gabriel Dutra A.K.A Dtr0x80
[+] dtr0x80@protonmail.com
[+] twitter.com/dtr0x80
```
