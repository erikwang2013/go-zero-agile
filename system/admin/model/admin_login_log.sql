CREATE TABLE `admin_login_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `admin_id` int NOT NULL,
  `login_ip` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `login_time` timestamp NOT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;