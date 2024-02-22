## Part1

Build and run
```shell
make
./myprogram -u https://news.ycombinator.com/ -o stdout
```

Related tests and linter
```shell
golangci-lint run
shellcheck test.sh
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
kubectl logs -l job-name=myprogram
```

## Part3

Github: https://github.com/Ant0wan/casestudy

Summary of a CI
`https://github.com/Ant0wan/casestudy/actions/runs/8006423474`


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

## Bonus

This case study ends with the deployment of a workload in Kubernetes. However, the current deployment process lacks rigor, and a more structured flow could significantly improve deployment ease and maintainability.
The current approach utilizes a "push" GitOps flow. Following the build process, the CI triggers an action that directly pushes the workload to the cluster.

While this "push" method offers simplicity and ease of coding, it can lead to discrepancies over time between the workload definition in the Git repository and the actual running workload in the cluster. This phenomenon, known as "drift," poses a significant challenge for DevOps and Platform Engineers, as it can introduce bugs and issues in the long run.

A "pull" GitOps flow offers a solution to this problem. In this model, a continuously running tool consistently monitors for discrepancies between the application definition (stored in the repository or registry) and the currently deployed resources. If any drift is detected, the tool, referred to as an operator, automatically rectifies the issue by applying the defined configurations.
