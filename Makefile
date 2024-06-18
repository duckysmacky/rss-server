build:
	@go build -o server.exe cmd/rss-server/main.go
	@echo Built server

run: build
	@./server

clean:
	@rm ./server.exe