/*
Navicat MySQL Data Transfer

Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50639
File Encoding         : 65001

Date: 2018-03-18 16:52:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
                                `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
                                `title` varchar(100) DEFAULT '' COMMENT '文章标题',
                                `desc` varchar(255) DEFAULT '' COMMENT '简述',
                                `content` text COMMENT '内容',
                                `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
                                `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
                                `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                                `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                                `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
                                `deleted_on` int(10) unsigned DEFAULT '0',
                                `state` tinyint(3) unsigned DEFAULT '1' COMMENT '删除时间',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `username` varchar(50) DEFAULT '' COMMENT '账号',
                             `password` varchar(50) DEFAULT '' COMMENT '密码',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES ('1', 'test', 'test123');

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
                            `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(100) DEFAULT '' COMMENT '标签名称',
                            `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
                            `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
                            `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
                            `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
                            `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
                            `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

CREATE TABLE `match_record` (
                                `id` bigint NOT NULL AUTO_INCREMENT,
                                `male_id` bigint NOT NULL DEFAULT '0' COMMENT '男生id',
                                `female_id` bigint NOT NULL DEFAULT '0' COMMENT '女生id',
                                `male_email` varchar(100) NOT NULL DEFAULT '' COMMENT '男生邮箱',
                                `female_email` varchar(100) NOT NULL DEFAULT '' COMMENT '女生邮箱',
                                `create_time` bigint NOT NULL DEFAULT '0' COMMENT '匹配时间',
                                `send_status` tinyint NOT NULL DEFAULT '0' COMMENT '邮件发送状态 0 未发送 1 已发送',
                                `confirm_status` tinyint NOT NULL COMMENT '配对确认状态 0 未联系 1 已联系',
                                `match_result` tinyint NOT NULL DEFAULT '0' COMMENT '匹配结果 0 未确认 1 合适 2 不合适',
                                `male_attitude` tinyint NOT NULL DEFAULT '0' COMMENT '男生意见 0 未确认 1 合适 2 不合适',
                                `female_attitude` tinyint NOT NULL DEFAULT '0' COMMENT '女生意见 0 未确认 1 合适 2 不合适',
                                `male_note` text NOT NULL COMMENT '男生反馈',
                                `female_note` text NOT NULL COMMENT '女生反馈',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='配对记录'


create table candidate
(
    id bigint auto_increment
        primary key,
    gender tinyint default 0 not null comment '0 未知 1 男生 2 女生',
    email varchar(100) default '' not null comment '邮箱',
    create_time bigint default 0 not null comment '创建时间',
    statue tinyint default 0 not null comment '0 正常 1 已删除',
    update_time bigint default 0 not null comment '更新时间'
)
    comment '候选人';


