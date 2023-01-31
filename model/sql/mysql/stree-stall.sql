/*
 Navicat Premium Data Transfer

 Source Server         : 本地环境
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : 127.0.0.1:3306
 Source Schema         : stree-stall

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 07/07/2022 18:52:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_admin
-- ----------------------------
DROP TABLE IF EXISTS `tb_admin`;
CREATE TABLE `tb_admin`  (
  `a_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员id',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '管理员昵称,在六个字符之内',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '管理员密码',
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '管理员头像链接',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`a_id`) USING BTREE,
  UNIQUE INDEX `username_unique`(`username`) USING BTREE COMMENT '管理员用户名唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_admin
-- ----------------------------

-- ----------------------------
-- Table structure for tb_booth
-- ----------------------------
DROP TABLE IF EXISTS `tb_booth`;
CREATE TABLE `tb_booth`  (
  `b_id` bigint(20) UNSIGNED NOT NULL COMMENT '摊位id',
  `type` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '摊位类型',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '摊位位置',
  `status` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '摊位状态 0 正常 1 禁用',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '摊位名称，在十个字之内',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '摊主用户id',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '摊主用户名',
  `longitude` decimal(10, 7) NULL DEFAULT NULL COMMENT '经度',
  `latitude` decimal(10, 7) NULL DEFAULT NULL COMMENT '纬度',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`b_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_booth
-- ----------------------------

-- ----------------------------
-- Table structure for tb_notice
-- ----------------------------
DROP TABLE IF EXISTS `tb_notice`;
CREATE TABLE `tb_notice`  (
  `n_id` bigint(20) UNSIGNED NOT NULL COMMENT '公告id',
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公告标题，10个字符之内',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公告内容，240个字符之内',
  `admin_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '管理员id',
  `admin_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '管理员用户名',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`n_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_notice
-- ----------------------------

-- ----------------------------
-- Table structure for tb_picture_wares
-- ----------------------------
DROP TABLE IF EXISTS `tb_picture_wares`;
CREATE TABLE `tb_picture_wares`  (
  `p_id` bigint(20) NOT NULL COMMENT '图片id',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片简介，20个字符之内',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片链接',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '图片状态  0 正常 1 禁用',
  `wares_id` bigint(20) NULL DEFAULT NULL COMMENT '商品id',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`p_id`) USING BTREE,
  UNIQUE INDEX `url_unique`(`url`) USING BTREE COMMENT '图片链接唯一'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_picture_wares
-- ----------------------------

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `u_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名，长度不能超过六个字符',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `email` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '电话，长度不超过11个字符，且全为数字',
  `user_type` smallint(1) UNSIGNED NULL DEFAULT NULL COMMENT '用户类型 0 商户（卖家） 1 游客用户（买家）',
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户头像链接',
  `create_at` datetime NULL DEFAULT NULL COMMENT '用户信息创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '用户信息更改时间',
  PRIMARY KEY (`u_id`, `username`) USING BTREE,
  UNIQUE INDEX `username_unique`(`username`) USING BTREE COMMENT '用户名唯一',
  UNIQUE INDEX `email_unique`(`email`) USING BTREE COMMENT '用户邮箱唯一',
  UNIQUE INDEX `phone_unique`(`phone`) USING BTREE COMMENT '用户电话唯一'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user
-- ----------------------------

-- ----------------------------
-- Table structure for tb_wares
-- ----------------------------
DROP TABLE IF EXISTS `tb_wares`;
CREATE TABLE `tb_wares`  (
  `w_id` bigint(20) UNSIGNED NOT NULL COMMENT '商品id',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商品名称，10个字符以内',
  `type` tinyint(255) UNSIGNED NULL DEFAULT NULL COMMENT '商品类型',
  `price` decimal(18, 2) UNSIGNED NULL DEFAULT NULL COMMENT '商品价格',
  `stock` int(255) UNSIGNED NULL DEFAULT NULL COMMENT '商品库存',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `status` tinyint(255) UNSIGNED NULL DEFAULT NULL COMMENT '商品状态',
  `booth_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '商品摊位id',
  `create_id` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_id` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`w_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_wares
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
