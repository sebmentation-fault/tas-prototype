# The golang-htmx prototype

## stack

1. frontend
    1. htmx - interactivity
    1. tailwindcss - styling
    1. daisyui - styling
1. backend
    1. golang - language
    1. go fiber web framework - client-server stuff
    1. templ - writing html templates
1. todo: database (supabase)
1. todo: authentication (supabase)
1. todo: payment

## requirements

1. [docker desktop](https://docs.docker.com/desktop/install/mac-install/)
   (to run the project in a container)
1. [go](https://go.dev/doc/install) (to install templ below)
1. [templ](https://templ.guide/quick-start/installation) (to compile templ into
   go)
1. [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) to
   install tailwindcss and daisyui, and also execute them
1. [supabase](https://supabase.com/docs/guides/cli/getting-started?queryGroups=platform&platform=macos)
   to start/stop the supabase services

## directory structure

* main entry point is [main.go](./main.go) (could change to `./cmd/v1/main.go`,
  `./cmd/v2/main.go` if we have multiple versions, etc)
* [pkg/](./pkg/) has server-related code, such as [handlers](./pkg/handlers/),
  and [db](./pkg/handlers/). If we have actual apps in the future, then these
  handlers should be aware of the client content-type and send JSON instead of
  HTML, for example
* publically availiable assets are in [public/](./public/)
* HTML/view-related code is in [views/](./views/)

## run the server with live-reload

### go server

In this dir:

```sh
make live
```

Then access the server at `http://localhost:7331`.

### database

In the `supabase/` dir:

```sh
supabase start
```

Then access the supabase interface at `http://localhost:54323`.

Stop it with:

```sh
supabase stop
```

## build/run the server (in docker)

```sh
make docker-build
make docker-run
```

## images

All the following images are royalty free. License states as long as we do not
sell the content as ours then is ok to be used on the platform.

| image                                                      | source                                                                             |
| ---------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| [default/selfie](./public/images/MenSilhouettesCamera.jpg) | [Pixabay](https://pixabay.com/photos/men-silhouettes-camera-photographer-1777352/) |
| [cafe](./public/images/CoffeeBeansSeed.jpg)                | [Pixabay](https://pixabay.com/photos/coffee-beans-seed-powder-wooden-2560260/)     |
| [pub](./public/images/IrishPubLocal.jpg)                   | [Pixabay](https://pixabay.com/photos/bar-local-ireland-irish-pub-pub-209148/)      |
| [park](./public/images/ParkBenchForest.jpg)                | [Pixabay](https://pixabay.com/photos/park-bench-park-forest-meadow-6607626/)       |
| [hike](./public/images/MountainHikeFall.jpg)               | [Pixabay](https://pixabay.com/photos/mountains-hike-fall-rosswald-8411954/)        |
