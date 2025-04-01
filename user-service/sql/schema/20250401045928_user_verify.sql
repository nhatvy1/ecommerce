-- +goose Up
-- +goose StatementBegin
create table if not exists user_verify (
  verify_id int auto_increment primary key,
  verify_otp varchar(6) not null,
  verify_key VARCHAR(255) NOT NULL,                     -- verify_key: User's email (or phone number) to identify the OTP recipient
  verify_key_hash varchar(255) not null,
  verify_type int default 1,  -- 1: Email
  is_verified int default 0,  -- 0: No, 1: Yes - OTP verification status
  is_deleted int default 0,   -- 0: No, 1: Yes - Deletion status

  verify_created_at timestamp default current_timestamp,
  verify_updated_at timestamp default current_timestamp on update current_timestamp,

  user_base_id int not null,
  foreign key (user_base_id) references user_base(user_id),

  index idex_verify_otp (verify_otp)
) engine=InnoDB default charset=utf8mb4
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `user_verify`;
-- +goose StatementEnd