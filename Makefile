build: buildserver buildstatic

run: 
	docker-compose stop; docker-compose up -d;

buildserver:
	cd writ && make dockerPackage;

buildstatic:
	cd writstatic && make;

reload: buildstatic run