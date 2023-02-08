/*
 Navicat Premium Data Transfer

 Source Server         : mysqlFromMyCloud
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : localhost:3306
 Source Schema         : dy_user

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 08/02/2023 18:24:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(0) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `follow_count` bigint(0) NULL DEFAULT 0,
  `follower_count` bigint(0) NULL DEFAULT 0,
  `created_time` timestamp(0) NOT NULL,
  `updated_time` timestamp(0) NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (446553213450061057, 'touko', 0, 0, '2023-02-06 23:07:37', '2023-02-06 23:07:37', '123456', NULL);
INSERT INTO `user` VALUES (446553238867543297, 'zhanglei', 0, 0, '2023-02-06 23:07:52', '2023-02-06 23:07:52', 'douyin', NULL);

SET FOREIGN_KEY_CHECKS = 1;
