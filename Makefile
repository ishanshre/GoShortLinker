createMongoDBcontainer:
	docker run -d -p 27017:27017 --name=UrlShortnerMongoDB -v mongo_data:/data/db mongo

startContainer:
	docker start UrlShortnerMongoDB
stopContainer:
	docker stop UrlShortnerMongoDB

run:
	go run cmd/api/main.go  