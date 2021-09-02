# benchmark-postgres
Repo to benchmark postgres performance

## Prerequisite
You need an instance of PostgreSQL. Here are some alternatives to deploy postgres:
- On Openshift you can use the Crunchy operator. [These insttructions worked for me](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/)
- As an [AWS RDS managed service](https://console.aws.amazon.com/rds/home?region=us-east-1#)

You will need the database URL for your Postgres instance.
```
postgresql://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_NAME}
```

## Run locally
Performance is significantly slower.
1. Open a port-forward to your postgres instance as described [here.](https://access.crunchydata.com/documentation/postgres-operator/v5/quickstart/)
2. Set the environment variables DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, and DB_NAME

## Build docker image and run in OCP cluster
Build the image
```
docker build . -t benchmark-postgres
```

Push image to quay.
**Get push access to quay.io/jlpadilla/benchmark-postgres from Jorge or Sherin.**
```
docker login -u {YOUR-QUAY_USER} quay.io
docker tag benchmark-postgres quay.io/${QUAY_USER}/benchmark-postgres:latest 
docker push quay.io/${QUAY_USER}/benchmark-postgres:latest 
```

Run Job in OCP cluster
```
oc login ...
oc apply -f deploy/benchmark-job.yaml
```

Updated: Sep 2
