
up:
	docker-compose up -d

down:
	docker-compose down

ifeq (build,$(firstword $(MAKECMDGOALS)))
ARGS = $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)
endif
.PHONY: build
build:
	@sudo docker-compose build $(ARGS)

ifeq (ps,$(firstword $(MAKECMDGOALS)))
ARGS = $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)
endif
.PHONY: ps
ps:
	@sudo docker-compose ps $(ARGS)

ifeq (exec,$(firstword $(MAKECMDGOALS)))
ARGS = $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)
endif
.PHONY: exec
exec:
	@sudo docker-compose exec $(ARGS)


ifeq (logs,$(firstword $(MAKECMDGOALS)))
ARGS = $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)
endif
.PHONY: logs
logs:
	@sudo docker-compose logs $(ARGS)


ifeq (recreate,$(firstword $(MAKECMDGOALS)))
ARGS = $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)
endif
.PHONY: recreate
recreate:
	@sudo docker-compose up --force-recreate --no-deps -d $(ARGS)
