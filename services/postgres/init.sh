#!/bin/sh

psql postgres -c "CREATE DATABASE horodata WITH ENCODING 'UTF8' TEMPLATE template0;"

# CREATE USER horodata WITH ENCRYPTED PASSWORD '';
# \c horodata
# creer les tables...
# GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public to horodata;
# GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public to horodata;

psql horodata -f schema.sql
psql horodata -f users.sql
psql horodata -f billing.sql
psql horodata -f jobs.sql
