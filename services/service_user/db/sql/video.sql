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

 Date: 08/02/2023 18:24:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` bigint(0) NOT NULL,
  `user_id` bigint(0) NULL DEFAULT NULL,
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `favorite_count` bigint(0) NULL DEFAULT NULL,
  `comment_count` bigint(0) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `user->video` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (446684273202693377, 446553213450061057, 'https://douyin-1313537069.cos.ap-guangzhou.myqcloud.com/video/2023/2/7/OBS_OBSS 无法获取游戏源、黑屏的解决办法！-百度经验 和另外 5 个页面 - 个人 - Microsoft​ Edge 2022-06-17 23-03-14.mp4', 'https://douyin-1313537069.cos.ap-guangzhou.myqcloud.com/picture/cover.jpg', 10, 10, '我的第一个视频', '2023-02-07 20:49:34');
INSERT INTO `video` VALUES (446685321345006205, 446553213450061057, 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 1, 2, '测试视频', '2023-02-06 12:00:00');
INSERT INTO `video` VALUES (446806371942270209, 446553213450061057, 'https://douyin-1313537069.cos.ap-guangzhou.myqcloud.com/video/2023/2/8/8cb90570953eca096383df5b1990172f.mp4', 'https://douyin-1313537069.cos.ap-guangzhou.myqcloud.com/picture/8d6fdf68df7d6d99f1fcd8c318482736.jpg', 0, 0, '人来人往', '2023-02-08 17:02:31');

SET FOREIGN_KEY_CHECKS = 1;
