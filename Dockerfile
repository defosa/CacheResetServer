FROM golang:alpine as builder


COPY . /app


WORKDIR /app
RUN go build -o app


FROM nginx:alpine


COPY --from=builder /app/app /usr/share/nginx/html/app


COPY nginx.conf /etc/nginx/conf.d/default.conf


EXPOSE 80 443 8999


CMD ["nginx", "-g", "daemon off;"]
ENTRYPOINT ["/app/app"]
