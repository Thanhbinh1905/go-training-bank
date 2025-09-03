#!/bin/sh

set -e

echo "run db migration"
source /app/.env

echo @$DB_SOURCE
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start app"
exec "$@"