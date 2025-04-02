-- name: InsertUserVerify :exec
insert into user_verify (
  verify_otp,
  verify_key,
  verify_key_hash,
  verify_type
) values (
  ?, ?, ?, ?
);

-- name: GetInfoOTP :one
select verify_id, verify_otp, verify_key, verify_key_hash, verify_type, is_verified, is_deleted, verify_created_at, verify_updated_at
from `user_verify`
where verify_key_hash = ?;

-- name: UpdateVerificationStatus :exec
update user_verify
set is_verified = 1,
    verify_created_at = now()
where verify_key_hash = ?;
