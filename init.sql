create table users (
    tg_id          integer unique primary key,
    tg_username    text    not null,
    sol_public_key text    not null,
    tg_chat_id     integer not null,
    created_at  timestamp without time zone default NOW()
);
