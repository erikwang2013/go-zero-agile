CREATE TABLE `ag_role_permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL,
  `permission_id` int NOT NULL,
  `permission_childer_id` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '子权限 多个用英文逗号隔开',
  `permission_button` json DEFAULT NULL COMMENT '按钮json',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 0=开启 1=关闭',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删 0=否 1=是',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色权限关系表';