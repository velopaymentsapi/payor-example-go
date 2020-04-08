db:
	docker-compose up -d db

env:
	# make key=123 secret=shhh payor=testers env
	cp .env.example .env
	sed -i.bak 's/VELO_API_APIKEY=contact_velo_for_info/VELO_API_APIKEY=$(key)/' .env && rm .env.bak
	sed -i.bak 's/VELO_API_APISECRET=contact_velo_for_info/VELO_API_APISECRET=$(secret)/' .env && rm .env.bak
	sed -i.bak 's/VELO_API_PAYORID=contact_velo_for_info/VELO_API_PAYORID=$(payor)/' .env && rm .env.bak

network:
	- docker network create payorexample

build:
	docker-compose build --no-cache api

up: clean network
	docker-compose run -d --service-ports api

sh: clean network
	docker-compose run --service-ports api sh

down:
	- docker-compose down

clean:
	- docker-compose rm -f

destroy:
	- docker rm /db
	- docker rm /redis
	- docker rmi -f payor-example-go_api

setdep:
	# make version=2.16.18 setdep
	sed -i.bak 's/"github.com\/velopaymentsapi\/velo-go .*/github.com\/velopaymentsapi\/velo-go v${version}/g' go.mod && rm go.mod.bak

updatedeps:
	# stub

dev:
	- docker-compose build api
	docker-compose run --service-ports api

refresh: down clean build up