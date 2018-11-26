/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50613
Source Host           : localhost:3306
Source Database       : gonovel

Target Server Type    : MYSQL
Target Server Version : 50613
File Encoding         : 65001

Date: 2018-11-26 09:46:10
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS `book`;
CREATE TABLE `book` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `author` varchar(100) DEFAULT NULL,
  `image` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of book
-- ----------------------------
INSERT INTO `book` VALUES ('4', '圣墟', 'huangzs', '', '2018-11-21 02:57:13', '2018-11-21 02:57:13');

-- ----------------------------
-- Table structure for chapter
-- ----------------------------
DROP TABLE IF EXISTS `chapter`;
CREATE TABLE `chapter` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_id` int(11) DEFAULT NULL,
  `url` varchar(255) NOT NULL,
  `title` varchar(100) DEFAULT NULL,
  `content` text,
  `sort` int(11) DEFAULT NULL,
  `pre` int(11) DEFAULT NULL,
  `next` int(11) DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `url_unq` (`url`)
) ENGINE=InnoDB AUTO_INCREMENT=6579 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of chapter
-- ----------------------------

-- ----------------------------
-- Table structure for undownload
-- ----------------------------
DROP TABLE IF EXISTS `undownload`;
CREATE TABLE `undownload` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `error_book_id` int(11) DEFAULT NULL,
  `chapter_title` varchar(255) DEFAULT NULL,
  `info` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of undownload
-- ----------------------------
