# jogchat_postgres

# Install postgres on Ubuntu 16.04:
* [How to install postgres](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-16-04)

# Schema:

```postgres
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
```
* [UUID Postgres](https://starkandwayne.com/blog/uuid-primary-keys-in-postgresql/) 

# Server JogchatPostgres:
ip: 206.189.212.172:5432

# Connect to database:
```
sudo -u postgres psql postgres

CREATE DATABASE jogchat;
\dt
\c jogchat;
```

## How to set new password for postgres:
[postgres_start_guide](http://suite.opengeo.org/docs/latest/dataadmin/pgGettingStarted/firstconnect.html)
```
Run the psql command from the postgres user account:
Set the password:
\password postgres
Enter a password.
Close psql.
\q
```

# Postgres
*[godoc pq](https://godoc.org/github.com/lib/pq)


