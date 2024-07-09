#!/bin/sh
chmod +x "$0"

set -e

usage() {
    echo "使用方法: $0 [コマンド] [マイグレーションパス]"
    echo "コマンド: migrateで実行するコマンド (例: up, down, version)"
    echo "マイグレーションパス: マイグレーションファイルのパス (デフォルト: /migrations)"
    exit 1
}

if [ "$#" -gt 2 ]; then
    usage
fi

CMD=${1:-version}
MIGRATE_PATH=${2:-/DDL}

migrate \
    -path $MIGRATE_PATH \
    -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=disable" \
    -verbose $CMD
