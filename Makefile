DBPASS?=mysql21

test:
	cd deathstar && make test
	cd destroyer && make test

build:
	cd deathstar && make build
	cd destroyer && make build

