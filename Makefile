
.PHONY: services-up
services-up:
	@docker-compose up -d

.PHONY: services-down
services-down:
	@docker-compose down
