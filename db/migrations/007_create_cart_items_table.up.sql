create table cart_items
(
    id         serial primary key,
    cart_id    integer not null references carts (id) on delete cascade,
    product_id integer not null references products (id) on delete cascade,
    quantity   integer not null check (quantity > 0),
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp,
    deleted_at timestamp with time zone,
    unique (cart_id, product_id)
);
create index idx_cart_items_cart_id on cart_items (cart_id);
create index idx_cart_items_product_id on cart_items (product_id);
create index idx_cart_items_deleted_at on cart_items (deleted_at);