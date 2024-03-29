# syntax = docker/dockerfile:1.3


##
## Build API
##

FROM golang:1.16-buster AS build-api

WORKDIR /app

COPY api/go.mod .
COPY api/go.sum .
RUN go mod download

COPY api .

RUN --mount=type=cache,target=/root/.cache/go-build go build -o out/codenv-api


##
## Build Frontend
##

FROM node:14-buster as build-frontend

WORKDIR /app

COPY frontend/package.json .
COPY frontend/yarn.lock .

RUN  --mount=type=cache,target=/root/.cache yarn install

COPY frontend .

RUN  --mount=type=cache,target=/root/.cache yarn build


##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build-api /app/out/codenv-api /app/codenv-api
COPY --from=build-frontend /app/build /app/static

EXPOSE 8080

CMD ["/app/codenv-api"]
