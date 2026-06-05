create table refresh_tokens
(
    id         serial primary key,
    user_id    integer      not null references users (id) on delete cascade,
    token      varchar(255) not null unique,
    expires_at timestamp with time zone,
    created_at timestamp with time zone default current_timestamp,
    deleted_at timestamp with time zone
);

create index idx_refresh_tokens_token on refresh_tokens (token);
create index idx_refresh_tokens_user_id on refresh_tokens (user_id);
create index idx_refresh_tokens_deleted_at on refresh_tokens (token, deleted_at);