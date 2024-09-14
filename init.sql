create table users (
    tg_id          bigint  unique primary key,
    tg_username    text    not null,
    sol_public_key text    unique not null,
    attempt        int     default 1,
    created_at  timestamp without time zone default NOW()
);