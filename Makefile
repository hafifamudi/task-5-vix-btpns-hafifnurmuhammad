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

.PHONY : postgresup postgresdown

postgresup:
	docker run --name go-api-postgres \
	-e POSTGRES_PASSWORD=$(DB_PWD) \
	-v $(PWD):/var/lib/postgresql/data \
	-p 5433:5432 \
	-d $(DB_NAME)

postgresdown:
	docker stop go-api-postgres && docker rm go-api-postgres 

psql:
	docker exec -it go-api-postgres psql $(PSQLURL)