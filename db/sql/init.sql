DROP DATABASE IF EXISTS test_db;
CREATE DATABASE test_db;
USE test_db;
CHARSET utf8mb4;

CREATE TABLE `users` (
  `id` VARCHAR(32) PRIMARY KEY UNIQUE NOT NULL,
  `name` VARCHAR(32) NOT NULL,
  `gender` TINYINT(1) UNSIGNED NOT NULL,
  `age` TINYINT(1) UNSIGNED NOT NULL
);

CREATE TABLE `coordinates` (
  `id` INT AUTO_INCREMENT PRIMARY KEY UNIQUE NOT NULL,
  `image` VARCHAR(32) NOT NULL,
  `user_id` VARCHAR(32) NOT NULL,
  FOREIGN KEY(`user_id`) REFERENCES users(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  `created_at` TIMESTAMP NOT NULL
);

CREATE TABLE `favorites` (
  `id` INT AUTO_INCREMENT PRIMARY KEY UNIQUE NOT NULL,
  `flag` BIT(1) NOT NULL DEFAULT b'0',
  `coordinate_id` INT NOT NULL,
  FOREIGN KEY(`coordinate_id`) REFERENCES coordinates(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  `user_id` VARCHAR(32) NOT NULL,
  FOREIGN KEY(`user_id`) REFERENCES users(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- user
insert into `users` (`id`, `name`, `gender`, `age`) values
  ('tQKMbEoRLicSYF0QhfTQaDIpz2e2', 'コード001', `0`, `22`);
insert into `users` (`id`, `name`, `gender`, `age`) values
  ('QHvLBq764VYfpBBUFp0wstryG833', 'コード002', `1`, `32`);
