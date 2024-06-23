DO
\$\$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_user WHERE usename = 'postgres') THEN
        CREATE ROLE postgres LOGIN PASSWORD 'n27qiGDJRaJ95bwbu';
    END IF;
END
\$\$;

SELECT 'CREATE DATABASE drone_calc'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'drone_calc')\gexec

GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;
