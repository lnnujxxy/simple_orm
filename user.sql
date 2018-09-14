CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `age` smallint(10) unsigned NOT NULL DEFAULT 0 COMMENT '年龄',
  `first_name` varchar(45) NOT NULL DEFAULT '' COMMENT '姓',
  `last_name` varchar(45) NOT NULL DEFAULT '' COMMENT '名',
  `email` varchar(45) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_age` (`age`),
  KEY `idx_first_name` (`first_name`),
  KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';