-- +goose Up
-- +goose StatementBegin
create table if not exists `user_base` (
  `user_id` bigint not null primary key auto_increment,
  `user_account` varchar(255) not null,
  `user_password` varchar(255) not null,
  `user_salt` varchar(255) not null,

  `user_login_time` datetime not null default now(),
  `user_logout_time` datetime not null default now(),
  `user_login_ip` varchar(45) null,
  
  `user_created_at` datetime not null default now(),
  `user_updated_at` datetime not null default now(),

  unique key unique_email (`user_account`)
) engine=InnoDB default charset=utf8mb4
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `user_base`
-- +goose StatementEnd
