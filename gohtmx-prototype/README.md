# The golang-htmx prototype

## stack

1. frontend
    1. golang
    1. go fiber web framework 1. templ for writing html templates
1. frontend
    1. htmx
    1. tailwindcss
    1. todo: daisyui
1. todo: database (supabase)
1. todo: authentication (supabase)
1. todo: payment

## requirements

1. [docker desktop](https://docs.docker.com/desktop/install/mac-install/)
   (to run the project in a container)
1. [go](https://go.dev/doc/install) (to install templ below)
1. [templ](https://templ.guide/quick-start/installation) (to compile templ into go)
1. [tailwindcss (standalone)](https://tailwindcss.com/blog/standalone-cli) or
   [tailwindcss (npm)](https://tailwindcss.com/docs/installation)
   (to compile the css)

## directory structure

* main entry point is [main.go](./main.go) (could change to `./cmd/v1/main.go`,
  `./cmd/v2/main.go` if we have multiple versions, etc)
* [pkg/](./pkg/) has server-related code, such as [handlers](./pkg/handlers/),
  and [db](./pkg/handlers/). If we have actual apps in the future, then these
  handlers should be aware of the client content-type and send JSON instead of
  HTML, for example
* publically availiable assets are in [public/](./public/)
* HTML/view-related code is in [views/](./views/)

## build/run the server (in docker)

```sh
make docker-build
make docker-run
```
