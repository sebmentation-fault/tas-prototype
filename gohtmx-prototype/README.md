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
1. authentication (supabase) (WIP - only sign up done so far)
1. todo: payment

## requirements

1. [docker desktop](https://docs.docker.com/desktop/install/mac-install/)
   (to run the project in a container)
1. [go](https://go.dev/doc/install) (to install templ below)
1. [templ](https://templ.guide/quick-start/installation) (to compile templ into
   go)
1. [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) to
   install tailwindcss and daisyui, and also execute them
1. [supabase](https://supabase.com/docs/guides/self-hosting/docker)
   to start/stop the supabase services

### using supabase

install/clone supabase from the link above in the root of this repo
(i.e. making `tas-prototype/supabase`) and not in this directory, as doing so
makes the `templ generate --watch` command attempt to watch _all_ the supabase
files.

you will need to copy the `./supabase-docker.env` file to
`../supabase/docker/.env`.

then when inside the `supabase/docker` directory, run the following to
start/stop the docker containers:

```sh
docker compose up -d
docker compose down
```

### making the tables

<!-- TODO: automate making the tables -->

Start supabase up, then make the following tables, with the following columns,
in the `public` database:

1. "celebrities":
   * celebrity_id - has only a foreign key to auth.users.id
   * created_at - timestamp
2. "events":
   * id - id
   * created_at - timestamp
   * is_deleted - bool - default: `FALSE`
   * celebrity_id - foreign key to public.celebrities.celebrity_id
   * is_reserved - bool - default: `FALSE`
   * price - text (and is the only that can be NULLable)
   * title - text
   * description - text
   * activity - the smallest int
   * city - text
   * country - text

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

Then access the web server at `http://localhost:7331`.

### supabase

In the `supabase/docker` dir:

```sh
docker compose up -d
```

Then access the supabase interface at `http://localhost:8000`.

Stop it with:

```sh
docker compose down
```

## build/run the server (in docker)

```sh
make docker-build
make docker-run
```

## images

All the following images are royalty free. License states as long as we do not
sell the content as ours then is ok to be used on the platform.

| image | source |
| ----- | ------ |
| ![default/selfie](./public/images/MenSilhouettesCamera.jpg) | [Pixabay](https://pixabay.com/photos/men-silhouettes-camera-photographer-1777352/) |
| ![cafe](./public/images/CoffeeBeansSeed.jpg)                | [Pixabay](https://pixabay.com/photos/coffee-beans-seed-powder-wooden-2560260/)     |
| ![pub](./public/images/IrishPubLocal.jpg)                   | [Pixabay](https://pixabay.com/photos/bar-local-ireland-irish-pub-pub-209148/)      |
| ![park](./public/images/ParkBenchForest.jpg)                | [Pixabay](https://pixabay.com/photos/park-bench-park-forest-meadow-6607626/)       |
| ![hike](./public/images/MountainHikeFall.jpg)               | [Pixabay](https://pixabay.com/photos/mountains-hike-fall-rosswald-8411954/)        |
