CREATE TABLE `ag_admin_role_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `admin_id` int NOT NULL,
  `role_id` int NOT NULL,
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 0=开启 1=关闭',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删 0=否 1=是',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员角色组';