# jogchat_postgres

createdb jogchat;

CREATE TABLE jogchat.users(
   id UUID PRIMARY KEY DEFAULT,
   nickname TEXT,
   email TEXT,
   password TEXT,
   activate boolean,
);


References:
https://starkandwayne.com/blog/uuid-primary-keys-in-postgresql/

