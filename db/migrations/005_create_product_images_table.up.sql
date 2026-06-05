create table product_images
(
    id         serial primary key,
    product_id integer      not null references products (id) on delete cascade,
    url        varchar(500) not null,
    alt_text   varchar(255),
    is_primary boolean               default false,
    created_at timestamp    not null default current_timestamp,
    deleted_at timestamp with time zone
);

create index idx_product_images_product_id on product_images (product_id);
create index idx_product_images_is_primary on product_images (is_primary);
create index idx_product_images_deleted_at on product_images (deleted_at);