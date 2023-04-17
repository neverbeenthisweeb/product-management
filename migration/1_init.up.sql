CREATE TABLE `orders_items` (
    `id` INT PRIMARY KEY AUTO_INCREMENT, 
    `name` VARCHAR(36), 
    `price` DECIMAL(10,2),
    `expired_at` TIMESTAMP, 
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    `deleted_at` TIMESTAMP NULL
);

CREATE TABLE `orders_histories` (
    `id` INT PRIMARY KEY AUTO_INCREMENT, 
    `user_id` INT,
    `order_item_id` INT, 
    `descriptions` VARCHAR(36),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    `deleted_at` TIMESTAMP NULL
);

CREATE TABLE `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT, 
    `full_name` VARCHAR(36),
    `first_order` VARCHAR(36), 
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
    `deleted_at` TIMESTAMP NULL
);
