# syntax=docker/dockerfile:1

# a golang + templ dockerfile, i had to read these a couple times
# docker docs: https://docs.docker.com/language/golang/build-images/
# templ docs: https://templ.guide/hosting-and-deployment/hosting-using-docker

# TODO: fix the build step so that the dockerfiles run:
#  * `make css`
#  * `make templ`
# as this will let us add `*_templ.go` and `public/css/style.css` to the
# gitignore :)

# build step
FROM golang:latest AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# deploy step
FROM gcr.io/distroless/static-debian11 AS release-stage

WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/public /public

EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
CMD ["/entrypoint"]
