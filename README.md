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
docker build -t myprogram/myprogram .
docker run myprogram/myprogram -u https://yahoo.fr
```

Linter and scan
```shell
hadolint Dockerfile
trivy image myprogram/myprogram > security-report.log
```

Using Minikube, with tiny conf
```shell
eval $(minikube -p minikube docker-env)
docker build -t myprogram/myprogram .

kubectl create -f myprogram-manifest.yaml --dry-run=client
kubectl create -f myprogram-manifest.yaml
kubectl logs pod/myprogram-xxxxx
```

## Part3

```shell
```

## Part4

```shell
grep -aoiE '[a-z]+.com' <<EOF | awk '{print tolower($0)}'
http://tiktok.com
https://ads.faceBoook.com.
https://sub.ads.faCebook.com
api.tiktok.com
Google.com.
aws.amazon.com
EOF
```
