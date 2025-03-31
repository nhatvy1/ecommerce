-- name: CheckUserBaseExists :one
select count(*)
from user_base
where user_account = ?;