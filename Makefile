SHELL = /bin/bash
CFG   ?= .env
# ------------------------------------------------------------------------------
-include $(CFG)
export
#-------------------------------------------------------------------------------
.PHONY: all help psql test down clean-docker down-all reup up

all: ## Full list options
all: help

up: ## Start pg & app containers
up: CMD=up -d
up: dc

reup: ## Restart pg & app containers
reup: CMD=up --force-recreate -d
reup: dc

down: ## Stop containers and remove them
down: CMD=rm -f -s
down: dc

clean-docker: ## Delete docker image
		docker rmi $(IMAGE_MARKETING):$(IMAGE_MARKETING_VER)

down-all: ## Stop containers and remove them & delete docker image
down-all: down clean-docker

# Wait for postgresql container start
docker-wait:
	@echo -n "Checking PG is ready..."
	@until [[ `docker inspect -f "{{.State.Health.Status}}" $(DCAPE_DB)` == healthy ]] ; do sleep 1 ; echo -n "." ; done
	@echo "Ok"

psql: ## Check PG db
psql: docker-wait
	@docker exec -it $(DCAPE_DB) psql -U $(DB_USER) -d $(DB_NAME)

# $$PWD используется для того, чтобы текущий каталог был доступен в контейнере по тому же пути
# и относительные тома новых контейнеров могли его использовать
## run docker-compose
dc: docker-compose.yml
	@docker run --rm  \
	  -v /var/run/docker.sock:/var/run/docker.sock \
	  -v $(PWD):$(PWD) \
	  -w $(PWD) \
	  docker/compose:$(DC_VER) \
	  -p $(PROJECT_NAME_MARKETING) \
	  $(CMD)

test: ## Run tests
	$(GO) test -v tpro/stage1g/utils
	echo $(PWD)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
