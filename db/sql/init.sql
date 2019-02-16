DROP DATABASE IF EXISTS test;
CREATE DATABASE test;
USE test;
CHARSET utf8mb4;

CREATE TABLE `users` (
  `id` VARCHAR(32) PRIMARY KEY UNIQUE NOT NULL,
  `name` VARCHAR(32) NOT NULL,
  `sex` BIT(1) NOT NULL,
  `age` TINYINT(1) UNSIGNED NOT NULL
);

CREATE TABLE `coordinates` (
  `id` INT AUTO_INCREMENT PRIMARY KEY UNIQUE NOT NULL,
  `image_name` VARCHAR(32) UNIQUE NOT NULL,
  `user_id` VARCHAR(32) NOT NULL,
  FOREIGN KEY(`user_id`) REFERENCES users(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  `created_at` TIMESTAMP NOT NULL
);

CREATE TABLE `favorites` (
  `coordinate_id` INT NOT NULL,
  FOREIGN KEY(`coordinate_id`) REFERENCES coordinates(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  `user_id` VARCHAR(32) NOT NULL,
  FOREIGN KEY(`user_id`) REFERENCES users(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY(`coordinate_id`, `user_id`)
);

insert into `users` (`id`, `name`, `sex`, `age`) values
  ('tQKMbEoRLicSYF0QhfTQaDIpz2e2', 'コード000', 1, 49),
  ('a1', 'コード001', 0, 22),
  ('a2', 'コード002', 1, 40),
  ('a3', 'コード003', 1, 32);

insert into `coordinates` (`id`, `image_name`, `user_id`, `created_at`) values
  (1001, 'woman1.jpg', 'a1', '2019-01-20 09:27:30'),
  (1002, 'woman2.jpg', 'a1', '2019-01-21 09:27:30'),
  (1003, 'man1.jpg', 'tQKMbEoRLicSYF0QhfTQaDIpz2e2', '2019-01-20 09:27:30'),
  (1004, 'man2.jpg', 'tQKMbEoRLicSYF0QhfTQaDIpz2e2', '2019-01-21 09:27:30');

insert into `favorites` (`coordinate_id`, `user_id`) values
  (1001, 'tQKMbEoRLicSYF0QhfTQaDIpz2e2'),
  (1001, 'a2'),
  (1001, 'a3'),
  (1002, 'a2'),
  (1002, 'a3'),
  (1003, 'a2');
