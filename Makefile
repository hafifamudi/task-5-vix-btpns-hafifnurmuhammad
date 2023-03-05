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

.PHONY : postgresup postgresdown

postgresup:
	docker run \
	--name postgres_db \
	--network $(NET_NAME) \
	-e POSTGRES_PASSWORD=$(DB_PWD) \
	-v $(PWD):/var/lib/postgresql/data \
	-p 5432:5432 \
	-d $(DB_NAME) \


postgresdown:
	docker stop go-api-postgres && docker rm go-api-postgres 

psql:
	docker exec -it go-api-postgres psql $(PSQLURL)

build:
	docker build -t go-rest-api:0.0.1 .

run: postgresup
	docker run -p 8080:8080 --network $(NET_NAME) -d -t go-rest-api:0.0.1