FROM gcr.io/distroless/base
EXPOSE 80
COPY ./ /
ENTRYPOINT ["/go-static-server"]