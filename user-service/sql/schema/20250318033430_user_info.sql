-- -- +goose Up
-- -- +goose StatementBegin
-- create table if not exists `user_info` (
--   `user_id` bigint unsigned primary key,
--   `user_first_name` nvarchar(50) not null,
--   `user_last_name` nvarchar(50) not null,
--   `user_nickname` nvarchar(255),
--   `user_avatar` varchar(255),
--   `user_state` tinyint unsigned not null comment 'User state: 0-Locked, 1-Activated, 2-Not Activated'
--   `user_gender` tinyint unsigned comment 'user gender: 0-Secret, 1-Male, 2-Female',
--   `user_birthday` date,
--   `user_phone` varchar(20)

--   `created_at` timestamp default current_timestamp,
--   `updated_at` timestamp default current_timestamp on update current_timestamp,

-- ) engine=InnoDB default charset=utf8mb4
-- -- +goose StatementEnd

-- -- +goose Down
-- -- +goose StatementBegin
-- drop table if exists `user_info`
-- -- +goose StatementEnd
