CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL,
    username VARCHAR NOT NULL,
    time_created TIMESTAMP NOT NULL,
    CONSTRAINT "pk_users_id" PRIMARY KEY ("id")
);
