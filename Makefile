# database name
DB_NAME ?= postgres

# database type
DB_TYPE ?= postgres

# database username
DB_USER ?= postgres

# database password
DB_PWD ?= mysecretpassword

# psql URL
IP=localhost

PSQLURL ?= $(DB_TYPE)://$(DB_USER):$(DB_PWD)@$(IP):5432/$(DB_NAME)

# container network name
NET_NAME ?= go-postgres
# container name
CON_NAME ?= postgres_db
# postgres version
POST_VERSION ?= postgres:15.2-alpine

.PHONY : postgresup postgresdown test build

postgresup:
	docker run \
	--name $(CON_NAME) \
	--network $(NET_NAME) \
	-p 5433:5432 \
	-e POSTGRES_PASSWORD=$(DB_PWD) \
	-v $(PWD):/var/lib/postgresql/data \
	-t $(POST_VERSION) \
	-d $(CON_NAME)

postgresdown:
	docker stop go-api-postgres && docker rm go-api-postgres 

psql:
	docker exec -it go-api-postgres psql $(PSQLURL)

test: 
	go test ./test -v

build:
	docker build -t go-rest-api:0.0.1 .

go_app:
	docker run --name go-rest-api \
	-p 8080:8080 \
	--network $(NET_NAME) \
	-d -t go-rest-api:0.0.1

run: go_app postgresup

# USEFUL TARGET FOR DESTROY or DELETE ALL RUNNING or CACHE CONTAINER
IMAGE_NAME ?= go-rest-api:0.0.1 # insert your custome image name or ID_

delete_container:
	docker rm -f $(CON_NAME)
delete_image:
	docker rmi -f $(IMAGE_NAME)