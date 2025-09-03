#!/bin/sh

set -e

echo "run db migration"
source /app/.env

sleep 60s

echo @$DB_SOURCE
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start app"
exec "$@"