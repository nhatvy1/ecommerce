1. Hash email lưu vào Redis
- Sử dụng key: u:email:otp

2. Giải pháp sử dụng redis để chống spam OTP
- Mỗi email chỉ được tạo 1 OTP trong 90 giây
- Nếu người dùng spam, hệ thống sẽ trả về OTP cũ

3. Sử dụng cơ chế lock trong redis
a. Cơ chế lock trong redis
- Key u:email:otp
- Được tạo khi OTP mới sinh ra
- TTL = 90 giây, ngăn không cho tạo OTP mới trong khoảng thời gian này
- Kh hết hạn, người dùng có thể yêu cầu OTP mới
b. Xử lý người dùng spam
- Nếu lockey tồn tại: 
  + Trả về OTPcũ (nếu có)
  + Hoặc thông báo lỗi: "Đợi 90 giây"
c. Đảm bảo tính nhất quán
- Sử dụng SET key value EX 90 NX (trong Redis) để tránh race condition: