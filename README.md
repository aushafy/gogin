### Dependencies
- golang v1.18
- docker
- docker-compose

### Run on local env
*Start mysql with docker-compose*

$ docker-compose up -d --build

*Run migration sql*

$ mysql -h 127.0.0.1 -D gogin -u root -p < migration.sql

*Testing by curl app endpoints*

$ curl -X GET http://localhost:8080/endpoint
$ curl -X GET http://localhost:8080/weather

$ cd terraform

$ terraform init

