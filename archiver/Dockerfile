# syntax=docker/dockerfile:1.2

# NOTE: because we need files from outside the `archiver` directory,
#       we build this image from project root, symlinking this file to
#       Dockerfile.archiver

FROM golang:1.16.5-buster as builder

WORKDIR /horahora/archiver

# download modules
COPY archiver/go.mod /horahora/archiver/
COPY archiver/go.sum /horahora/archiver/
COPY video_service/go.mod /horahora/video_service/
COPY video_service/go.sum /horahora/video_service/

RUN go mod download

# build binary
COPY archiver /horahora/archiver
COPY video_service/protocol /horahora/video_service/protocol

RUN --mount=type=cache,target=/root/.cache/go-build go build -o /archiver.bin

ENTRYPOINT ["/archiver.bin"]
