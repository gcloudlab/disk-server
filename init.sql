/*
 Navicat Premium Data Transfer
 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 127.0.0.1:3306
 Source Schema         : gcloud
 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001
 Date: 05/05/2022 21:05:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for repository_pool
-- ----------------------------
DROP TABLE IF EXISTS `repository_pool`;
CREATE TABLE `repository_pool` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `hash` varchar(32) DEFAULT NULL COMMENT '文件的唯一标识',
  `name` varchar(255) DEFAULT NULL,
  `ext` varchar(30) DEFAULT NULL COMMENT '文件扩展名',
  `size` int(11) DEFAULT NULL COMMENT '文件大小',
  `path` varchar(255) DEFAULT NULL COMMENT '文件路径',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for share_basic
-- ----------------------------
DROP TABLE IF EXISTS `share_basic`;
CREATE TABLE `share_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `repository_identity` varchar(36) DEFAULT NULL COMMENT '公共池中的唯一标识',
  `user_repository_identity` varchar(36) DEFAULT NULL COMMENT '用户池子中的唯一标识',
  `expired_time` int(11) DEFAULT NULL COMMENT '失效时间，单位秒, 【0-永不失效】',
  `click_num` int(11) DEFAULT '0' COMMENT '点击次数',
  `desc` varchar(36) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `name` varchar(60) DEFAULT NULL,
  `password` varchar(32) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `avatar` varchar(100) DEFAULT NULL,
  `capacity` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_repository
-- ----------------------------
DROP TABLE IF EXISTS `user_repository`;
CREATE TABLE `user_repository` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `repository_identity` varchar(36) DEFAULT NULL,
  `ext` varchar(255) DEFAULT NULL COMMENT '文件或文件夹类型',
  `name` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for public_repository
-- ----------------------------
DROP TABLE IF EXISTS `public_repository`;
CREATE TABLE `public_repository` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `repository_identity` varchar(36) DEFAULT NULL,
  `ext` varchar(255) DEFAULT NULL COMMENT '文件或文件夹类型',
  `name` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for posts_basic
-- ----------------------------
DROP TABLE IF EXISTS `posts_basic`;
CREATE TABLE `posts_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `click_num` int(11) DEFAULT '0' COMMENT '点击次数',
  `title` varchar(36) DEFAULT NULL,
  `tags` varchar(255) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `mention` varchar(255) DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for posts_comment_basic
-- ----------------------------
DROP TABLE IF EXISTS `posts_comment_basic`;
CREATE TABLE `posts_comment_basic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `posts_identity` varchar(36) DEFAULT NULL,
  `reply_identity` varchar(36) DEFAULT NULL,
  `reply_name` varchar(36) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `mention` varchar(255) DEFAULT NULL,
  `like1` int(11) DEFAULT '0',
  `dislike` int(11) DEFAULT '0',
  `read` int(11) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for posts_fb
-- ----------------------------
DROP TABLE IF EXISTS `posts_fb`;
CREATE TABLE `posts_fb` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL,
  `user_identity` varchar(36) DEFAULT NULL,
  `posts_identity` varchar(36) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `count` int(11) DEFAULT '0',
  `read` int(11) DEFAULT '0',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;