#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE event;
    CREATE TABLE `targets` (
  `id` varchar(40) NOT NULL DEFAULT '',
  `message` text NOT NULL DEFAULT '',
  `created_on` varchar(40) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
EOSQL