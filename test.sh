#!/bin/sh
#Quick tests
PRGNAME="./myprogram"
CGO_ENABLED=0 GOOS=linux go build -v
# Working cases
$PRGNAME -u "https://news.ycombinator.com/"
$PRGNAME -u "https://news.ycombinator.com/" -u "https://arstechnica.com/"
$PRGNAME --url "https://news.ycombinator.com/" -o json
$PRGNAME --url "https://news.ycombinator.com/" --url "https://arstechnica.com/"
$PRGNAME --url "https://news.ycombinator.com/" --output json --output stdout
$PRGNAME --url "https://news.ycombinator.com/" --url "https://arstechnica.com/"
$PRGNAME --url "https://news.ycombinator.com/" --url "https://arstechnica.com/"

# Error cases
$PRGNAME
$PRGNAME --url
$PRGNAME --url "https://news.ycombinator.com/" --url "toto"
$PRGNAME --url "https://news.ycombinator.com/" --output --output stdout
exit 0
