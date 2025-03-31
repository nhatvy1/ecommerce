-- name: InsertUserVerify :exec
insert into user_verify (
  verify_otp,
  verify_key,
  verify_key_hash,
  verify_type
) values (
  ?, ?, ?, ?
)