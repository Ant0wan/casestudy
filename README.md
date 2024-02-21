## Part1

Build and run
```shell
make
./myprogram -u https://news.ycombinator.com/ -o stdout
```

Related tests and linter
```shell
golangci-lint run
shellcheck test
checkmake Makefile
make test
```


## Part2

Build and run
```shell
docker build -t myprogram:1 .
docker run myprogram:1 -u https://yahoo.fr
```

Linter and scan
```shell
hadolint Dockerfile
trivy image myprogram:1 > security-report.log
```


## Part3

```shell
```

## Part4

```shell
grep -aoiE '[a-z]+.com' file-part4.txt | awk '{print tolower($0)}'
```
