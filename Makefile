ifneq (,$(wildcard ./.env))
    include .env
    export
endif

atlas-apply:
	atlas migrate apply \
  --url "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}&sslmode=disable"