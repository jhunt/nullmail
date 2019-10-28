nullmail: main.go
	go build .

docker:
	docker build -t huntprod/nullmail .

push: docker
	docker push huntprod/nullmail
