# jogchat_postgres

createdb jogchat;

CREATE TABLE jogchat.users(
   id UUID PRIMARY KEY,
   username TEXT,
   email TEXT,
   password TEXT,
   activate boolean
);

CREATE TABLE companies(
   id UUID PRIMARY KEY,
   email TEXT,
   name TEXT
)

CREATE TABLE edus(
   id UUID PRIMARY KEY,
   email TEXT,
   name TEXT
)

References:
https://starkandwayne.com/blog/uuid-primary-keys-in-postgresql/

How to install postgres:
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-16-04


Connect to database:
CREATE DATABASE jogchat;
\dt
\c jogchat;



How to set new password for postgres:
http://suite.opengeo.org/docs/latest/dataadmin/pgGettingStarted/firstconnect.html
Run the psql command from the postgres user account:

sudo -u postgres psql postgres
Set the password:

\password postgres
Enter a password.

Close psql.

\q
