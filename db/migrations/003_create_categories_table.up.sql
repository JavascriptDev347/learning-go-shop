create table categories
(
    id          serial primary key,
    name        varchar(255) not null,
    description text,
    is_active   boolean                  default true,
    created_at  timestamp with time zone default current_timestamp,
    updated_at  timestamp with time zone default current_timestamp,
    deleted_at  timestamp with time zone
);

create index idx_categries_is_active on categories (is_active);
create index idx_categries_deleted_at on categories (deleted_at);