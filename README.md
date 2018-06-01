# jogchat_postgres

createdb jogchat;

CREATE TABLE jogchat.users(
   id UUID PRIMARY KEY DEFAULT,
   username TEXT,
   email TEXT,
   password TEXT,
   activate boolean,
);

CREATE TABLE companies(
   id UUID PRIMARY KEY DEFAULT,
   email TEXT,
   name TEXT,
)

CREATE TABLE edu(
   id UUID PRIMARY KEY DEFAULT,
   email TEXT,
   name TEXT,
)

References:
https://starkandwayne.com/blog/uuid-primary-keys-in-postgresql/

