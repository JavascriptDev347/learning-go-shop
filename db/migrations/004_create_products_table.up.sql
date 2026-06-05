create table products
(
    id          serial primary key,
    category_id integer             not null references categories (id) on delete cascade,
    name        varchar(255)        not null,
    description text,
    price       decimal(10, 2) not null,
    stock       integer                  default 0,
    sku         varchar(100) unique not null,
    is_active   boolean                  default true,
    created_at  timestamp with time zone default current_timestamp,
    updated_at  timestamp with time zone default current_timestamp,
    deleted_at  timestamp with time zone
);

create index idx_products_category_id on products (category_id);
create index idx_products_sku on products (sku);
create index idx_products_is_active on products (is_active);
create index idx_products_deleted_at on products (deleted_at);
