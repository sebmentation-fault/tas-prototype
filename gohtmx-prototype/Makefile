# below is for building the docker container

docker-build:
	docker build -t gohtmx:latest .

docker-run:
	docker run -p 8080:8080 -it --rm --name gohtmxserver gohtmx:latest

# NOTE: source - https://templ.guide/commands-and-tools/live-reload-with-other-tools
#
# below is used for development on local machine

# run templ generation in watch mode to detect all .templ files and
# re-create _templ.txt files on change, then send reload event to browser.
# Default url: http://localhost:7331
live/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
live/tailwind:
	npx tailwindcss -i ./views/css/tailwind.css -o ./public/css/style.css --minify --watch

# run esbuild to generate the index.js bundle in watch mode.
#live/esbuild:
#	npx esbuild js/index.ts --bundle --outdir=assets/ --watch

# watch for any js or css or jpg change in the public/ folder, then reload the browser via templ proxy.
live/sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "public" \
	--build.include_ext "js,css,jpg"

# start all 5 watch processes in parallel.
live: 
	make -j4 live/templ live/server live/tailwind live/sync_assets
#	make -j5 live/templ live/server live/tailwind live/esbuild live/sync_assets
