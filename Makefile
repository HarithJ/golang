DOCKER_DEV_COMPOSE_FILE := docker/dev/docker-compose.yml
DOCKER_PROD_FILE := docker/prod/Dockerfile

start-build:
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) build
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) run app /bin/ash

start-nobuild:
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) run app /bin/ash

production:
	@ docker-compose -f $(DOCKER_DEV_COMPOSE_FILE) run app ash -c "cd src/create-pod && \
		go install"
	@ docker build -f $(DOCKER_PROD_FILE) -t harithj/create-pod .
	@ docker push harithj/create-pod

ssh-production:
	@ docker run -it gcr.io/apprenticeship-projects-1/pod-creation /bin/ash
