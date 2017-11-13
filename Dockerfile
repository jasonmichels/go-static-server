FROM gcr.io/distroless/base
EXPOSE 80
COPY ./go-static-server /
COPY ./public /public
ENTRYPOINT ["/go-static-server"]