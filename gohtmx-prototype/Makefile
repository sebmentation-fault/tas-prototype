run: build
	@./bin/app

build:
	@go build -o bin/app .

templ:
	@templ generate

css:
	tailwindcss -i views/css/tailwind.css -o public/css/style.css --minify

docker-build:
	docker build -t gohtmx:latest .

docker-run:
	docker run -p 8080:8080 -it --rm --name gohtmxserver gohtmx:latest

dev-css:
	tailwindcss -i views/css/tailwind.css -o public/css/style.css --watch

