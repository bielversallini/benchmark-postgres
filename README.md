# benchmark-postgres
Repo to benchmark postgres performance


## Run locally


## Build docker image and run in OCP cluster
Build the image
```
docker build . -t benchmark-postgres
```

Push image to quay
```
export QUAY_USER=<your-quay-user-id>
docker login -u ${QUAY_USER} quay.io
docker tag benchmark-postgres quay.io/${QUAY_USER}/benchmark-postgres:latest 
docker push quay.io/${QUAY_USER}/benchmark-postgres:latest 
```

Run Job in OCP cluster
```
oc login ...
oc apply -f deploy/benchmark-job.yaml
```
