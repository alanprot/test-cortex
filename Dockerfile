FROM amazonlinux:latest AS build
RUN yum -y update && rm -rf /var/cache/yum/*
RUN yum install -y  \
      ca-certificates \
      git \
      bash \
      go

RUN mkdir /cortex
WORKDIR /cortex
COPY go.mod .
COPY go.sum .

RUN go env -w GOPROXY=direct
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/cortex


FROM       alpine:3.13
RUN        apk add --no-cache ca-certificates
COPY --from=build /bin/cortex /bin/cortex
EXPOSE     80
ENTRYPOINT [ "/bin/cortex" ]