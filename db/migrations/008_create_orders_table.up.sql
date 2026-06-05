create type order_status as enum ('pending', 'confirmed', 'shipped', 'delivered', 'cancelled');

create table orders
(
    id           serial primary key,
    user_id      integer not null references users (id) on delete cascade,
    total_amount decimal(10, 2) not null,
    status       order_status             default 'pending',
    created_at   timestamp with time zone default current_timestamp,
    updated_at   timestamp with time zone default current_timestamp,
    deleted_at   timestamp with time zone
);

create index idx_orders_user_id on orders (user_id);
create index idx_orders_status on orders (status);
create index idx_orders_deleted_at on orders (deleted_at);