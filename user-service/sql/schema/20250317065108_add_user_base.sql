-- +goose Up
-- +goose StatementBegin
create table if not exists `users` (
  `id` bigint not null primary key auto_increment,
  `first_name` nvarchar(50) not null,
  `last_name` nvarchar(50) not null,
  `email` varchar(255) not null,
  `password` varchar(255) not null,

  `created_at` datetime not null default current_timestamp,
  `updated_at` datetime not null default current_timestamp,

  unique key unique_email (email)
) engine=InnoDB default charset=utf8mb4
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `users`
-- +goose StatementEnd
