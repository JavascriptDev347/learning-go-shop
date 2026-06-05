create type user_role as enum ('admin', 'customer');

create table users
(
    id         serial primary key,
    email      varchar(255) not null unique,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    password   varchar(255) not null,
    role       user_role    not null    default 'customer',
    phone      varchar(20),
    is_active  boolean                  default true,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp,
    deleted_at timestamp with time zone
);

create index idx_users_email on users (email);
create index idx_users_deleted_at on users (deleted_at);