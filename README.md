# jogchat_postgres

# How to allow remote connections to PostgreSQL database server
[Remote connections to PostgreSQL database server] (https://bosnadev.com/2015/12/15/allow-remote-connections-postgresql-database-server/)
* grep listen /etc/postgresql/9.5/main/postgresql.conf  
* listen_addresses = 'localhost'		# what IP address(es) to listen on;
* vim /etc/postgresql/9.5/main/pg_hba.conf
* add new line:
```
host all all 0.0.0.0/0 md5
```
* restart
```
/etc/init.d/postgresql restart
```

* how to connect to remote postgres using commands
```
psql -h 206.189.212.172 -U postgres -d jogchat
```

# Install postgres on Ubuntu 16.04:
* [How to install postgres](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-16-04)

# Schema:

```postgres
createdb jogchat;

CREATE TABLE users(
   id BINARY(16) PRIMARY KEY,
   username TEXT,
   email TEXT,
   phone INT(10),
   password TEXT,
   activate boolean
);

CREATE TABLE companies(
   id BINARY(16) PRIMARY KEY,
   category TEXT,
   domain TEXT,
   name TEXT
);

CREATE TABLE schools(
   id BINARY(16) PRIMARY KEY,
   category TEXT,
   domain TEXT,
   name TEXT
);
```


```
Cell
DROP TABLE IF EXISTS cell;

SHOW WARNINGS;

CREATE TABLE cell
(
    added_at         INTEGER PRIMARY KEY AUTO_INCREMENT,
    row_key          VARCHAR(36) NOT NULL,
    column_name      VARCHAR(64) NOT NULL,
    ref_key          INTEGER NOT NULL,
    body             JSON,
    created_at       DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE `cell_idx`(`row_key`, `column_name`, `ref_key`)
) ENGINE=InnoDB;

SHOW WARNINGS;
```
* [UUID Postgres](https://starkandwayne.com/blog/uuid-primary-keys-in-postgresql/) 

# How to drop tables:
```
DROP TABLE table0, table1 ...;
```

# Server JogchatPostgres:
ip: 206.189.212.172:5432

# Connect to database:
```
sudo -u postgres psql postgres

CREATE DATABASE jogchat;
\c jogchat;
\dt
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
# References:
* [UUID Postgres](https://starkandwayne.com/blog/uuid-primary-keys-in-postgresql/) 
* [How to install postgres](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-16-04)
* [Go auth example](https://github.com/sohamkamani/go-password-auth-example)
