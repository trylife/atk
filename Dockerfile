FROM golang:1.15 as compiler

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o atk .

RUN mkdir publish && cp atk publish && cp app.yaml.example publish/app.yaml

# 第二阶段
FROM alpine

WORKDIR /app

COPY --from=compiler /app/publish .

# 注意修改端口
EXPOSE 8080

ENTRYPOINT ["./atk"]