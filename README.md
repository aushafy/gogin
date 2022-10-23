# Gogin API Service

## Prerequisite
- golang v1.18 or latest
- docker
- docker-compose
- terraform
- ansible
- ansible-playbook

  

## Run on local env

**Start mysql with docker-compose**

    docker-compose up -d --build

  

**Run migration sql**

    mysql -h 127.0.0.1 -D gogin -u root -p < migration.sql

  

**Testing by curl app endpoints**

    curl -X GET http://localhost:8080/endpoint
    curl -X GET http://localhost:8080/weather

  
## Provisioning GCE VMs

    cd terraform
    terraform init
    terraform apply

## Setup Mysql and Docker in GCE VMs

    cd playbook
    ansible-playbook -i ./hosts playbook-api.yaml
    ansible-playbook -i ./hosts playbook-mysql.yaml
