#! /bin/sh
%connect local
# Success: 'local' connection established and activated for user 'demo_user'
CREATE DATABASE jaffle_shop AS PERM = 1e9;
DELETE DATABASE "jaffle_shop";
