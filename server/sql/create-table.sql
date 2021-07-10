use `ecommerce_fashion`;
CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE NOT NULL,
  `email` varchar(255),
  `password` varchar(255) NOT NULL,
  `role` ENUM ('admin', 'customer') NOT NULL
);
CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255)
);
CREATE TABLE `products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  `price` int,
  `sale_percent` float,
  `category_id` int,
  `quantity` int
);
CREATE TABLE `product_properties` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `key` varchar(255) NOT NULL,
  `value` varchar(255) NOT NULL
);
CREATE TABLE `product_images` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` int,
  `uri` varchar(255) NOT NULL
);
CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `status` ENUM (
    'Chờ thanh toán',
    'Đang vận chuyển',
    'Đã thanh toán',
    'Giao hàng thành công',
    'Đã hủy đơn hàng'
  ) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255),
  `phone` varchar(11) NOT NULL,
  `location` varchar(255) NOT NULL,
  `total_price` int
);
CREATE TABLE `order_items` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` int,
  `order_id` int,
  `quantity` int
);
CREATE TABLE `carts` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int
);
CREATE TABLE `cart_items` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `cart_id` int,
  `product_id` int,
  `quantity` int
);
ALTER TABLE
  `products`
ADD
  FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);
ALTER TABLE
  `product_properties`
ADD
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
ALTER TABLE
  `product_images`
ADD
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
ALTER TABLE
  `orders`
ADD
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE
  `order_items`
ADD
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
ALTER TABLE
  `order_items`
ADD
  FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);
ALTER TABLE
  `carts`
ADD
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
ALTER TABLE
  `cart_items`
ADD
  FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`);
ALTER TABLE
  `cart_items`
ADD
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
