CREATE TABLE IF NOT EXISTS users
(
    id         serial
        constraint user_pk
            primary key,
    uid        varchar(36)  not null,
    email      varchar(255) not null,
    name       varchar(36)  not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

CREATE UNIQUE INDEX user_uid_uniq
    on users (uid);