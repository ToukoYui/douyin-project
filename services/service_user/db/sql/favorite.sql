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

 Date: 08/02/2023 18:23:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`  (
  `id` bigint(0) NOT NULL,
  `user_id` bigint(0) NULL DEFAULT NULL,
  `video_id` bigint(0) NULL DEFAULT NULL,
  `is_favorite` tinyint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `favorite_user`(`user_id`) USING BTREE,
  INDEX `favorite_video`(`video_id`) USING BTREE,
  CONSTRAINT `favorite_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `favorite_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite
-- ----------------------------
INSERT INTO `favorite` VALUES (446808854836020481, 446553213450061057, 446684273202693377, 1);
INSERT INTO `favorite` VALUES (446808900906255617, 446553213450061057, 446685321345006205, 0);
INSERT INTO `favorite` VALUES (446808995831743745, 446553213450061057, 446806371942270209, 1);

SET FOREIGN_KEY_CHECKS = 1;
