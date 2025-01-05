# ---------------------------------------------------
# base stage
# ---------------------------------------------------
FROM golang:1.23.0-alpine AS base

ARG GID=1000 \
    UID=1000 \
    USERNAME=user

RUN <<EOF
apk add --no-cache git tzdata autoconf automake libtool nasm gcc g++ make build-base
cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
addgroup -g $GID usergroup
adduser -u $UID -G usergroup -D $USERNAME
echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
EOF

USER $USERNAME

ENV TZ=Asia/Tokyo

WORKDIR /app

# ---------------------------------------------------
# builder-stage
# ---------------------------------------------------
FROM base AS deploy-builder

COPY go.mod go.sum ./
RUN go mod download

COPY --chown=$USERNAME:usergroup . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -trimpath -ldflags "-s -w" -installsuffix cgo -mod=readonly \
    -v -o /app/tmp/main ./cmd/app/main.go

# ---------------------------------------------------
# deploy-stage
# ---------------------------------------------------
FROM debian:bullseye-slim AS deploy

ARG GID=1000 \
    UID=1000 \
    USERNAME=user

RUN <<EOF
apt-get update
apt-get install -y --no-install-recommends ca-certificates tzdata
rm -rf /var/lib/apt/lists/*
ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
groupadd -g $GID usergroup
useradd -u $UID -g usergroup -m $USERNAME
echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
EOF

USER $USERNAME

ENV TZ=Asia/Tokyo

COPY --from=deploy-builder /app/tmp/main /app/main

EXPOSE 8080

CMD ["/app/main"]

# ---------------------------------------------------
# dev-stage
# ---------------------------------------------------
FROM base AS dev

USER root

RUN <<EOF
apk add --no-cache postgresql-client zlib-dev libpng-dev curl
go install github.com/air-verse/air@v1.52.3
go install github.com/go-task/task/v3/cmd/task@latest
EOF

ENV PATH="${PATH}:/root/go/bin"

COPY go.mod go.sum ./

RUN go mod download && go mod tidy

COPY --chown=$USERNAME:usergroup . .
