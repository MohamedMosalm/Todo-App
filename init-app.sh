#!/bin/sh

echo "Waiting for PostgreSQL to start..."
until nc -z -v -w30 $DB_HOST $DB_PORT; do
  echo "Waiting for database connection..."
  sleep 3
done

echo "PostgreSQL is up - starting application"
exec ./main
