FROM golang:1.19-buster as builder

WORKDIR /app

COPY . ./

