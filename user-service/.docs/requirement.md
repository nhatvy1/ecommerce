# Tài liệu thiết kế cơ sở dữ liệu quản lý phiên đăng nhập

## 1. Giới thiệu
Tài liệu này mô tả thiết kế cơ sở dữ liệu (database) cho hệ thống quản lý phiên đăng nhập trên từng thiết bị. Các bảng và mối quan hệ giữa chúng được mô tả chi tiết để hỗ trợ việc phát triển.

---

## 2. Các bảng trong cơ sở dữ liệu

### 2.1. Bảng `Users`
Lưu trữ thông tin người dùng.

| Tên cột       | Kiểu dữ liệu | Mô tả                     |
|---------------|--------------|---------------------------|
| `user_id`     | INT (PK)     | ID duy nhất của người dùng|
| `username`    | VARCHAR(50)  | Tên đăng nhập             |
| `password`    | VARCHAR(255) | Mật khẩu (đã mã hóa)      |
| `email`       | VARCHAR(100) | Email người dùng          |
| `created_at`  | DATETIME     | Thời gian tạo tài khoản   |

### 2.2. Bảng `Devices`
Lưu trữ thông tin về các thiết bị mà người dùng đăng nhập.

| Tên cột       | Kiểu dữ liệu | Mô tả                     |
|---------------|--------------|---------------------------|
| `device_id`   | INT (PK)     | ID duy nhất của thiết bị  |
| `user_id`     | INT (FK)     | ID người dùng             |
| `device_name` | VARCHAR(100) | Tên thiết bị              |
| `device_type` | VARCHAR(50)  | Loại thiết bị (PC, Mobile)|
| `created_at`  | DATETIME     | Thời gian thiết bị được thêm|

### 2.3. Bảng `Sessions`
Lưu trữ thông tin về các phiên đăng nhập của người dùng trên từng thiết bị.

| Tên cột       | Kiểu dữ liệu | Mô tả                     |
|---------------|--------------|---------------------------|
| `session_id`  | INT (PK)     | ID duy nhất của phiên     |
| `user_id`     | INT (FK)     | ID người dùng             |
| `device_id`   | INT (FK)     | ID thiết bị               |
| `login_time`  | DATETIME     | Thời gian đăng nhập       |
| `logout_time` | DATETIME     | Thời gian đăng xuất       |
| `ip_address`  | VARCHAR(45)  | Địa chỉ IP đăng nhập      |
| `is_active`   | BOOLEAN      | Trạng thái phiên (active/inactive)|

### 2.4. Bảng `Session_Tokens`
Lưu trữ token để xác thực phiên đăng nhập.

| Tên cột       | Kiểu dữ liệu | Mô tả                     |
|---------------|--------------|---------------------------|
| `token_id`    | INT (PK)     | ID duy nhất của token     |
| `session_id`  | INT (FK)     | ID phiên đăng nhập        |
| `token`       | VARCHAR(255) | Token xác thực            |
| `expires_at`  | DATETIME     | Thời gian hết hạn token   |

### 2.5. Bảng `Activity_Logs`
Lưu trữ các hoạt động của người dùng trong phiên đăng nhập.

| Tên cột       | Kiểu dữ liệu | Mô tả                     |
|---------------|--------------|---------------------------|
| `log_id`      | INT (PK)     | ID duy nhất của log       |
| `session_id`  | INT (FK)     | ID phiên đăng nhập        |
| `activity`    | TEXT         | Mô tả hoạt động           |
| `timestamp`   | DATETIME     | Thời gian hoạt động       |

---

## 3. Mối quan hệ giữa các bảng
- **Users** và **Devices**: Một người dùng có thể có nhiều thiết bị (`user_id` trong bảng `Devices` là khóa ngoại tham chiếu đến `user_id` trong bảng `Users`).
- **Users** và **Sessions**: Một người dùng có thể có nhiều phiên đăng nhập (`user_id` trong bảng `Sessions` là khóa ngoại tham chiếu đến `user_id` trong bảng `Users`).
- **Devices** và **Sessions**: Một thiết bị có thể có nhiều phiên đăng nhập (`device_id` trong bảng `Sessions` là khóa ngoại tham chiếu đến `device_id` trong bảng `Devices`).
- **Sessions** và **Session_Tokens**: Một phiên đăng nhập có thể có nhiều token (`session_id` trong bảng `Session_Tokens` là khóa ngoại tham chiếu đến `session_id` trong bảng `Sessions`).
- **Sessions** và **Activity_Logs**: Một phiên đăng nhập có thể có nhiều hoạt động (`session_id` trong bảng `Activity_Logs` là khóa ngoại tham chiếu đến `session_id` trong bảng `Sessions`).

---

## 4. SQL để tạo các bảng
Dưới đây là mã SQL để tạo các bảng trong cơ sở dữ liệu:

```sql
CREATE TABLE Users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Devices (
    device_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    device_name VARCHAR(100),
    device_type VARCHAR(50),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE Sessions (
    session_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    device_id INT,
    login_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    logout_time DATETIME,
    ip_address VARCHAR(45),
    is_active BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (device_id) REFERENCES Devices(device_id)
);

CREATE TABLE Session_Tokens (
    token_id INT AUTO_INCREMENT PRIMARY KEY,
    session_id INT,
    token VARCHAR(255) NOT NULL,
    expires_at DATETIME NOT NULL,
    FOREIGN KEY (session_id) REFERENCES Sessions(session_id)
);

CREATE TABLE Activity_Logs (
    log_id INT AUTO_INCREMENT PRIMARY KEY,
    session_id INT,
    activity TEXT,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (session_id) REFERENCES Sessions(session_id)
);