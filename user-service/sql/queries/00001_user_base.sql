-- name: FindUserByEmail :one
select *
from `user_base`
where user_email = ?;

-- name: FindUserInfoAdmin :one
select `user_id`, `user_password`, `user_salt`, `user_email`
      `user_created_at`, `user_updated_at`
from `user_base`
where user_email = ?;

-- name: FindListUsers :many
select `user_id`, `user_password`, `user_salt`, `user_email`
      `user_created_at`, `user_updated_at`
from `user_base`
limit 0, 20;

-- name: CheckUserExists :one
select count(*)
from `user_base`
where user_email = ?;

-- name: InsertUserBase :execresult
insert into `user_base` (
  `user_email`, `user_password`, `user_salt`,
  `user_created_at`, `user_updated_at`
) values (
  ?, ?, ?, NOW(), NOW()
);

-- name: UpdateUserBase :exec
update `user_base` 
set
  user_email = coalesce(sql.narg('user_email'), user_email),
  user_password = coalesce(sqlc.narg('user_password'), user_password),
  user_salt = coalesce(sqlc.narg('user_salt'), user_salt),
  
  updated_at = NOW()
where user_id = sqlc.narg('user_id');