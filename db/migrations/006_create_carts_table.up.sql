create table carts
(
    id         serial primary key,
    user_id    integer   not null references users (id) on delete cascade,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default current_timestamp,
    deleted_at timestamp with time zone
);

create index idx_carts_user_id on carts (user_id);
create index id_carts_deleted_at on carts (deleted_at);