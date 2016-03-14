#!/bin/sh

psql postgres -c "CREATE DATABASE horodata WITH ENCODING 'UTF8' TEMPLATE template0"

psql horodata -f schema.sql
psql horodata -f users.sql
psql horodata -f billing.sql
psql horodata -f jobs.sql
