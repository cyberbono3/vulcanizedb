#!/usr/bin/env bash
# Travis doesn't support postgres 10
# https://github.com/travis-ci/travis-ci/issues/8537

set -ex

echo "Installing Postgres 10"
sudo service postgresql stop
sudo apt-get remove -q 'postgresql-*'
sudo apt-get update -q
sudo apt-get install -q postgresql-10 postgresql-client-10
sudo cp /etc/postgresql/{9.6,10}/main/pg_hba.conf

echo "Restarting Postgres 10"
sudo service postgresql restart

sudo psql -c 'CREATE ROLE travis SUPERUSER LOGIN CREATEDB;' -U postgres