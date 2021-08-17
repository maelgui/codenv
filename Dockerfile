##
## Build API
##

FROM golang:1.16-buster AS build-api

WORKDIR /app

COPY api/go.mod .
COPY api/go.sum .
RUN go mod download

COPY api .

RUN go build -o out/codenv-api

##
## Build Frontend
##

FROM node:14-buster as build-frontend

WORKDIR /app

COPY frontend/package.json .
COPY frontend/yarn.lock .

RUN yarn install

COPY frontend .

RUN yarn build

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build-api /app/out/codenv-api /app/codenv-api
COPY --from=build-frontend /app/build /app/static

EXPOSE 8080

CMD ["/app/codenv-api"]
