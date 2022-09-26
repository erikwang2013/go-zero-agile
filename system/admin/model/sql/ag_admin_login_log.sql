CREATE TABLE `ag_admin_login_log` (
  `id` bigint unsigned NOT NULL,
  `admin_id` int NOT NULL,
  `access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `login_ip` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `login_time` timestamp NOT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;