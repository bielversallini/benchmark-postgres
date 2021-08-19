# benchmark-postgres
Repo to benchmark postgres performance

## Prerequisite
You need an instance of PostgreSQL. This can be deployed in many different ways here are some alternatives:
- On Openshift you can use the Crunchy operator. [These insttructions worked for me](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/)
- As an [AWS RDS managed service](https://console.aws.amazon.com/rds/home?region=us-east-1#)

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
