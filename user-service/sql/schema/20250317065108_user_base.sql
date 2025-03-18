-- +goose Up
-- +goose StatementBegin
create table if not exists `user_base` (
  `user_id` bigint not null primary key auto_increment,
  `user_email` varchar(255) not null,
  `user_password` varchar(255) not null,
  `user_salt` varchar(255) not null,

  `user_created_at` datetime not null default now(),
  `user_updated_at` datetime not null default now(),

  unique key unique_email (`user_email`)
) engine=InnoDB default charset=utf8mb4
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `user_base`
-- +goose StatementEnd
