<p align="center">
  <h3 align="center">rapt0r</h3>
  <p align="center">Rapt0r is a web recognition tool.</p>

  <p align="center">
    <a href="https://twitter.com/dtr0x80">
      <img src="https://img.shields.io/badge/twitter-@dtr0x80-blue.svg">
    </a>
    <a href="https://travis-ci.org/dtr0x80/rapt0r">
      <img src="https://travis-ci.org/dtr0x80/rapt0r.svg?branch=master">
    </a>
    <a href="https://www.gnu.org/licenses/gpl-3.0">
      <img src="https://img.shields.io/badge/License-GPLv3-blue.svg">
    </a>
  </p>
</p>
<hr>

<p>
Rapt0r is a web application fuzzer. It is possible to analyze subdomains, directories and web crawler on targets.
</p>

[![asciicast](https://asciinema.org/a/14.png)](https://asciinema.org/a/326342)

## Build

```
$ make or go build main.go -o rapt0r
```

<hr>

## Examples

### Searching subdomains
```
./rapt0r -host google.com -wordlist databases/subdomains/subdomains.txt -subdomain
```

### Searching directories

```
./rapt0r -host google.com -wordlist databases/directorys/directorys.txt -dir
```

### Web crawler (WIP)

```
./rapt0r -host https://google.com -crawler -verbose
```

<hr>

## Docker environment

```
docker build -t rapt0r .
```

## Example

```
docker run -it rapt0r -host clubedostrinta.com.br -wordlist databases/subdomains/subdomains.txt -subdomain
```

## Developer

```
[+] Gabriel Dutra A.K.A Dtr0x80
[+] dtr0x80@protonmail.com
[+] twitter.com/dtr0x80
```
