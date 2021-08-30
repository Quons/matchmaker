CREATE TABLE `candidate` (
                             `id` bigint NOT NULL AUTO_INCREMENT,
                             `gender` tinyint NOT NULL DEFAULT '0' COMMENT '0 未知 1 男生 2 女生',
                             `email` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
                             `create_time` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
                             `statue` tinyint NOT NULL DEFAULT '0' COMMENT '0 正常 1 已删除',
                             `update_time` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
                             `age` int NOT NULL DEFAULT '0' COMMENT '年龄',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='候选人';

CREATE TABLE `match_record` (
                                `id` bigint NOT NULL AUTO_INCREMENT,
                                `male_id` bigint NOT NULL DEFAULT '0' COMMENT '男生id',
                                `female_id` bigint NOT NULL DEFAULT '0' COMMENT '女生id',
                                `male_email` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '男生邮箱',
                                `female_email` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '女生邮箱',
                                `create_time` bigint NOT NULL DEFAULT '0' COMMENT '匹配时间',
                                `male_send_status` tinyint NOT NULL DEFAULT '0' COMMENT '男生邮件发送状态 0 未发送 1 已发送',
                                `female_send_status` int NOT NULL DEFAULT '0' COMMENT '男生邮件发送状态 0 未发送 1 已发送',
                                `match_result` tinyint NOT NULL DEFAULT '0' COMMENT '匹配结果 0 未确认 1 合适 2 不合适',
                                `male_attitude` tinyint NOT NULL DEFAULT '0' COMMENT '男生意见 0 未确认 1 合适 2 不合适',
                                `female_attitude` tinyint NOT NULL DEFAULT '0' COMMENT '女生意见 0 未确认 1 合适 2 不合适',
                                `male_note` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '男生反馈',
                                `female_note` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '女生反馈',
                                `confirm_status` tinyint NOT NULL COMMENT '配对确认状态 0 未联系 1 已联系',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=155 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='配对记录';



