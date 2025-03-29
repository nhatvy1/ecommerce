-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_info` (
    verify_id INT AUTO_INCREMENT PRIMARY KEY,             -- ID of the OTP record
    verify_otp VARCHAR(6) NOT NULL,                       -- OTP code (verification code)
    verify_key VARCHAR(255) NOT NULL,                     -- verify_key: User's email (or phone number) to identify the OTP recipient
    verify_key_hash VARCHAR(255) NOT NULL,                -- verify_key_hash: User's email (or phone number) to identify the OTP recipient
    verify_type INT DEFAULT 1,                            -- 1: Email, 2: Phone, 3:... (Type of verification)
    is_verified INT DEFAULT 0,                            -- 0: No, 1: Yes - OTP verification status (default is not verified)
    is_deleted INT DEFAULT 0,                             -- 0: No, 1: Yes - Deletion status
    verify_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation time
    verify_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Record update time

    INDEX idx_verify_otp (verify_otp),

    UNIQUE KEY unique_verify_key (verify_key)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_info`;
-- +goose StatementEnd

