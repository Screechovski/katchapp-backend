#!/usr/bin/env bash
set -euo pipefail

# Load .env if present
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs -d '\n' -I {} sh -c 'echo {}' | xargs)
fi

echo "🛑 Останавливаем процесс на порту 8080, если есть..."
lsof -ti:8080 | xargs kill -9 2>/dev/null || echo "Нет процессов на порту 8080"

echo "Starting Postgres via docker compose..."
docker compose up -d db

echo "Waiting for database to be ready..."
until docker compose exec -T db pg_isready -U "${DB_USER:-postgres}" -d "${DB_NAME:-postgres}" 2>/dev/null; do
  sleep 1
done

echo "Database is ready. Starting app with go run main.go"
nohup go run main.go > app.log 2>&1 &
echo $! > app.pid
echo "✅ App started with PID $(cat app.pid)"