# Đây là phần hướng dẫn cài đặt backend cho hệ thống

    1. Tạo cơ sở dữ liệu trong mysql, file sql về các bảng và dữ liệu để có trong thư mục ```sql```
    2. Chạy lệnh ```go mod tidy``` để tải các các dependency mà trong project sử dụng
    3. Mở file ```dev.yml``` để để chỉnh sửa cấu hình phù hợp như port, username, password mysql,...
    4. Chạy lệnh ```go run app.go``` để khởi chạy server
