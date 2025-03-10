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

CREATE TABLE `customers` (
  `customerNumber` int(11) NOT NULL,
  `customerName` varchar(50) NOT NULL,
  `contactLastName` varchar(50) NOT NULL,
  `contactFirstName` varchar(50) NOT NULL,
  `phone` varchar(50) NOT NULL,
  `addressLine1` varchar(50) NOT NULL,
  `addressLine2` varchar(50) DEFAULT NULL,
  `city` varchar(50) NOT NULL,
  `state` varchar(50) DEFAULT NULL,
  `postalCode` varchar(15) DEFAULT NULL,
  `country` varchar(50) NOT NULL,
  `salesRepEmployeeNumber` int(11) DEFAULT NULL,
  `creditLimit` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`customerNumber`),
  KEY `salesRepEmployeeNumber` (`salesRepEmployeeNumber`),
  CONSTRAINT `customers_ibfk_1` FOREIGN KEY (`salesRepEmployeeNumber`) REFERENCES `employees` (`employeeNumber`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE `employees` (
  `employeeNumber` int(11) NOT NULL,
  `lastName` varchar(50) NOT NULL,
  `firstName` varchar(50) NOT NULL,
  `extension` varchar(10) NOT NULL,
  `email` varchar(100) NOT NULL,
  `officeCode` varchar(10) NOT NULL,
  `reportsTo` int(11) DEFAULT NULL,
  `jobTitle` varchar(50) NOT NULL,
  PRIMARY KEY (`employeeNumber`),
  KEY `reportsTo` (`reportsTo`),
  KEY `officeCode` (`officeCode`),
  CONSTRAINT `employees_ibfk_1` FOREIGN KEY (`reportsTo`) REFERENCES `employees` (`employeeNumber`),
  CONSTRAINT `employees_ibfk_2` FOREIGN KEY (`officeCode`) REFERENCES `offices` (`officeCode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
  `officeCode` varchar(10) NOT NULL,
  `city` varchar(50) NOT NULL,
  `phone` varchar(50) NOT NULL,
  `addressLine1` varchar(50) NOT NULL,
  `addressLine2` varchar(50) DEFAULT NULL,
  `state` varchar(50) DEFAULT NULL,
  `country` varchar(50) NOT NULL,
  `postalCode` varchar(15) NOT NULL,
  `territory` varchar(10) NOT NULL,
  PRIMARY KEY (`officeCode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE `orderdetails` (
  `orderNumber` int(11) NOT NULL,
  `productCode` varchar(15) NOT NULL,
  `quantityOrdered` int(11) NOT NULL,
  `priceEach` decimal(10,2) NOT NULL,
  `orderLineNumber` smallint(6) NOT NULL,
  PRIMARY KEY (`orderNumber`,`productCode`),
  KEY `productCode` (`productCode`),
  CONSTRAINT `orderdetails_ibfk_1` FOREIGN KEY (`orderNumber`) REFERENCES `orders` (`orderNumber`),
  CONSTRAINT `orderdetails_ibfk_2` FOREIGN KEY (`productCode`) REFERENCES `products` (`productCode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE `orders` (
  `orderNumber` int(11) NOT NULL,
  `orderDate` date NOT NULL,
  `requiredDate` date NOT NULL,
  `shippedDate` date DEFAULT NULL,
  `status` varchar(15) NOT NULL,
  `comments` text,
  `customerNumber` int(11) NOT NULL,
  PRIMARY KEY (`orderNumber`),
  KEY `customerNumber` (`customerNumber`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`customerNumber`) REFERENCES `customers` (`customerNumber`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE `payments` (
  `customerNumber` int(11) NOT NULL,
  `checkNumber` varchar(50) NOT NULL,
  `paymentDate` date NOT NULL,
  `amount` decimal(10,2) NOT NULL,
  PRIMARY KEY (`customerNumber`,`checkNumber`),
  CONSTRAINT `payments_ibfk_1` FOREIGN KEY (`customerNumber`) REFERENCES `customers` (`customerNumber`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE TABLE `products` (
  `productCode` varchar(15) NOT NULL,
  `productName` varchar(70) NOT NULL,
  `productLine` varchar(50) NOT NULL,
  `productScale` varchar(10) NOT NULL,
  `productVendor` varchar(50) NOT NULL,
  `productDescription` text NOT NULL,
  `quantityInStock` smallint(6) NOT NULL,
  `buyPrice` decimal(10,2) NOT NULL,
  `MSRP` decimal(10,2) NOT NULL,
  PRIMARY KEY (`productCode`),
  KEY `productLine` (`productLine`),
  CONSTRAINT `products_ibfk_1` FOREIGN KEY (`productLine`) REFERENCES `productlines` (`productLine`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
hãy giúp tôi đặt nhưng câu hỏi và bài tập liên quan thiết database trên