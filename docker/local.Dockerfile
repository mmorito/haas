FROM golang:1.19-alpine AS build

RUN apk --no-cache add ca-certificates curl

# RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
#   wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
#   chmod +x /bin/grpc_health_probe

# cosmtrek/airをダウンロード
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

FROM golang:1.19-alpine

# ビルドステージでダウンロードしたバイナリをコピー
WORKDIR /usr/bin
COPY --from=build /go/bin/air /go/bin/air
# COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe

# mnesユーザーの作成
ARG USER=mnes
ARG GROUPNAME=mnes
ARG GODIR=/go

RUN addgroup -S ${GROUPNAME} && adduser -S ${USER} -G ${GROUPNAME}

# mnesユーザーにgoディレクトリ配下の操作権限を付与
RUN chown -R ${USER}:${GROUPNAME} ${GODIR}

USER ${USER}

# airのバイナリを実行
WORKDIR /go/bin
CMD ["air"]
