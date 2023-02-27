FROM golang:1.19-alpine as builder

WORKDIR /workspace

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . /workspace/
RUN CGO_ENABLED=0 GOOS=linux go build -o dashboard

FROM alpine:3.17

RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone && \
    apk del tzdata
COPY --from=builder /workspace/dashboard /app/dashboard
COPY /static-content /app/static-content
RUN chmod +x /dashboard
ENTRYPOINT [ "/dashboard" ]