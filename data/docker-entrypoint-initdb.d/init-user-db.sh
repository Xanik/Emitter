#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE event;
    CREATE TABLE `targets` (
    `id` text NOT NULL DEFAULT '',
    `message` text NOT NULL DEFAULT '',
    `created_on` text NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
EOSQL