/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : library

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 04/05/2023 11:03:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `phone`(`phone`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES (1, 'admin', '123456');

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`  (
  `book_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `book_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `author` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amount` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `position` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `total_amount` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `borrowed_times` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `status` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`book_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of books
-- ----------------------------
INSERT INTO `books` VALUES ('', '数据结构', '严蔚敏', 120, '03#02C#3-9', 0, 0, 0);
INSERT INTO `books` VALUES ('1', '论中国共产党历史', '中央文献出版社', 230, '03#02C#3-2', 234, 9, 1);
INSERT INTO `books` VALUES ('2', '中国共产党简史', '人民出版社', 431, '03#02C#3-1', 432, 2, 1);
INSERT INTO `books` VALUES ('2002', '简·爱', '夏洛蒂·勃朗特', 13, '02#01B#2-2', 15, 3, 1);
INSERT INTO `books` VALUES ('2003', '平凡的世界', '路遥', 105, '03#02B#3-1', 105, 0, 1);
INSERT INTO `books` VALUES ('3', '向北方', '李红梅、刘仰东', 12, '03#02B#3-2', 12, 0, 1);
INSERT INTO `books` VALUES ('4', '觉醒年代', '龙平平', 12, '03#02C#3-3', 12, 1, 1);
INSERT INTO `books` VALUES ('5', '靠山', '铁流', 34, '03#02C#3-4', 34, 0, 1);
INSERT INTO `books` VALUES ('6', '大医马海德', '陈敦德', 31, '03#02C#3-5', 32, 1, 1);
INSERT INTO `books` VALUES ('7', '数字解读中国：中国的发展坐标与发展成就', '贺耀敏、甄峰', 78, '03#02C#3-6', 78, 0, 0);
INSERT INTO `books` VALUES ('8', '三体1：地球往事', '刘慈欣', 520, '01#12A#1-1', 520, 0, 1);

-- ----------------------------
-- Table structure for borrows
-- ----------------------------
DROP TABLE IF EXISTS `borrows`;
CREATE TABLE `borrows`  (
  `id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reader_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `book_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `borrow_date` datetime(3) NULL DEFAULT NULL,
  `return_date` datetime(3) NULL DEFAULT NULL,
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `real_date` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`, `reader_id`, `book_id`) USING BTREE,
  INDEX `idx_reader_id`(`reader_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of borrows
-- ----------------------------
INSERT INTO `borrows` VALUES ('12e450c3-e6fe-11ed-98d4-38f3abf14c5a', '3f26279f-e441-11ed-b0ab-38f3abf14c5a', '2002', '2023-04-30 10:24:00.000', '2023-05-30 10:24:00.000', '已还', '2023-04-30 10:24:30.336');
INSERT INTO `borrows` VALUES ('24ebf6b0-e044-11ed-8300-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-20 23:16:21.000', '2023-05-20 23:16:21.000', '已还', '2023-04-23 23:22:04.747');
INSERT INTO `borrows` VALUES ('698ce21e-e23f-11ed-a9db-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '4', '2023-04-24 09:28:58.000', '2023-05-24 09:28:58.000', '已还', '2023-04-24 16:20:22.299');
INSERT INTO `borrows` VALUES ('6e6b6255-e2a0-11ed-a90e-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-24 21:03:43.000', '2023-05-24 21:03:43.000', '已还', '2023-04-24 21:03:46.014');
INSERT INTO `borrows` VALUES ('71a9297d-e2a1-11ed-a90e-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2002', '2023-04-24 21:10:39.000', '2023-05-24 21:10:39.000', '未还', NULL);
INSERT INTO `borrows` VALUES ('7fe7d1bd-e2a4-11ed-9135-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2', '2023-04-25 16:20:50.000', '2023-05-25 16:20:50.000', '已还', '2023-04-25 16:20:54.985');
INSERT INTO `borrows` VALUES ('a5541140-e056-11ed-bb91-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '6', '2023-04-21 23:10:13.000', '2023-05-21 23:10:13.000', '已还', '2023-04-23 23:15:16.635');
INSERT INTO `borrows` VALUES ('d13c1e25-e2a4-11ed-b26e-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2002', '2023-04-25 15:50:44.000', '2023-05-25 15:50:44.000', '续借', NULL);
INSERT INTO `borrows` VALUES ('e4f0be35-e292-11ed-833a-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2', '2023-04-24 19:26:40.000', '2023-05-24 19:26:40.000', '已还', '2023-04-24 19:43:06.602');
INSERT INTO `borrows` VALUES ('ecb3853a-e2a1-11ed-9be8-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-24 21:14:05.000', '2023-05-24 21:14:05.000', '已还', '2023-04-24 21:48:12.103');
INSERT INTO `borrows` VALUES ('ee79ee00-e6ff-11ed-a6ba-38f3abf14c5a', '3f26279f-e441-11ed-b0ab-38f3abf14c5a', '1', '2023-04-30 10:37:59.000', '2023-05-30 10:37:59.000', '已还', '2023-04-30 10:38:11.160');
INSERT INTO `borrows` VALUES ('f0d903c5-e278-11ed-abd0-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-24 19:26:08.000', '2023-05-24 19:26:08.983', '已还', '2023-04-24 21:03:21.158');

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `comment_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reader_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `book_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `date` datetime(3) NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `praise` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `status` bigint(0) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`comment_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES ('2c8f153f-e5a0-11ed-a60f-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-28 16:39:22.000', '大家好，我是小白', 2, 0);
INSERT INTO `comments` VALUES ('9454581c-e700-11ed-800b-38f3abf14c5a', '3f26279f-e441-11ed-b0ab-38f3abf14c5a', '2002', '2023-04-30 10:41:59.000', '大家好', 0, 0);
INSERT INTO `comments` VALUES ('c1d649ef-e8b2-11ed-9e80-38f3abf14c5a', 'admin', '1', '2023-05-02 14:29:57.000', '大家好，我是管理员', 0, 1);
INSERT INTO `comments` VALUES ('c7adda68-e426-11ed-b00d-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-26 19:37:52.000', '1', 0, 1);
INSERT INTO `comments` VALUES ('e54e0e02-e374-11ed-ada0-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2002', '2023-04-25 22:24:32.000', '爱了爱了', 3, 1);

-- ----------------------------
-- Table structure for readers
-- ----------------------------
DROP TABLE IF EXISTS `readers`;
CREATE TABLE `readers`  (
  `reader_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reader_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `borrow_times` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `ovd_times` bigint(0) UNSIGNED NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`reader_id`) USING BTREE,
  UNIQUE INDEX `phone`(`phone`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for reports
-- ----------------------------
DROP TABLE IF EXISTS `reports`;
CREATE TABLE `reports`  (
  `comment_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reporter_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `report_date` datetime(3) NOT NULL,
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`comment_id`, `reporter_id`, `report_date`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of reports
-- ----------------------------
INSERT INTO `reports` VALUES ('2c8f153f-e5a0-11ed-a60f-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2023-05-02 14:05:09.000', '已驳回');
INSERT INTO `reports` VALUES ('2c8f153f-e5a0-11ed-a60f-38f3abf14c5a', '3f26279f-e441-11ed-b0ab-38f3abf14c5a', '2023-05-03 11:01:12.000', '已驳回');
INSERT INTO `reports` VALUES ('2c8f153f-e5a0-11ed-a60f-38f3abf14c5a', '3f26279f-e441-11ed-b0ab-38f3abf14c5a', '2023-05-03 11:02:01.000', '已驳回');
INSERT INTO `reports` VALUES ('2c8f153f-e5a0-11ed-a60f-38f3abf14c5a', '3f26279f-e441-11ed-b0ab-38f3abf14c5a', '2023-05-03 20:31:23.000', '已驳回');
INSERT INTO `reports` VALUES ('9454581c-e700-11ed-800b-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2023-05-03 15:01:30.000', '已驳回');
INSERT INTO `reports` VALUES ('9454581c-e700-11ed-800b-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2023-05-03 20:31:12.000', '已通过');
INSERT INTO `reports` VALUES ('c1d649ef-e8b2-11ed-9e80-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2023-05-03 10:59:53.000', '已驳回');

-- ----------------------------
-- Table structure for reserves
-- ----------------------------
DROP TABLE IF EXISTS `reserves`;
CREATE TABLE `reserves`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reader_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `book_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `date` datetime(3) NULL DEFAULT NULL,
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of reserves
-- ----------------------------
INSERT INTO `reserves` VALUES ('113958d2-e6fe-11ed-98d4-38f3abf14c5a', '3f26279f-e441-11ed-b0ab-38f3abf14c5a', '2002', '2023-04-30 10:24:00.000', '已借阅');
INSERT INTO `reserves` VALUES ('2ffa972d-e440-11ed-8358-38f3abf14c5a', '2834dbc2-e440-11ed-8358-38f3abf14c5a', '1', '2023-04-26 22:39:44.000', '已借阅');
INSERT INTO `reserves` VALUES ('4eda6449-df8e-11ed-8f9f-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-20 23:16:21.000', '已借阅');
INSERT INTO `reserves` VALUES ('62eab758-e23f-11ed-a9db-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '4', '2023-04-24 09:28:58.000', '已借阅');
INSERT INTO `reserves` VALUES ('68ceec62-e2a1-11ed-a90e-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2002', '2023-04-24 21:10:39.000', '已借阅');
INSERT INTO `reserves` VALUES ('691c2671-e2a0-11ed-a90e-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-24 21:03:30.000', '已借阅');
INSERT INTO `reserves` VALUES ('7c42e4c4-e2a4-11ed-9135-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2', '2023-04-24 21:32:40.000', '已借阅');
INSERT INTO `reserves` VALUES ('908bfdab-e440-11ed-921f-38f3abf14c5a', '8b63365a-e440-11ed-921f-38f3abf14c5a', '1', '2023-04-26 22:42:26.000', '已借阅');
INSERT INTO `reserves` VALUES ('9e10aab2-e056-11ed-bb91-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '6', '2023-04-21 23:10:13.000', '已借阅');
INSERT INTO `reserves` VALUES ('ce049f94-e2a4-11ed-b26e-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2002', '2023-04-24 21:34:57.000', '已借阅');
INSERT INTO `reserves` VALUES ('e1f4762c-e292-11ed-833a-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '2', '2023-04-24 19:26:40.000', '已借阅');
INSERT INTO `reserves` VALUES ('e38cf7b1-e2a1-11ed-9be8-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-24 21:14:05.000', '已借阅');
INSERT INTO `reserves` VALUES ('ec5f417b-e6ff-11ed-a6ba-38f3abf14c5a', '3f26279f-e441-11ed-b0ab-38f3abf14c5a', '1', '2023-04-30 10:37:17.000', '已借阅');
INSERT INTO `reserves` VALUES ('ef0c74a1-e278-11ed-abd0-38f3abf14c5a', '30339572-db53-11ed-89f4-38f3abf14c5a', '1', '2023-04-24 16:20:54.000', '已借阅');
INSERT INTO `reserves` VALUES ('f4054930-e440-11ed-822e-38f3abf14c5a', 'ef188782-e440-11ed-822e-38f3abf14c5a', '1', '2023-04-26 22:45:14.000', '已借阅');

SET FOREIGN_KEY_CHECKS = 1;