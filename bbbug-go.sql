/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : bbbug-go

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 08/10/2023 16:57:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sa_access
-- ----------------------------
DROP TABLE IF EXISTS `sa_access`;
CREATE TABLE `sa_access`  (
  `access_id` int NOT NULL AUTO_INCREMENT,
  `access_user` int NOT NULL DEFAULT 0 COMMENT '用户ID',
  `access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'AccessToken',
  `access_plat` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'all' COMMENT '登录平台',
  `access_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'IP',
  `access_status` int NOT NULL DEFAULT 0 COMMENT '状态',
  `access_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `access_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`access_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '授权信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_access
-- ----------------------------
INSERT INTO `sa_access` VALUES (9, 0, '', 'all', '', 0, 12345678, 12345678);
INSERT INTO `sa_access` VALUES (10, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3MDM1LCJpYXQiOjE2NTg5OTIyMzV9.MORvr-18avVe8Wkehe-GexkitDlJYz7EDMv3yfLaT6M', 'test', '34.92.112.221', 0, 0, 0);
INSERT INTO `sa_access` VALUES (11, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3MDgyLCJpYXQiOjE2NTg5OTIyODJ9.2sB3DD4CJAwjDL9oeYq4NU7-CqEIb8dsKg26Ci5k4Ag', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (12, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3MTAyLCJpYXQiOjE2NTg5OTIzMDJ9.CHqlXUBv3HmHHWwPUUu4s_bAdFOKcyB_erRdJ0s9qXg', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (13, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3MTA1LCJpYXQiOjE2NTg5OTIzMDV9.6fBMi9vGgBBm-Ta5gQRY8TfSlM8UCZzuIxKEPcfWjLE', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (14, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3MTE1LCJpYXQiOjE2NTg5OTIzMTV9.yVGkC1OyyM6olwh1CMSzT9uLvNxOe9_EGipK9-XbCNw', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (15, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NDMyLCJpYXQiOjE2NTg5OTI2MzJ9.RXp5Ox7i51WgHny_AbEmBjd4nEwK56tLb6HAN4JU_v4', 'test', '34.92.112.221', 0, 0, 0);
INSERT INTO `sa_access` VALUES (16, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NDg4LCJpYXQiOjE2NTg5OTI2ODh9.vkM48RtdGq0TEg52rxCuio_s1ltw46Viak5CfuZqr04', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (17, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NDk0LCJpYXQiOjE2NTg5OTI2OTR9.nPIea-SUxVuk1InJ56yo_CWPaLppddCvgqZM_niau0M', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (18, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NTAzLCJpYXQiOjE2NTg5OTI3MDN9.hL9EL_JdXplFBi7zKdXeFWu6XSVa__uo9JGkoKOInvE', 'test', '34.92.112.221', 0, 0, 0);
INSERT INTO `sa_access` VALUES (19, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NTE0LCJpYXQiOjE2NTg5OTI3MTR9.1xxRgU1VtvothhvYuiFS4HUopa4j6z11G44Odh_caAY', 'test', '34.92.112.221', 0, 0, 0);
INSERT INTO `sa_access` VALUES (20, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NTU1LCJpYXQiOjE2NTg5OTI3NTV9.Bs_87M97xG9xEExaeWdIkil0tqTbT3x51kci_lOBtlM', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (21, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NTY0LCJpYXQiOjE2NTg5OTI3NjR9.SfaTWCQt_IIrDU-f9O9nOgYJqt4BRU6MHZj9vE9wzFM', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (22, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NTcyLCJpYXQiOjE2NTg5OTI3NzJ9.goQFLJGNo-lPoTWw84hxaChZnQSCPbuSJ29AGlcUc4A', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (23, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NTc1LCJpYXQiOjE2NTg5OTI3NzV9.5umrq61JJn8MRr9pyruapZXhqRziKQo2ye3qEwxx_es', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (24, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NjM4LCJpYXQiOjE2NTg5OTI4Mzh9.IBxV2eeJjs9i34Rp5cyWPSpCzfhao9vkHJVnwZNqohU', 'test', '34.92.112.221', 0, 0, 0);
INSERT INTO `sa_access` VALUES (25, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NTk3NjQ3LCJpYXQiOjE2NTg5OTI4NDd9.YVGu8Fb8MI6DWhC8yBWR2OVmI7tky9IwqE5bzBYnCjU', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (26, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NjA3NjcxLCJpYXQiOjE2NTkwMDI4NzF9.Vlfw15ZvnwXJpLy2SIza9pOJo1KNADRR1knjchCQLtw', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (27, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQbGF0IjoidnVlIiwiVXNlcklkIjoxLCJVc2VyTmFtZSI6ImFkbWluQGJiYnVnLmNvbSIsImV4cCI6MTY1OTIzNzE1OSwiaWF0IjoxNjU5MDY0MzU5fQ.eFiNNOGwoiTbiwUKHiQFUNJZ2KBnJqzLmG8TKQ6l7Lc', 'vue', '121.4.166.249', 0, 1659064359, 1659064359);
INSERT INTO `sa_access` VALUES (28, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NjcwMjU2LCJpYXQiOjE2NTkwNjU0NTZ9.l8fuMKALDZsNIhBROOim7dcCW9DrUAitwtjV3_gfZKc', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (29, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NjcwMjc1LCJpYXQiOjE2NTkwNjU0NzV9.QczzhtlTpYDpWDmkJ_2JtoBoyslxPAIrSOTa3PKqCo8', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (30, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5NjcwNzQ3LCJpYXQiOjE2NTkwNjU5NDd9.xB5Wf87flC_QYDk0LQ094d4udu3V3Our-vzj6plcu4w', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (31, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY1OTY4OTM3MCwiaWF0IjoxNjU5MDg0NTcwfQ.RkNsUuGjD9Rv8GWgpppAmX2pDd-5y7X9A0VZNwA1mFM', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (32, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjU5Njk1MTcyLCJpYXQiOjE2NTkwOTAzNzJ9.SufACaHFu6ss6LDhUGHqIgLuZlO-VFJ-1_NEjR9b12c', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (33, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAyNjI0NCwiaWF0IjoxNjU5NDIxNDQ0fQ.0W1LrX98LJQhAidEaJZzIkQ1gNGZ_Y1oWAQlV3rCYL0', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (34, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAyNzM5MCwiaWF0IjoxNjU5NDIyNTkwfQ.iYXtrp9lmcJwapau9IZrWs4LQXbBPjhJFyD0wnx62Xw', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (35, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYwMDI3NDEyLCJpYXQiOjE2NTk0MjI2MTJ9.jN-oAcDi_lMXeElKFTvkARf-k-7jWVLi6-VS9o3WN9U', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (36, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMjcwNSwiaWF0IjoxNjU5NDI3OTA1fQ.iFkrPvT8HVHUzPug0Tt2kGooFO1a0hns7nGfZs3PbGc', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (37, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMjgzNCwiaWF0IjoxNjU5NDI4MDM0fQ.oxr329TYdg7-bZDHYfMCiDHGTlKIKirnGSHH9xMx1cY', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (38, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMjkxOCwiaWF0IjoxNjU5NDI4MTE4fQ.Z_6dFOMxYETriQISRt41oruDKRI9-I5kOlu5y-CUN18', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (39, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMzAwMywiaWF0IjoxNjU5NDI4MjAzfQ.CZMsrBWmX8iCG3--tjuAHa5NlZTh5fXnu8QHyGJ7cvM', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (40, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMzAxMSwiaWF0IjoxNjU5NDI4MjExfQ.4rDXcdRl_qfEjZaLOhF8vCwOiqKio3-Vb05-OI37wGk', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (41, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMzAzMiwiaWF0IjoxNjU5NDI4MjMyfQ.gTs0JieFAgy7-3UvA0MdjyIhizdheTIpGAvhL-vn_AQ', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (42, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMzQ4MCwiaWF0IjoxNjU5NDI4NjgwfQ.rRrUPA8UH5dN5FMSQ3fiyw6dwvnssUCX0HGk71nSblE', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (43, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMzU1MCwiaWF0IjoxNjU5NDI4NzUwfQ.cyY9Z6QBluM0zLPxvgZoBxL5fWAaUiVeYOu0UuCQi7c', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (44, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzMzg4OCwiaWF0IjoxNjU5NDI5MDg4fQ.AZ0yI5BYGuJgJ6AMd6lcnSDTC1duxNL4DEvXcrChcUk', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (45, 200, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwMCwiVXNlck5hbWUiOiJ0ZXN0QHFxLmNvbSIsIlVzZXJQbGF0IjoidGVzdCIsImV4cCI6MTY2MDAzNDAyNSwiaWF0IjoxNjU5NDI5MjI1fQ.ePD4fYmAkrjmsUO96cIYofW-3tNi8TKrtp6zIaDC0iI', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (46, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYwMDM0MTc0LCJpYXQiOjE2NTk0MjkzNzR9.FPVn40S0iaSNNMx4YLaI-HQSJ_49uc5ok6Yafm-MHHA', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (47, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYwMDM3NDg5LCJpYXQiOjE2NTk0MzI2ODl9.2RA8GOVbR6WqWvVLhBf8tZ7508RpHtIQzOo_x5h4zew', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (48, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYwMTAzNjgzLCJpYXQiOjE2NTk0OTg4ODN9.xC9-q2W7tDIbPt773r_rXSY6CUqr3lKTM3F4dFhvK0U', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (49, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYwMTE1OTg2LCJpYXQiOjE2NTk1MTExODZ9.0rhGYFS7k0Z-xAmwEOGFU4fbdol7Rr4JWl5BTf4HNGk', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (50, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYwMTE2MTUzLCJpYXQiOjE2NTk1MTEzNTN9.NBXm2yGw4t6Ig6eH40pnce9YFydzp5iCA6k9UGsdS9A', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (51, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYwMTE3MDU0LCJpYXQiOjE2NTk1MTIyNTR9.5bWBBUsn8GFzA3LwnItihvfXCUlkzNP8rtLX0XjTUI4', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (52, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYxMTQ5MTY4LCJpYXQiOjE2NjA1NDQzNjh9.egpnL-eB179MR8I2nK_LJV7_uG9P7qQeu8naj5TYcQk', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (53, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYxMTQ5NTM0LCJpYXQiOjE2NjA1NDQ3MzR9.btNa-5PHCxgMgI3uXnCAsuD3V_9c8kKfS1N31xg_1FE', 'test', '121.4.166.249', 0, 0, 0);
INSERT INTO `sa_access` VALUES (54, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlVzZXJOYW1lIjoiYWRtaW5AYmJidWcuY29tIiwiVXNlclBsYXQiOiJ0ZXN0IiwiZXhwIjoxNjYxMTUyNDMyLCJpYXQiOjE2NjA1NDc2MzJ9.tx0dZuIvagsycWT4ZmhMysjibrCjRSvg61nqCLki7N4', 'test', '192.168.18.76', 0, 0, 0);

-- ----------------------------
-- Table structure for sa_app
-- ----------------------------
DROP TABLE IF EXISTS `sa_app`;
CREATE TABLE `sa_app`  (
  `app_id` int NOT NULL AUTO_INCREMENT,
  `app_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'key',
  `app_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'name',
  `app_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'url',
  `app_user` int NOT NULL DEFAULT 0 COMMENT 'user',
  `app_scope` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'scope',
  `app_status` int NOT NULL DEFAULT 0 COMMENT '状态',
  `app_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `app_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`app_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1005 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '应用表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_app
-- ----------------------------
INSERT INTO `sa_app` VALUES (1, '请重置后对接', 'BBBUG', 'https://bbbug.com', 1, '', 0, 0, 0);
INSERT INTO `sa_app` VALUES (1001, '请重置后对接', 'Gitee', 'https://gitee.com/#extra#', 1, '', 0, 0, 0);
INSERT INTO `sa_app` VALUES (1002, '请重置后对接', 'OSChina', 'https://my.oschina.net/#extra#', 1, '', 0, 0, 0);
INSERT INTO `sa_app` VALUES (1003, '请重置后对接', 'QQ', 'https://hamm.cn', 1, '', 0, 0, 0);
INSERT INTO `sa_app` VALUES (1004, '请重置后对接', '钉钉', 'https://hamm.cn', 1, '', 0, 0, 0);
INSERT INTO `sa_app` VALUES (1005, '请重置后对接', '微信小程序', 'https://hamm.cn', 1, '', 0, 0, 0);

-- ----------------------------
-- Table structure for sa_attach
-- ----------------------------
DROP TABLE IF EXISTS `sa_attach`;
CREATE TABLE `sa_attach`  (
  `attach_id` int NOT NULL AUTO_INCREMENT,
  `attach_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '路径',
  `attach_used` int NOT NULL DEFAULT 0,
  `attach_thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `attach_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '类型',
  `attach_sha` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `attach_size` int NOT NULL DEFAULT 0 COMMENT '大小',
  `attach_user` int NOT NULL DEFAULT 0 COMMENT '用户',
  `attach_status` int NOT NULL DEFAULT 0 COMMENT '状态',
  `attach_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `attach_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`attach_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '附件表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_attach
-- ----------------------------
INSERT INTO `sa_attach` VALUES (1, '', 0, 'uploads/thumb/image/89386380c834351cd716ef0652a1a24d.png', 'png', '89386380c834351cd716ef0652a1a24d', 0, 0, 0, 1648611218, 1648611218);
INSERT INTO `sa_attach` VALUES (2, '75cabffb5ad0061352ed05eca0fbe1d8.png', 0, 'uploads/thumb/image/75cabffb5ad0061352ed05eca0fbe1d8.png', 'png', '75cabffb5ad0061352ed05eca0fbe1d8', 10977, 0, 0, 1648611303, 1648611303);
INSERT INTO `sa_attach` VALUES (3, '', 0, 'uploads/thumb/image/75cabffb5ad0061352ed05eca0fbe1d8.png', 'png', '75cabffb5ad0061352ed05eca0fbe1d8', 0, 0, 0, 1648611412, 1648611412);
INSERT INTO `sa_attach` VALUES (4, '75cabffb5ad0061352ed05eca0fbe1d8.png', 0, 'uploads/thumb/image/75cabffb5ad0061352ed05eca0fbe1d8.png', 'png', '75cabffb5ad0061352ed05eca0fbe1d8', 10977, 0, 0, 1648611549, 1648611549);
INSERT INTO `sa_attach` VALUES (5, '75cabffb5ad0061352ed05eca0fbe1d8.png', 0, 'uploads/thumb/image/75cabffb5ad0061352ed05eca0fbe1d8.png', 'png', '75cabffb5ad0061352ed05eca0fbe1d8', 10977, 2, 0, 1648619529, 1648619529);
INSERT INTO `sa_attach` VALUES (6, '75cabffb5ad0061352ed05eca0fbe1d8.png', 0, 'uploads/thumb/image/75cabffb5ad0061352ed05eca0fbe1d8.png', 'png', '75cabffb5ad0061352ed05eca0fbe1d8', 10977, 2, 0, 1648619633, 1648619633);
INSERT INTO `sa_attach` VALUES (7, '75cabffb5ad0061352ed05eca0fbe1d8.png', 0, 'uploads/thumb/image/75cabffb5ad0061352ed05eca0fbe1d8.png', 'png', '75cabffb5ad0061352ed05eca0fbe1d8', 10977, 2, 0, 1648619844, 1648619844);
INSERT INTO `sa_attach` VALUES (8, '75cabffb5ad0061352ed05eca0fbe1d8.png', 0, '', 'png', '', 66628, 2, 0, 1648620223, 1648620223);
INSERT INTO `sa_attach` VALUES (9, '75cabffb5ad0061352ed05eca0fbe1d8.png', 0, 'uploads/thumb/image/head/75cabffb5ad0061352ed05eca0fbe1d8.png', 'png', '75cabffb5ad0061352ed05eca0fbe1d8', 66628, 2, 0, 1648620306, 1648620306);
INSERT INTO `sa_attach` VALUES (10, '75cabffb5ad0061352ed05eca0fbe1d8.png', 0, 'uploads/thumb/image/head/75cabffb5ad0061352ed05eca0fbe1d8.png', 'png', '75cabffb5ad0061352ed05eca0fbe1d8', 66628, 2, 0, 1648627404, 1648627404);

-- ----------------------------
-- Table structure for sa_conf
-- ----------------------------
DROP TABLE IF EXISTS `sa_conf`;
CREATE TABLE `sa_conf`  (
  `conf_id` int NOT NULL AUTO_INCREMENT,
  `conf_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '参数名',
  `conf_value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '参数值',
  `conf_desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数描述',
  `conf_int` int NOT NULL DEFAULT 0 COMMENT '参数到期',
  `conf_status` int NOT NULL DEFAULT 0,
  `conf_createtime` int NOT NULL DEFAULT 0,
  `conf_updatetime` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`conf_id`) USING BTREE,
  INDEX `conf_key`(`conf_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 60 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_conf
-- ----------------------------
INSERT INTO `sa_conf` VALUES (11, 'weapp_appid', '', '小程序APPID', 0, 0, 0, 0);
INSERT INTO `sa_conf` VALUES (12, 'weapp_appkey', '', '小程序SECRET', 0, 0, 0, 0);
INSERT INTO `sa_conf` VALUES (39, 'upload_max_file', '2097152', '最大文件上传限制', 0, 0, 0, 0);
INSERT INTO `sa_conf` VALUES (40, 'upload_file_type', 'jpg,png,gif,jpeg,bmp,txt,pdf,mp3,mp4,amr,m4a,xls,xlsx,ppt,pptx,doc,docx', '允许文件上传类型', 0, 0, 0, 0);
INSERT INTO `sa_conf` VALUES (41, 'upload_max_image', '2097152', '最大图片上传限制', 0, 0, 0, 0);
INSERT INTO `sa_conf` VALUES (42, 'upload_image_type', 'jpg,png,gif,jpeg,bmp', '允许上传图片类型', 0, 0, 0, 0);
INSERT INTO `sa_conf` VALUES (47, 'default_group', '5', '注册默认用户组', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (48, 'email_account', 'admin@mail.bbbug.com', '邮箱账号', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (49, 'email_password', '123456', '邮箱密码', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (50, 'email_host', 'smtp.exmail.qq.com', '邮箱服务器', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (51, 'email_remark', 'BBBUG TEAM', '邮箱签名', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (52, 'email_port', '465', '邮箱端口号', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (53, 'websocket_http', 'http://127.0.0.1:10012/', 'WebsocketHTTP请求地址', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (54, 'websocket_token', 'wss_bbbug_com', 'Websocket验证串', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (55, 'api_guest_token', '45af3cfe44942c956e026d5fd58f0feffbd3a237', '临时用户access_token', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (56, 'frontend_url', '', '前端地址', 0, 0, 0, 0);
INSERT INTO `sa_conf` VALUES (57, 'api_url', '', 'API地址', 0, 0, 0, 0);
INSERT INTO `sa_conf` VALUES (58, 'tencent_ai_appid', '', '腾讯AI的APPID', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (59, 'tencent_ai_appkey', '', '腾讯AI的APPKEY', 0, 0, 0, 1598539052);
INSERT INTO `sa_conf` VALUES (60, 'static_url', '', 'Static文件地址', 0, 0, 0, 0);

-- ----------------------------
-- Table structure for sa_keywords
-- ----------------------------
DROP TABLE IF EXISTS `sa_keywords`;
CREATE TABLE `sa_keywords`  (
  `keywords_id` int NOT NULL AUTO_INCREMENT,
  `keywords_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '原串',
  `keywords_target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '替换',
  `keywords_all` int NOT NULL DEFAULT 0 COMMENT '全替换',
  `keywords_status` int NOT NULL DEFAULT 0 COMMENT '状态',
  `keywords_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `keywords_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`keywords_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '关键词表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_keywords
-- ----------------------------
INSERT INTO `sa_keywords` VALUES (3, '<script>', '我傻乎乎的想试试能不能XSS，结果被系统拦截了。。。', 1, 0, 1592574791, 1592575266);

-- ----------------------------
-- Table structure for sa_message
-- ----------------------------
DROP TABLE IF EXISTS `sa_message`;
CREATE TABLE `sa_message`  (
  `message_id` bigint NOT NULL AUTO_INCREMENT,
  `message_user` int NOT NULL DEFAULT 0 COMMENT 'user',
  `message_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'type',
  `message_where` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `message_to` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `message_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT 'content',
  `message_status` int NOT NULL DEFAULT 0 COMMENT '状态',
  `message_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `message_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`message_id`) USING BTREE,
  INDEX `message_user`(`message_user`) USING BTREE,
  INDEX `message_createtime`(`message_createtime`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 79 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '消息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_message
-- ----------------------------
INSERT INTO `sa_message` VALUES (1, 1, 'text', 'channel', '888', 'test', 1, 0, 0);
INSERT INTO `sa_message` VALUES (2, 200, 'text', 'channel', '888', '%E6%81%B6%E5%BF%83', 1, 0, 0);
INSERT INTO `sa_message` VALUES (3, 200, 'text', 'channel', '888', '%F0%9F%91%81%F0%9F%91%81', 1, 0, 0);
INSERT INTO `sa_message` VALUES (4, 1, 'text', 'channel', '888', 'shuonimane', 1, 0, 0);
INSERT INTO `sa_message` VALUES (5, 200, 'text', 'channel', '888', '%E9%82%A3%E6%B2%A1%E4%BA%8B%E4%BA%86', 1, 0, 0);
INSERT INTO `sa_message` VALUES (6, 1, 'text', 'channel', '888', '%E6%80%8E%E4%B9%88%E6%92%AD%E6%94%BE%E5%91%A2', 1, 0, 0);
INSERT INTO `sa_message` VALUES (7, 1, 'text', 'channel', '888', '%E6%88%91%E9%9D%A0', 1, 0, 0);
INSERT INTO `sa_message` VALUES (8, 1, 'text', 'channel', '888', '%40', 1, 0, 0);
INSERT INTO `sa_message` VALUES (9, 1, 'text', 'channel', '888', 'c', 1, 0, 0);
INSERT INTO `sa_message` VALUES (10, 1, 'text', 'channel', '888', '111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (11, 1, 'text', 'channel', '888', '%40%E6%9C%BA%E5%99%A8%E4%BA%BA%20111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (12, 1, 'text', 'channel', '888', '%40%E6%9C%BA%E5%99%A8%E4%BA%BA%20111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (13, 1, 'text', 'channel', '888', '%40%E6%9C%BA%E5%99%A8%E4%BA%BA%20111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (14, 1, 'text', 'channel', '888', '%40%E6%9C%BA%E5%99%A8%E4%BA%BA%20111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (15, 1, 'text', 'channel', '888', '1', 1, 0, 0);
INSERT INTO `sa_message` VALUES (16, 1, 'text', 'channel', '888', '111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (17, 1, 'text', 'channel', '888', '%40%E6%9C%BA%E5%99%A8%E4%BA%BA%20111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (18, 1, 'text', 'channel', '888', '%40%E6%9C%BA%E5%99%A8%E4%BA%BA%20111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (19, 1, 'text', 'channel', '888', '111', 1, 0, 0);
INSERT INTO `sa_message` VALUES (20, 200, 'text', 'channel', '888', 'test', 1, 0, 0);
INSERT INTO `sa_message` VALUES (21, 1, 'text', 'channel', '888', 'baby%20', 1, 0, 0);
INSERT INTO `sa_message` VALUES (22, 1, 'text', 'channel', '888', '%E5%8D%95%E6%9B%B2%E5%BE%AA%E7%8E%AF%E4%BA%86', 1, 0, 0);
INSERT INTO `sa_message` VALUES (23, 1, 'text', 'channel', '888', '%E6%88%91%E6%B2%A1%E8%AE%BE%E7%BD%AE%E5%95%8A', 1, 0, 0);
INSERT INTO `sa_message` VALUES (24, 1, 'text', 'channel', '888', '%E8%87%AA%E5%B7%B1%E9%80%89', 1, 0, 0);
INSERT INTO `sa_message` VALUES (25, 1, 'text', 'channel', '888', '%E6%9D%A5%E9%A6%96%E6%91%87%E6%BB%9A', 1, 0, 0);
INSERT INTO `sa_message` VALUES (26, 1, 'text', 'channel', '888', '%E5%A5%BD%E6%94%BE%EF%BC%81', 1, 0, 0);
INSERT INTO `sa_message` VALUES (27, 1, 'text', 'channel', '888', '%E6%80%8E%E4%B9%88%E6%98%AF%E7%8E%B0%E5%9C%BA%E7%89%88%E7%9A%84', 1, 0, 0);
INSERT INTO `sa_message` VALUES (28, 1, 'text', 'channel', '888', '%E9%9A%8F%E6%9C%BA%E6%90%9C%E7%9A%84', 1, 0, 0);
INSERT INTO `sa_message` VALUES (29, 1, 'text', 'channel', '888', '%E4%BD%A0%E6%80%8E%E4%B9%88%E5%90%AC%E5%87%BA%E6%9D%A5%E7%9A%84', 1, 0, 0);
INSERT INTO `sa_message` VALUES (30, 1, 'text', 'channel', '888', '%E6%9C%89%E7%8E%B0%E5%9C%BA%E8%A7%82%E4%BC%97%E7%9A%84%E5%B0%96%E5%8F%AB%E5%A3%B0', 1, 0, 0);
INSERT INTO `sa_message` VALUES (31, 1, 'text', 'channel', '888', '%E6%80%8E%E4%B9%88%E5%8F%88%E8%B7%B3%E5%88%B0%E8%BF%99%E9%A6%96%E4%BA%86%E3%80%82%E3%80%82%E3%80%82', 1, 0, 0);
INSERT INTO `sa_message` VALUES (32, 1, 'text', 'channel', '888', '%E5%93%88%E5%93%88', 1, 0, 0);
INSERT INTO `sa_message` VALUES (33, 1, 'text', 'channel', '888', '%E4%BD%A0%E8%83%BD%E4%B8%8D%E8%83%BD%E6%8D%A2%E4%B8%AA%E5%8F%B7', 1, 0, 0);
INSERT INTO `sa_message` VALUES (34, 1, 'text', 'channel', '888', '%E5%A4%A9%E5%A4%A9%E7%94%A8%E6%88%91%E7%9A%84%E7%AE%A1%E7%90%86%E5%91%98%E8%B4%A6%E6%88%B7', 1, 0, 0);
INSERT INTO `sa_message` VALUES (35, 1, 'text', 'channel', '888', '%20%E6%88%91%E8%A6%81%E9%A9%AC%E7%94%B2', 1, 0, 0);
INSERT INTO `sa_message` VALUES (36, 200, 'text', 'channel', '888', 'didi', 1, 0, 0);
INSERT INTO `sa_message` VALUES (37, 200, 'text', 'channel', '888', 'rnm%EF%BC%8C%E5%AD%A4%E5%8B%87%E8%80%85%E9%83%BD%E7%82%B9%E4%B8%8D%E4%BA%86', 1, 0, 0);
INSERT INTO `sa_message` VALUES (38, 1, 'text', 'channel', '888', '%E4%B8%8D%E6%98%AF%E5%90%A7', 1, 0, 0);
INSERT INTO `sa_message` VALUES (39, 200, 'text', 'channel', '888', '%F0%9F%92%83%F0%9F%8F%BB%F0%9F%92%83%F0%9F%8F%BB%F0%9F%92%83%F0%9F%8F%BB', 1, 0, 0);
INSERT INTO `sa_message` VALUES (40, 1, 'text', 'channel', '888', '%E5%A5%BD%E5%8F%91', 1, 0, 0);
INSERT INTO `sa_message` VALUES (41, 200, 'text', 'channel', '888', '%E6%94%B6%E5%91%B3%E5%84%BF', 1, 0, 0);
INSERT INTO `sa_message` VALUES (42, 1, 'text', 'channel', '888', '%EF%BC%9F', 1, 0, 0);
INSERT INTO `sa_message` VALUES (43, 200, 'text', 'channel', '888', '%E5%B0%91%E5%A5%B3%E6%97%B6%E4%BB%A3%E7%9A%84%E6%AD%8C%E4%B9%9F%E6%94%BE%E4%B8%8D%E4%BA%86%E5%90%97%EF%BC%9F', 1, 0, 0);
INSERT INTO `sa_message` VALUES (44, 1, 'text', 'channel', '888', '%E4%B8%8D%E5%BA%94%E8%AF%A5', 1, 0, 0);
INSERT INTO `sa_message` VALUES (45, 1, 'text', 'channel', '888', '%E5%88%ABdarling%E4%BA%86', 1, 0, 0);
INSERT INTO `sa_message` VALUES (46, 200, 'text', 'channel', '888', '%E6%88%91%E6%B2%A1%E5%8A%A8', 1, 0, 0);
INSERT INTO `sa_message` VALUES (47, 1, 'text', 'channel', '888', '%E6%88%91%E7%9F%A5%E9%81%93', 1, 0, 0);
INSERT INTO `sa_message` VALUES (48, 200, 'text', 'channel', '888', '%E4%B8%80%E7%9B%B4%E5%8D%95%E6%9B%B2', 1, 0, 0);
INSERT INTO `sa_message` VALUES (49, 200, 'text', 'channel', '888', '%E6%9D%A5%E9%A6%96GENIE', 1, 0, 0);
INSERT INTO `sa_message` VALUES (50, 200, 'text', 'channel', '888', '%E6%88%91%E8%BF%99%E9%87%8C%E7%BB%9D%E5%AF%B9%E6%9C%89%E9%97%AE%E9%A2%98', 1, 0, 0);
INSERT INTO `sa_message` VALUES (51, 200, 'text', 'channel', '888', '%E6%B0%B8%E8%BF%9C%E6%92%AD%E6%94%BE%E5%87%BA%E9%94%99', 1, 0, 0);
INSERT INTO `sa_message` VALUES (52, 200, 'text', 'channel', '888', '%E4%BD%A0%E8%AF%B4%E4%B8%8D%E5%AE%9A%E6%9D%A5linkin%20park%E4%B9%9F%E8%83%BD%E6%94%BE', 1, 0, 0);
INSERT INTO `sa_message` VALUES (53, 1, 'text', 'channel', '888', '%E4%B8%80%E4%B8%AAapi', 1, 0, 0);
INSERT INTO `sa_message` VALUES (54, 1, 'text', 'channel', '888', '%E8%82%AF%E5%AE%9A%E6%98%AF%E7%BD%91%E7%BB%9C%E6%B3%A2%E5%8A%A8', 1, 0, 0);
INSERT INTO `sa_message` VALUES (55, 1, 'text', 'channel', '888', 'wc', 1, 0, 0);
INSERT INTO `sa_message` VALUES (56, 200, 'text', 'channel', '888', 'hing%E8%8A%AF', 1, 0, 0);
INSERT INTO `sa_message` VALUES (57, 200, 'text', 'channel', '888', '%E3%80%82%E3%80%82%E3%80%82', 1, 0, 0);
INSERT INTO `sa_message` VALUES (58, 200, 'text', 'channel', '888', '%E3%80%82%E3%80%82%E3%80%82', 1, 0, 0);
INSERT INTO `sa_message` VALUES (59, 1, 'text', 'channel', '888', '%3F', 1, 0, 0);
INSERT INTO `sa_message` VALUES (60, 1, 'text', 'channel', '888', '%E5%A4%9A%E5%A5%BD%E5%90%AC', 1, 0, 0);
INSERT INTO `sa_message` VALUES (61, 1, 'text', 'channel', '888', '%E5%A5%BD%E6%94%BE', 1, 0, 0);
INSERT INTO `sa_message` VALUES (62, 1, 'text', 'channel', '888', '%E5%A5%BD%E6%94%BE', 1, 0, 0);
INSERT INTO `sa_message` VALUES (63, 1, 'text', 'channel', '888', '%E5%A4%AA%E6%B6%A9%E4%BA%86', 1, 0, 0);
INSERT INTO `sa_message` VALUES (64, 1, 'text', 'channel', '888', '%E9%9B%B6%E9%9B%B6%E5%AD%90%EF%BC%8C%E5%BE%88%E6%AD%A3%E5%B8%B8', 1, 0, 0);
INSERT INTO `sa_message` VALUES (65, 1, 'text', 'channel', '888', '%E5%A4%AA%E9%9A%BE%E9%A1%B6%E4%BA%86', 1, 0, 0);
INSERT INTO `sa_message` VALUES (66, 1, 'text', 'channel', '888', '%E8%BA%AB%E4%B8%8A%E6%9C%89%E8%9A%82%E8%9A%81%E7%88%AC', 1, 0, 0);
INSERT INTO `sa_message` VALUES (67, 1, 'text', 'channel', '888', '%E5%8E%9F%E6%9D%A5%E6%98%AF%E8%BF%99%E9%A6%96%E6%AD%8C', 1, 0, 0);
INSERT INTO `sa_message` VALUES (68, 1, 'text', 'channel', '888', '%E5%99%9C%E5%95%A6%E5%99%9C%E5%95%A6%E5%98%9E%E7%BB%BF%E7%BB%BF%E7%BB%BF%E7%BB%BF%E7%BB%BF%E7%BB%BF%E7%BB%BF', 1, 0, 0);
INSERT INTO `sa_message` VALUES (69, 1, 'text', 'channel', '888', '%E5%99%9C%E5%95%A6%E5%99%9C%E5%95%A6%E5%98%9E', 1, 0, 0);
INSERT INTO `sa_message` VALUES (70, 1, 'text', 'channel', '888', '%E8%BF%99%E6%80%8E%E4%B9%88%E6%B2%A1%E5%90%AC%E8%BF%87', 1, 0, 0);
INSERT INTO `sa_message` VALUES (71, 1, 'text', 'channel', '888', '%E5%A5%BD%EF%BC%81', 1, 0, 0);
INSERT INTO `sa_message` VALUES (72, 1, 'text', 'channel', '888', '%E4%BD%A0%E6%94%BE%E7%9A%84%E5%A5%BD%E5%95%8A%E4%BD%A0%E6%94%BE%E7%9A%84%E5%A5%BD', 1, 0, 0);
INSERT INTO `sa_message` VALUES (73, 1, 'text', 'channel', '888', '%E5%85%B8', 1, 0, 0);
INSERT INTO `sa_message` VALUES (74, 1, 'text', 'channel', '888', '%E5%A4%AA%E6%9C%89%E5%91%B3%E5%84%BF%E4%BA%86', 1, 0, 0);
INSERT INTO `sa_message` VALUES (75, 1, 'text', 'channel', '888', '%E6%80%8E%E4%B9%88%E6%B2%A1%E5%A3%B0%EF%BC%9F', 1, 0, 0);
INSERT INTO `sa_message` VALUES (76, 1, 'text', 'channel', '888', '%E6%B2%A1%E5%88%B0', 1, 0, 0);
INSERT INTO `sa_message` VALUES (77, 1, 'text', 'channel', '888', '%E5%8F%AE%E5%92%9A', 1, 0, 0);
INSERT INTO `sa_message` VALUES (78, 1, 'text', 'channel', '888', '%E6%80%80%E7%96%91%E6%98%AF%E6%88%91%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%9A%84%E5%8E%9F%E5%9B%A0', 1, 0, 0);
INSERT INTO `sa_message` VALUES (79, 1, 'text', 'channel', '888', 'test', 1, 0, 0);

-- ----------------------------
-- Table structure for sa_room
-- ----------------------------
DROP TABLE IF EXISTS `sa_room`;
CREATE TABLE `sa_room`  (
  `room_id` int NOT NULL AUTO_INCREMENT,
  `room_user` int NOT NULL DEFAULT 0 COMMENT '所有者ID',
  `room_addsongcd` int NOT NULL DEFAULT 60 COMMENT '点歌CD',
  `room_addcount` int NOT NULL DEFAULT 5 COMMENT '点歌数量',
  `room_pushdaycount` int NOT NULL DEFAULT 5 COMMENT '顶歌日限额',
  `room_pushsongcd` int NOT NULL DEFAULT 3600 COMMENT '顶歌CD',
  `room_online` int NOT NULL DEFAULT 0 COMMENT '已登录在线',
  `room_realonline` int NOT NULL DEFAULT 0 COMMENT '所有在线',
  `room_hide` int NOT NULL DEFAULT 0 COMMENT '是否从列表隐藏',
  `room_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '房间名称',
  `room_type` int NOT NULL DEFAULT 1 COMMENT '房间类型',
  `room_public` int NOT NULL DEFAULT 0,
  `room_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '房间密码',
  `room_notice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '进入房间提醒',
  `room_addsong` int NOT NULL DEFAULT 0,
  `room_sendmsg` int NOT NULL DEFAULT 0,
  `room_robot` int NOT NULL DEFAULT 0,
  `room_order` int NOT NULL DEFAULT 0,
  `room_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `room_playone` int NOT NULL DEFAULT 0 COMMENT '0随机1单曲',
  `room_votepass` int NOT NULL DEFAULT 1,
  `room_votepercent` int NOT NULL DEFAULT 30,
  `room_background` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '房间背景图',
  `room_app` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '插件地址',
  `room_fullpage` int NOT NULL DEFAULT 0 COMMENT '插件是否全屏',
  `room_status` int NOT NULL DEFAULT 0 COMMENT '状态',
  `room_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `room_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`room_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1889 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '房间表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_room
-- ----------------------------
INSERT INTO `sa_room` VALUES (888, 1, 60, 5, 5, 3600, 0, 2, 0, 'BBBUG音乐大厅', 1, 0, '', '大厅为电台播放模式，欢迎大家点歌，房间已支持自定义点歌/顶歌等CD和数量，快去房间管理页面看看吧~', 1, 0, 0, 10000000, '', 1, 1, 30, '2', '', 0, 0, 1598539777, 1648883680);
INSERT INTO `sa_room` VALUES (1888, 200, 60, 5, 5, 3600, 1, 2, 0, 'BBBUG音乐大厅', 1, 0, '1111', '大厅为电台播放模式，欢迎大家点歌，房间已支持自定义点歌/顶歌等CD和数量，快去房间管理页面看看吧~', 0, 0, 0, 10000000, '', 0, 1, 30, '', '', 0, 0, 1598539777, 1648888144);

-- ----------------------------
-- Table structure for sa_song
-- ----------------------------
DROP TABLE IF EXISTS `sa_song`;
CREATE TABLE `sa_song`  (
  `song_id` int NOT NULL AUTO_INCREMENT,
  `song_user` int NOT NULL DEFAULT 0,
  `song_mid` bigint NOT NULL DEFAULT 0,
  `song_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '歌曲名称',
  `song_singer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '歌手',
  `song_pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `song_length` int NOT NULL DEFAULT 0,
  `song_play` int NOT NULL DEFAULT 1 COMMENT '被点次数',
  `song_week` int NOT NULL DEFAULT 0 COMMENT '本周被点次数',
  `song_fav` int NOT NULL DEFAULT 0 COMMENT '0点歌 1收藏',
  `song_status` int NOT NULL DEFAULT 0 COMMENT '状态',
  `song_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `song_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`song_id`) USING BTREE,
  INDEX `song_mid`(`song_mid`) USING BTREE,
  INDEX `song_user`(`song_user`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 92 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '歌曲表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_song
-- ----------------------------
INSERT INTO `sa_song` VALUES (1, 1, 23655613, '单身情歌', '林志炫', 'https://img4.kuwo.cn/star/albumcover/120/57/55/276338272.jpg', 275, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (2, 1, 180004855, '城市傍晚', '毛不易', 'https://img4.kuwo.cn/star/albumcover/120/40/35/3993015688.jpg', 251, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (3, 1, 274791, '挪威的森林', '伍佰&nbsp;And&nbsp;China&nbsp;Blue', 'https://img3.kuwo.cn/star/albumcover/500/77/7/973071322.jpg', 391, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (4, 1, 28618517, 'Dusk&nbsp;Till&nbsp;Dawn', 'ZAYN&Sia', 'https://img4.kuwo.cn/star/albumcover/120/61/11/2392722128.jpg', 265, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (5, 1, 55827880, 'We&nbsp;Will&nbsp;Rock&nbsp;You-《波西米亚狂想曲》电影插曲', 'Queen', 'https://img4.kuwo.cn/star/albumcover/120/89/38/158513320.jpg', 122, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (6, 1, 55827893, 'Bohemian&nbsp;Rhapsody-《波西米亚狂想曲》电影插曲', 'Queen', 'https://img4.kuwo.cn/star/albumcover/120/89/38/158513320.jpg', 390, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (7, 1, 3493772, 'Bohemian&nbsp;Rhapsody', 'Queen', 'https://img4.kuwo.cn/star/albumcover/120/74/40/3759290886.jpg', 355, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (8, 1, 56514931, '霓虹甜心', '马赛克', 'https://img4.kuwo.cn/star/albumcover/120/20/52/3746506517.jpg', 271, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (9, 1, 68465355, '夏日漱石&nbsp;(Summer&nbsp;Cozy&nbsp;Rock)', '橘子海', 'https://img4.kuwo.cn/star/albumcover/120/12/53/1509423259.jpg', 263, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (10, 1, 240306, 'Innocence', 'Avril&nbsp;Lavigne', 'https://img4.kuwo.cn/star/albumcover/120/93/84/101718004.jpg', 232, 3, 2, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (11, 1, 69115560, 'Everybody&nbsp;Hurts', 'Avril&nbsp;Lavigne', 'https://img4.kuwo.cn/star/albumcover/120/49/50/3301596756.jpg', 221, 2, 1, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (12, 200, 240306, 'Innocence', 'Avril&nbsp;Lavigne', 'https://img4.kuwo.cn/star/albumcover/120/93/84/101718004.jpg', 232, 2, 1, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (13, 200, 69115560, 'Everybody&nbsp;Hurts', 'Avril&nbsp;Lavigne', 'https://img4.kuwo.cn/star/albumcover/120/49/50/3301596756.jpg', 221, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (14, 200, 23655613, '单身情歌', '林志炫', 'https://img4.kuwo.cn/star/albumcover/120/57/55/276338272.jpg', 275, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (15, 1, 205517490, '孤勇者', '陈奕迅', 'https://img4.kuwo.cn/star/starheads/500/23/6/1897539410.jpg', 62, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (16, 1, 1248953, 'DAY&nbsp;BY&nbsp;DAY', 'T-ara', 'https://img4.kuwo.cn/star/albumcover/120/19/62/1443364194.jpg', 209, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (17, 200, 1248953, 'DAY&nbsp;BY&nbsp;DAY', 'T-ara', 'https://img4.kuwo.cn/star/albumcover/120/19/62/1443364194.jpg', 209, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (18, 1, 3437085, 'Sexy&nbsp;Love&nbsp;(Japanese&nbsp;ver.)', 'T-ara', 'https://img4.kuwo.cn/star/albumcover/120/86/53/2139831889.jpg', 226, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (19, 1, 543883, 'MY&nbsp;ALL', '浜崎あゆみ', 'https://img4.kuwo.cn/star/albumcover/120/11/50/2019718377.jpg', 326, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (20, 1, 40718304, 'GENIE', '少女时代', 'https://img4.kuwo.cn/star/albumcover/120/63/2/2088671823.jpg', 224, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (21, 1, 28409250, 'Numb-英雄联盟代表音乐', 'Linkin&nbsp;Park', 'https://img4.kuwo.cn/star/albumcover/120/24/24/2693959198.jpg', 187, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (22, 200, 28409250, 'Numb-英雄联盟代表音乐', 'Linkin&nbsp;Park', 'https://img4.kuwo.cn/star/albumcover/120/24/24/2693959198.jpg', 187, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (23, 1, 28409528, 'In&nbsp;The&nbsp;End-《QQ飞车手游》游戏背景歌曲', 'Linkin&nbsp;Park', 'https://img4.kuwo.cn/star/albumcover/120/45/36/1535239246.jpg', 216, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (24, 200, 28409528, 'In&nbsp;The&nbsp;End-《QQ飞车手游》游戏背景歌曲', 'Linkin&nbsp;Park', 'https://img4.kuwo.cn/star/albumcover/120/45/36/1535239246.jpg', 216, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (25, 1, 144527321, 'ご唱和ください&nbsp;我の名を！-《泽塔奥特曼》特摄剧主题曲', '遠藤正明', 'https://img1.kuwo.cn/star/starheads/120/26/1/2603812966.jpg', 92, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (26, 1, 1553982, 'Butterfly', 'Smile.DK', 'https://img4.kuwo.cn/star/albumcover/120/36/83/2315009878.jpg', 180, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (27, 1, 5254550, 'Butter-Fly&nbsp;(劇場サイズ#3)-《数码宝贝》TV动画片头曲', '和田光司', 'https://img1.kuwo.cn/star/starheads/120/28/18/2894513305.jpg', 67, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (28, 1, 164400763, '天地劫之幽城再临', '呦猫UNEKO', 'https://img4.kuwo.cn/star/albumcover/120/23/31/2556067907.jpg', 239, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (29, 1, 169036108, '天地劫', '黄龄', 'https://img4.kuwo.cn/star/albumcover/120/65/94/2561905124.jpg', 162, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (30, 1, 185361801, '问劫-《天地劫》手游主题曲', '徐佳莹', 'https://img4.kuwo.cn/star/albumcover/120/35/90/1861650493.jpg', 210, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (31, 1, 202367109, '叹', '黄龄&Tăng&nbsp;Duy&nbsp;Tân', 'https://img4.kuwo.cn/star/albumcover/120/77/22/3463145793.jpg', 251, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (32, 1, 153518681, '牵丝戏（梦幻西游盘丝洞门派曲）', '黄龄', 'https://img4.kuwo.cn/star/albumcover/120/45/50/2044843992.jpg', 251, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (33, 1, 149499336, '无价之姐', '李宇春&郑希怡&郁可唯&蓝盈莹&万茜&黄龄&张含韵&张萌&袁咏琳&宁静&金晨&伊能静&孟佳&王霏霏&张雨绮&李斯丹妮&陈松伶&吴昕&沈梦辰&金莎&王智&阿朵&钟丽缇&黄圣依&海陆&刘芸&白冰&王丽坤&许飞', 'https://img4.kuwo.cn/star/albumcover/120/86/69/2134354566.jpg', 329, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (34, 1, 807748, 'High歌', '黄龄', 'https://img4.kuwo.cn/star/albumcover/120/92/84/657730923.jpg', 283, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (35, 1, 228258581, '星河叹-《星汉灿烂》电视剧女主人物曲', '黄龄', 'https://img4.kuwo.cn/star/albumcover/120/3/82/730524348.jpg', 264, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (36, 1, 4856712, '惊鸿一面', '许嵩&黄龄', 'https://img3.kuwo.cn/star/albumcover/500/59/23/2851294451.jpg', 256, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (37, 1, 2920161, '伤不起', '王麟&何鹏', 'https://img4.kuwo.cn/star/albumcover/120/80/43/882756362.jpg', 328, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (38, 1, 67307403, '只因你太美', 'SWIN-S', 'https://img1.kuwo.cn/star/starheads/120/39/16/1810373762.jpg', 219, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (39, 1, 218607362, 'M八七', '米津玄师', 'https://img4.kuwo.cn/star/albumcover/120/46/57/3885765967.jpg', 263, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (40, 1, 146386177, 'Lemon', '米津玄师', 'https://img4.kuwo.cn/star/albumcover/120/14/69/1739667594.jpg', 255, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (41, 1, 2254622, 'Jupiter', '平原綾香', 'https://img4.kuwo.cn/star/albumcover/120/15/27/860127043.jpg', 361, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (42, 1, 703564, 'Time&nbsp;After&nbsp;Time', '倉木麻衣', 'https://img4.kuwo.cn/star/albumcover/120/54/21/3316802890.jpg', 246, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (43, 1, 85814691, '渡月橋&nbsp;～君&nbsp;想ふ～', '倉木麻衣', 'https://img4.kuwo.cn/star/albumcover/120/26/53/311959806.jpg', 246, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (44, 1, 1651004, '月半小夜曲', '李克勤', 'https://img4.kuwo.cn/star/albumcover/120/82/52/4095053955.jpg', 270, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (45, 1, 20490865, '夜半小夜曲', '李克群、陈苑淇', 'https://img1.kuwo.cn/star/starheads/', 291, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (46, 1, 28483813, '一生不变(Album&nbsp;Version)', '李克勤', 'https://img1.kuwo.cn/star/starheads/120/18/31/3491428260.jpg', 260, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (47, 1, 98316686, '红日-《他来自天堂》电视剧主题曲', '李克勤', 'https://img4.kuwo.cn/star/albumcover/120/76/27/1148976275.jpg', 291, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (48, 1, 24582262, '夜猫', '张蔷', 'https://img4.kuwo.cn/star/albumcover/120/71/15/1114090424.jpg', 247, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (49, 1, 840198, '猪猪侠-《猪猪侠》动漫主题曲', '陈洁丽', 'https://img4.kuwo.cn/star/albumcover/120/96/12/4065572910.jpg', 211, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (50, 1, 1077770, '超兽武装-《超兽武装之仁者无敌》动画片主题曲', '陈洁丽&刘罡', 'https://img4.kuwo.cn/star/albumcover/120/31/55/2175429781.jpg', 226, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (51, 1, 500963, '爱在心中-《东方神娃》第二部动画主题曲', '孙晔', 'https://img4.kuwo.cn/star/albumcover/120/44/33/2941449854.jpg', 122, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (52, 1, 443607, '噢买尬', '五月天', 'https://img4.kuwo.cn/star/albumcover/120/76/73/545763767.jpg', 176, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (53, 1, 547911, '人生不过是一百年-《虹猫蓝兔七侠传》片头曲', '影视原声', 'https://img4.kuwo.cn/star/albumcover/120/8/1/3087985440.jpg', 99, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (54, 1, 51597366, 'GBL女神殿', 'dnf', 'https://img1.kuwo.cn/star/starheads/120/80/90/2969467752.jpg', 112, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (55, 1, 4958680, '我为厨艺狂-《神厨小福贵》动画片尾曲', '群星', 'https://img1.kuwo.cn/star/starheads/120/41/73/4157743756.jpg', 86, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (56, 1, 547872, '豪情走天涯-《神厨小福贵》动画片头曲', '影视原声', 'https://img1.kuwo.cn/star/starheads/120/63/51/3109218617.jpg', 93, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (57, 1, 534239, '爱不会绝迹-《恐龙宝贝》动画片主题曲', '林俊杰', 'https://img4.kuwo.cn/star/albumcover/120/84/41/4159059729.jpg', 240, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (58, 1, 696528, '大角牛之歌-《大角牛》动画片主题曲', '沸点乐队', 'https://img4.kuwo.cn/star/albumcover/120/92/40/917707347.jpg', 109, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (59, 1, 122701, '梦的光点-《神兵小将》动画片片头曲', '王心凌', 'https://img4.kuwo.cn/star/albumcover/120/77/28/1947667828.jpg', 231, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (60, 1, 21234429, '青春之火-《火力少年王》动漫主题歌', '潘松和', 'https://img4.kuwo.cn/star/albumcover/120/78/29/3705082739.jpg', 160, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (61, 1, 23642751, '月亮船', '杨钰莹', 'https://img4.kuwo.cn/star/albumcover/120/20/22/3289005013.jpg', 323, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (62, 1, 666200, '月亮船-《快乐星球》电视剧片尾曲', '少儿天唱组合', 'https://img1.kuwo.cn/star/starheads/120/94/51/1611200171.jpg', 240, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (63, 1, 51383903, '别看我只是一只羊-《喜羊羊与灰太狼》动画主题曲|《喜羊羊与灰太狼之开心方程式》动画主题曲|《喜羊羊与灰太狼之深海历险》动画片尾曲', '古倩敏&杨沛宜', 'https://img4.kuwo.cn/star/albumcover/120/27/0/1164104754.jpg', 168, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (64, 1, 281762, '离开地球表面-《开心超人》电影片尾曲', '五月天', 'https://img4.kuwo.cn/star/albumcover/120/2/60/3827709536.jpg', 275, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (65, 1, 6124411, '永远的奥特曼-《迪迦奥特曼》动画片尾曲', '电视原声', 'https://img4.kuwo.cn/star/albumcover/120/77/66/239870060.jpg', 197, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (66, 1, 5961360, '月光-《秦时明月》动画片头曲|《秦时明月之君临天下》动画片头曲', '胡彦斌', 'https://img4.kuwo.cn/star/albumcover/120/59/77/2376077807.jpg', 272, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (67, 1, 497264, '风云决-《风云决》动画电影插曲', '任贤齐', 'https://img4.kuwo.cn/star/albumcover/120/37/60/552902605.jpg', 287, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (68, 1, 6903639, '开心往前飞-《开心宝贝》动画片主题曲|《开心超人》电影片头曲', 'ViVi', 'https://img1.kuwo.cn/star/starheads/55/77/11/2458662407.jpg', 78, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (69, 1, 22865865, '十二生肖闯江湖-《十二生肖闯江湖》动画片主题曲', '十二生肖闯江湖', 'https://img1.kuwo.cn/star/starheads/120/45/41/4071809482.jpg', 87, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (70, 1, 1108363, '难念的经-《天龙八部》电视剧主题曲', '周华健', 'https://img1.kuwo.cn/star/albumcover/500/18/34/348086044.jpg', 288, 8, 7, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (71, 1, 41196396, '刀剑如梦', '周华健', 'https://img2.kuwo.cn/star/albumcover/500/51/43/4239375159.jpg', 202, 6, 5, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (72, 1, 38513799, '神话情话-《天下有情人》粤语版|《神雕侠侣》电视剧主题曲', '周华健&齐豫', 'http://img3.kuwo.cn/star/albumcover/500/44/76/2101242625.jpg', 266, 11, 10, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (73, 1, 7075846, '天下有情人', '周华健&齐豫', 'https://img4.kuwo.cn/star/albumcover/500/62/63/4165334067.jpg', 172, 7, 6, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (74, 1, 67821751, '知否知否', '齐豫&胡夏', 'https://img3.kuwo.cn/star/albumcover/500/31/16/1794259845.jpg', 290, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (75, 1, 5496028, '倾城', '陈奕迅', 'https://img1.kuwo.cn/star/starheads/120/23/6/1897539410.jpg', 249, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (76, 1, 198554068, '孤勇者-《英雄联盟：双城之战》动画剧集中文主题曲', '陈奕迅', 'https://img3.kuwo.cn/star/albumcover/500/49/79/1011385325.jpg', 256, 4, 3, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (77, 1, 317730, '富士山下-《爱情转移》粤语版', '陈奕迅', 'https://img3.kuwo.cn/star/albumcover/500/61/8/208529230.jpg', 259, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (78, 1, 204830470, '她会魔法吧', 'DJ小鱼儿', 'https://img1.kuwo.cn/star/albumcover/500/92/86/1884598203.jpg', 180, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (79, 1, 38548792, '爱情错觉', '王娅', 'https://img4.kuwo.cn/star/albumcover/500/91/3/1560096476.jpg', 244, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (80, 1, 80958029, '我要找到你', '小阿枫', 'https://img3.kuwo.cn/star/albumcover/500/85/4/1207575179.jpg', 157, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (81, 1, 92443539, '达拉崩吧', '周深', 'https://img3.kuwo.cn/star/albumcover/500/22/29/754779525.jpg', 245, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (82, 1, 7139252, '九九八十一', '双笙&nbsp;(陈元汐)&南久&易言&樊棋', 'http://img3.kuwo.cn/star/albumcover/500/64/43/1977085465.jpg', 287, 17, 16, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (83, 1, 23856230, '锦鲤抄-天刀2016秋季版本外装同名主题', '银临', 'https://img2.kuwo.cn/star/albumcover/500/86/67/3731560114.jpg', 245, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (84, 1, 6304356, '牵丝戏', '银临&Aki阿杰', 'https://img3.kuwo.cn/star/albumcover/500/95/24/2202064696.jpg', 239, 15, 14, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (85, 1, 6784156, '明月天涯', '五音Jw', 'https://img2.kuwo.cn/star/albumcover/500/71/81/917902547.jpg', 241, 2, 1, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (86, 1, 5237999, '棠梨煎雪', '银临', 'https://img4.kuwo.cn/star/albumcover/500/13/96/2012123181.jpg', 245, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (87, 1, 512842, '生生世世爱-《仙剑奇侠传》电视剧插曲', '吴雨霏', 'https://img3.kuwo.cn/star/albumcover/500/74/48/4168686413.jpg', 281, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (88, 1, 159679426, '兰亭序', '周杰伦', 'https://img1.kuwo.cn/star/starheads/500/8/10/2150960774.jpg', 234, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (89, 1, 206515968, '奢香夫人&nbsp;(2021时光音乐会第3期现场)', '林志炫', 'https://img2.kuwo.cn/star/starheads/500/73/59/1217516097.jpg', 255, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (90, 1, 922516, '凤凰花开的路口', '林志炫', 'https://img1.kuwo.cn/star/albumcover/500/50/23/700436519.jpg', 251, 1, 0, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (91, 1, 3248366, '烟花易冷-原唱：周杰伦', '林志炫', 'http://img2.kuwo.cn/star/albumcover/500/49/51/837580951.jpg', 316, 6, 5, 0, 0, 0, 0);
INSERT INTO `sa_song` VALUES (92, 1, 199423299, '天涯', '林志炫', 'http://img3.kuwo.cn/star/albumcover/500/47/20/2043152394.jpg', 279, 6, 5, 0, 0, 0, 0);

-- ----------------------------
-- Table structure for sa_user
-- ----------------------------
DROP TABLE IF EXISTS `sa_user`;
CREATE TABLE `sa_user`  (
  `user_id` int NOT NULL AUTO_INCREMENT COMMENT 'UID',
  `user_icon` int NOT NULL DEFAULT 0,
  `user_sex` int NOT NULL DEFAULT 0,
  `user_account` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '帐号',
  `user_password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `user_salt` varchar(4) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `user_head` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'new/images/nohead.jpg',
  `user_remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '每个人都应该有签名,但偏偏我没有.',
  `user_group` int NOT NULL DEFAULT 0 COMMENT '用户组',
  `user_ipreg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '注册IP',
  `user_openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `user_extra` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `user_app` int NOT NULL DEFAULT 1,
  `user_device` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `user_touchtip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `user_vip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `user_status` int NOT NULL DEFAULT 0 COMMENT '1被禁用',
  `user_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `user_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `admin_account`(`user_account`) USING BTREE,
  INDEX `admin_group`(`user_group`) USING BTREE,
  INDEX `admin_name`(`user_name`) USING BTREE,
  INDEX `admin_password`(`user_password`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 301 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sa_user
-- ----------------------------
INSERT INTO `sa_user` VALUES (1, 1, 0, 'admin@bbbug.com', '8b527446ed9bdfdd264716282f29f959', 'zXZy', '%E6%9C%BA%E5%99%A8%E4%BA%BA', 'new/images/nohead.jpg', '别@我,我只是个测试号', 1, '127.0.0.1', '', '', 1, 'iPhone', '%EF%BC%8C%E6%9C%BA%E5%99%A8%E4%BA%BA%E5%B7%AE%E7%82%B9%E7%88%BD%E7%BF%BB%E5%A4%A9%E3%80%82', '', 0, 0, 1605004436);
INSERT INTO `sa_user` VALUES (200, 0, 1, 'test@qq.com', '8b527446ed9bdfdd264716282f29f959', 'zXZy', 'Hammaa', 'http://localhost:8080/uploads/thumb/image/75cabffb5ad0061352ed05eca0fbe1d8.png', '欢迎来BBBUG聊天听歌划水呀~', 0, '::1', '', '', 0, 'pc', '帅气的脸庞', '', 0, 1648523967, 1648619443);
INSERT INTO `sa_user` VALUES (300, 0, 1, '6266@qq.com', '814f0f4863f42145ba86912a63a7f23c', 'CquV', '12312321', 'http://localhost:8080/uploads/thumb/image/75cabffb5ad0061352ed05eca0fbe1d8.png', '欢迎来BBBUG聊天听歌划水呀~', 0, '::1', '', '', 0, 'pc', '帅气的脸庞', '', 0, 1648523967, 1648619443);

-- ----------------------------
-- Table structure for sa_weapp
-- ----------------------------
DROP TABLE IF EXISTS `sa_weapp`;
CREATE TABLE `sa_weapp`  (
  `weapp_id` int NOT NULL AUTO_INCREMENT,
  `weapp_openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'OPENID',
  `weapp_status` int NOT NULL DEFAULT 0 COMMENT '状态',
  `weapp_createtime` int NOT NULL DEFAULT 0 COMMENT '创建时间',
  `weapp_updatetime` int NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`weapp_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '小程序用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sa_weapp
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
