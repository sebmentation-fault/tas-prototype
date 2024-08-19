run: build
	@./bin/app

build:
	@go build -o bin/app .

templ:
	@templ generate

css:
	tailwindcss -i views/css/tailwind.css -o public/css/style.css --minify

dev-css:
	tailwindcss -i views/css/tailwind.css -o public/css/style.css --watch

