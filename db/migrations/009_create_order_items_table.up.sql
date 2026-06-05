create table order_items
(
    id         serial primary key,
    order_id   integer not null references orders (id) on delete cascade,
    product_id integer not null references products (id) on delete cascade,
    quantity   integer not null check ( quantity > 0 ),
    price      decimal(10, 2) not null,
    created_at timestamp with time zone default current_timestamp,
    deleted_at timestamp with time zone
);

create index idx_order_items_order_id on order_items (order_id);
create index idx_order_items_product_id on order_items (product_id);
create index idx_order_items_deleted_at on order_items (deleted_at);