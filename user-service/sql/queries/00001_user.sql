-- name: FindByEmail :one
select *
from `users`
where email = ?;

-- name: FindById :one
select *
from `users`
where id = ?;

-- name: Create :execresult
insert into `users` (
  first_name, last_name, email, `password`,
  created_at, updated_at
) values (
  ?, ?, ?, ?, NOW(), NOW()
);

-- name: UpdateById :exec
update `users` 
set
  first_name = coalesce(sqlc.narg('first_name'), first_name),
  last_name = coalesce(sqlc.narg('last_name'), last_name),
  updated_at = NOW()
where id = sqlc.narg('id');