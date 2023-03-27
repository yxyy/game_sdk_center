/*
 Navicat Premium Data Transfer

 Source Server         :
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           :
 Source Schema         : game_sdk_center

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 27/03/2023 18:53:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS `game_sdk_center`;
CREATE DATABASE `game_sdk_center`;
use `game_sdk_center`;

-- ----------------------------
-- Table structure for game
-- ----------------------------
DROP TABLE IF EXISTS `game`;
CREATE TABLE `game`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_project_id` mediumint(8) NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `game_name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '应用名',
  `cp_interface_url` varchar(300) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'CP 发货正式接口',
  `cp_pay_secret` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '发货key',
  `app_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用类型：1 Android 2 IOS 3 MiniGame',
  `app_config` varchar(4096) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '应用配置,即对接配置',
  `product_key` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '产品密钥，可用于初始化及加密传输',
  `algo` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT 'AES' COMMENT '数据加密算法',
  `link_h5` varchar(300) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'H5 链接 ',
  `app_store_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '苹果应用ID',
  `down_url` varchar(500) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '包uri',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态: 关闭(1)、正常(0)',
  `add_time` int(11) NOT NULL DEFAULT 0 COMMENT '添加时间',
  `last_time` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `pkg_name` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '包名请勿重复，否则可能会影响第三方登录和支付',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '游戏应用包，主要用于分发' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for game_iap_goods
-- ----------------------------
DROP TABLE IF EXISTS `game_iap_goods`;
CREATE TABLE `game_iap_goods`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0,
  `goods_type` tinyint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品类型：1（苹果）、2（谷歌）',
  `app_store_id` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '苹果应用ID',
  `developer` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '苹果开发者',
  `goods_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '内购商品名',
  `goods_desc` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '内购商品描述',
  `goods_id` varchar(120) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '内购ID',
  `amount` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '金额单位分',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '商品状态: 关闭(1)、正常(0)',
  `add_time` int(11) NOT NULL DEFAULT 0 COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `game_product_id`(`goods_id`, `game_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '内购商品表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for game_notice
-- ----------------------------
DROP TABLE IF EXISTS `game_notice`;
CREATE TABLE `game_notice`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `notice_type` int(10) NOT NULL DEFAULT 1 COMMENT '状态: 0 正常  1流畅 ',
  `game_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `start_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `end_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `account` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '管理员账号',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态: 0 正常  1流畅 ',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for game_project
-- ----------------------------
DROP TABLE IF EXISTS `game_project`;
CREATE TABLE `game_project`  (
  `id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `project_name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `game_class` tinyint(1) NULL DEFAULT 1 COMMENT '游戏类别',
  `cp_info` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `add_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `last_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `g`(`project_name`) USING BTREE,
  INDEX `game_class`(`game_class`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '游戏资料表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for game_sdk_ui
-- ----------------------------
DROP TABLE IF EXISTS `game_sdk_ui`;
CREATE TABLE `game_sdk_ui`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `app_id` int(10) NOT NULL DEFAULT 0 COMMENT '游戏应用ID',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `game_id`(`game_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '应用Sdk UI' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for infra_config
-- ----------------------------
DROP TABLE IF EXISTS `infra_config`;
CREATE TABLE `infra_config`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `group` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '参数分组',
  `type` tinyint(4) NOT NULL COMMENT '参数类型',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '参数名称',
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '参数键名',
  `value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '参数键值',
  `sensitive` bit(1) NOT NULL COMMENT '是否敏感',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '参数配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ios_blur
-- ----------------------------
DROP TABLE IF EXISTS `ios_blur`;
CREATE TABLE `ios_blur`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `game_id` int(11) NOT NULL DEFAULT 0,
  `url_list` json NULL,
  `params_map` json NULL,
  `url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `function_map` json NULL,
  `status` tinyint(2) NOT NULL DEFAULT 1 COMMENT '状态，1=禁用。2=启用',
  `gs_url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `game_id_blur` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  `algo` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `only_key`(`game_id`) USING BTREE,
  INDEX `url`(`url`) USING BTREE,
  INDEX `game_id_blur`(`game_id_blur`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'ios混淆表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for ios_shm_url
-- ----------------------------
DROP TABLE IF EXISTS `ios_shm_url`;
CREATE TABLE `ios_shm_url`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '提审域名',
  `game_id` int(11) NULL DEFAULT 0 COMMENT '游戏id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `url`(`url`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '提审域名与ios渠道包' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for money_record_all
-- ----------------------------
DROP TABLE IF EXISTS `money_record_all`;
CREATE TABLE `money_record_all`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `order_num` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `money_type` tinyint(1) NULL DEFAULT NULL COMMENT '帐变类型',
  `money` int(10) NOT NULL DEFAULT 0 COMMENT '帐变金额,单位分',
  `money_after` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帐变后',
  `money_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帐变时间',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `money_time`(`money_time`) USING BTREE,
  INDEX `order_num`(`order_num`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_channel
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel`;
CREATE TABLE `pay_channel`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `merchant_id` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付商户ID',
  `channel` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '支付通道：wxpay、alipay',
  `code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '支付编号：扫码、App、Wap、小程序等',
  `display_name` varchar(45) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `fee_rate` double NOT NULL DEFAULT 0 COMMENT '渠道费率，单位：百分比\',',
  `channel_sort` int(10) NOT NULL DEFAULT 1 COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态: 关闭(1)、正常(0)',
  `config` varchar(4096) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '配置',
  `remark` varchar(800) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `tc`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '支付通道配置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_google_service
-- ----------------------------
DROP TABLE IF EXISTS `pay_google_service`;
CREATE TABLE `pay_google_service`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏id',
  `package_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '应用的包名',
  `public_key` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公钥',
  `json_content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT 'googlePay 配置信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `package`(`game_id`, `package_name`) USING BTREE,
  INDEX `productId`(`game_id`) USING BTREE,
  INDEX `packageName`(`package_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'googlePay 配置信息表-一个游戏一个' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for pay_merchant
-- ----------------------------
DROP TABLE IF EXISTS `pay_merchant`;
CREATE TABLE `pay_merchant`  (
  `id` smallint(5) NOT NULL COMMENT '商户编号',
  `no` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商户号',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商户全称',
  `short_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商户简称',
  `status` tinyint(4) NOT NULL COMMENT '开启状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '‘’' COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '支付商户信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_order_google
-- ----------------------------
DROP TABLE IF EXISTS `pay_order_google`;
CREATE TABLE `pay_order_google`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `order_num` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '平台订单号',
  `order_no` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'google订单号',
  `hash_value` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'purchaseToken',
  `product_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '商品id',
  `ctime` int(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `order_num`(`order_num`) USING BTREE,
  UNIQUE INDEX `order_no`(`order_no`) USING BTREE,
  INDEX `hash_value`(`hash_value`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for pay_orders
-- ----------------------------
DROP TABLE IF EXISTS `pay_orders`;
CREATE TABLE `pay_orders`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `merchant_id` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '支付商户ID',
  `pay_channel` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '支付通道对应表pay_channel的 channel_code',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `openid` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '三方开放的帐号ID 如 vivo、wx 等',
  `order_type` tinyint(1) NULL DEFAULT 1 COMMENT '订单类型：直购(1)、应用购(2)、充余额(10)、提现(20)',
  `order_num` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `order_price` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '定价,单位分',
  `subject` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '物品描述',
  `order_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `pay_money` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '实付金额,单位分',
  `currency` varchar(45) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT 'USD' COMMENT '货币代码:USD(美元)、HKG(港元) 、MAC(澳门元) 、TWD(新台币) ...',
  `pay_status` tinyint(1) NULL DEFAULT -1,
  `pay_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `trade_id` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `trade_account` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `trade_status` tinyint(1) NOT NULL DEFAULT 0,
  `trade_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '三方交易时间',
  `trade_data` varchar(500) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '交易数据',
  `goods_id` varchar(120) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '对就商口goods_id',
  `cp_order_num` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT 'CP 订单号',
  `cp_order_status` tinyint(1) NULL DEFAULT -1 COMMENT '发货状态： -1 未发 0 失败 1 成功',
  `cp_order_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发货时间',
  `game_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `game_name` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '游戏名',
  `zone_id` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '0' COMMENT '区服',
  `zone_name` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '区服名',
  `game_role_id` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '0',
  `game_role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `game_role_level` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '0' COMMENT '角色级别',
  `os` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '操作系统：android、ios',
  `device_id` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `ip` varchar(15) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `sandbox` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '环境: 测试(1)、生产(0)',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `order_num`(`order_num`) USING BTREE,
  INDEX `member_id`(`user_id`) USING BTREE,
  INDEX `order_time`(`order_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 109 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '订单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for pay_switch
-- ----------------------------
DROP TABLE IF EXISTS `pay_switch`;
CREATE TABLE `pay_switch`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_channel` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '渠道即带有SDK的渠道如 VIVO',
  `game_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `start_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开始时间',
  `end_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '结束时间',
  `pay_times` int(10) NOT NULL DEFAULT 0 COMMENT '充值次数',
  `pay_money` int(10) NOT NULL DEFAULT 0 COMMENT '充值总额',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态: 关闭(1)、正常(0)',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE,
  INDEX `channel`(`app_channel`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '应用切支付' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for promote
-- ----------------------------
DROP TABLE IF EXISTS `promote`;
CREATE TABLE `promote`  (
  `id` bigint(20) UNSIGNED NOT NULL,
  `game_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `app_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '应用类型：1 Android 2 IOS 3 MiniGame',
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `media` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT 'my' COMMENT '媒体编号',
  `ad_agent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '广告代理ID',
  `ad_account_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '媒体帐号',
  `advertiser_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '广告主ID',
  `ad_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '媒体广告名',
  `ad_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '媒体广告ID',
  `group_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '广告组名',
  `group_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '广告组id',
  `ad_status` tinyint(1) NOT NULL DEFAULT -1 COMMENT '广告状态: 创建中(-1)、正常(1)、暂停(0)',
  `ad_data` varchar(4096) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '媒体广告计划数据',
  `promote_link` varchar(245) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '推广链',
  `land_page_id` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '落地页ID',
  `land_page_link` varchar(245) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '落地页链',
  `monitor_link` varchar(245) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '监控链',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `promote_code`(`promote_code`) USING BTREE,
  INDEX `media`(`media`) USING BTREE,
  INDEX `ad_id`(`ad_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '广告计划' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for promote_app
-- ----------------------------
DROP TABLE IF EXISTS `promote_app`;
CREATE TABLE `promote_app`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '应用ID',
  `app_name` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '应用名',
  `pkg_name` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '包名',
  `media` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '媒体',
  `media_config` varchar(4096) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '媒体配置,包括SDK',
  `pack_time` int(10) NOT NULL DEFAULT -1 COMMENT '打包状态: -1未打 0失败 1成功',
  `pack_status` tinyint(1) NOT NULL DEFAULT -1 COMMENT '打包状态: 未打(-1)、关闭(0)、正常(1)',
  `file_uri` varchar(500) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '包uri',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态: 关闭(1)、正常(0)',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '应用分包' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for promote_domain
-- ----------------------------
DROP TABLE IF EXISTS `promote_domain`;
CREATE TABLE `promote_domain`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `domain` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT '',
  `domain_status` int(11) NOT NULL DEFAULT 1,
  `res_data` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for report_daily
-- ----------------------------
DROP TABLE IF EXISTS `report_daily`;
CREATE TABLE `report_daily`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `days` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `promote_code` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告码',
  `app_channel` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '应用渠道',
  `agent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理ID,即主播',
  `media` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '广告媒体如：oe、tc、ks',
  `advertiser_id` bigint(20) NULL DEFAULT 0 COMMENT '广告主ID',
  `app_id` int(10) NOT NULL DEFAULT 0 COMMENT '游戏应用ID',
  `pv` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `uv` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '独立IP',
  `click` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `active` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '激活即启动数',
  `dnu` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `dau` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `players` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏注册数：一个dnu对应N个players ',
  `enter` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `pay_money` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '付费金额',
  `fee` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '手续费',
  `pay_times` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '付费人次',
  `pay_numbers` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '付费人数',
  `ad_money` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '广告费，即消耗',
  `lt` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT 'LT生命周期(LT:Life Time)：一个用户从第1次到最后1次参与游戏之间的时间段，一般按月计算平均值',
  `keep1` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '次留',
  `keep2` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `keep3` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `keep7` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `keep15` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `keep30` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `dtaaa`(`days`, `app_channel`, `app_id`, `advertiser_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '广告日报' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for report_pay_channel
-- ----------------------------
DROP TABLE IF EXISTS `report_pay_channel`;
CREATE TABLE `report_pay_channel`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `days` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `tenant_id` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '租户编号',
  `merchant_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商户ID',
  `channel` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT 'default' COMMENT '渠道即带有SDK的渠道如 VIVO',
  `pay_channel` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '支付通道',
  `pay_money` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00,
  `pay_times` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '付费人次',
  `pay_numbers` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '付费人数',
  `order_price` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_ban
-- ----------------------------
DROP TABLE IF EXISTS `sys_ban`;
CREATE TABLE `sys_ban`  (
  `id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ban_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '封禁类型',
  `reason` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '原因',
  `object_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '被封禁的对象：IP(1)、设备(2)、帐号ID(3)',
  `objects` varchar(800) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '被封禁的对象',
  `start_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开始时间',
  `end_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '结束时间',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态： 1:被封 　0：解封',
  `account` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '登录名，不可改',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `member_id`(`objects`, `ban_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '封禁表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '父部门id',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `leader_user_id` bigint(20) NULL DEFAULT 0 COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint(4) NOT NULL COMMENT '部门状态（0正常 1停用）',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_at` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '部门表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '字典排序',
  `label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典标签',
  `value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `color_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '颜色类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'css 样式',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_at` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1255 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '字典数据表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典名称',
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_at` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 175 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '字典类型表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_id_generator
-- ----------------------------
DROP TABLE IF EXISTS `sys_id_generator`;
CREATE TABLE `sys_id_generator`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = 'id生成' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `display_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '中文名',
  `permission` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限标识',
  `path` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '访问的URL',
  `sorts` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '排序',
  `component` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '页面标签',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '菜单状态（0正常 1停用）',
  `icon` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'icon 暂用于商户后台',
  `resource_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '1：目录 2：菜单 3；按钮',
  `module_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1:中央后台 2:商户后台',
  `create_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `visible` tinyint(1) NOT NULL DEFAULT 1 COMMENT '可视：是(1)、否(0)',
  `keep_alive` tinyint(1) NOT NULL DEFAULT 0 COMMENT '缓存：是（1）、否（0）',
  `component_name` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '组件名字',
  `always_show` tinyint(1) NOT NULL DEFAULT 1 COMMENT '总是显示：是(1)、否(0)',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `module_sort`(`sorts`) USING BTREE,
  INDEX `module_type`(`module_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 72 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统菜单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位编码',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '岗位名称',
  `sort` int(11) NOT NULL COMMENT '显示顺序',
  `status` tinyint(4) NOT NULL COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_at` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '岗位信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：1 正常 0 禁用',
  `code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色权限字符串',
  `sort` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示顺序',
  `data_scope` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `data_scope_menu_ids` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单权限',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '角色类型',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `creator` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updater` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `delete_at` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission`  (
  `role_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '台子ID',
  `permission_id` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`role_id`, `permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_sessions
-- ----------------------------
DROP TABLE IF EXISTS `sys_sessions`;
CREATE TABLE `sys_sessions`  (
  `id` char(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `val` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `expiry` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `ip` char(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'IP ',
  `ua` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `add_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `expiry`(`expiry`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_setting
-- ----------------------------
DROP TABLE IF EXISTS `sys_setting`;
CREATE TABLE `sys_setting`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `source_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '来源ID：platformID or gameID',
  `set_name` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `set_val` mediumtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ss`(`set_name`, `source_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统设置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_sms_channel
-- ----------------------------
DROP TABLE IF EXISTS `sys_sms_channel`;
CREATE TABLE `sys_sms_channel`  (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `tenant_id` int(10) UNSIGNED NOT NULL,
  `signature` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '短信签名',
  `code` varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '渠道编码',
  `status` tinyint(4) NOT NULL COMMENT '开启状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `api_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '短信 API 的账号',
  `api_secret` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '短信 API 的秘钥',
  `callback_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '短信发送回调 URL',
  `creator` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '短信渠道' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_sms_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_sms_log`;
CREATE TABLE `sys_sms_log`  (
  `id` bigint(20) NOT NULL,
  `tenant_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '台子ID',
  `channel_id` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `channel_code` varchar(63) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `template_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模板ＩＤ',
  `template_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '模板类型',
  `mobile` char(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `code` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '短信验证码',
  `content` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `send_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '发送状态: 成功（1 ） 、失败（0）',
  `send_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `res` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '响应结果',
  `ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'IP',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `mobile`(`mobile`) USING BTREE,
  INDEX `ip`(`ip`) USING BTREE,
  INDEX `send_status`(`send_status`) USING BTREE,
  INDEX `platform_id`(`tenant_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '短信发送日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_sms_template
-- ----------------------------
DROP TABLE IF EXISTS `sys_sms_template`;
CREATE TABLE `sys_sms_template`  (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `type` tinyint(4) NOT NULL COMMENT '短信类型',
  `status` tinyint(4) NOT NULL COMMENT '开启状态',
  `code` varchar(63) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '模板编码',
  `name` varchar(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '模板名称',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '模板内容',
  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '参数数组',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '备注',
  `api_template_id` varchar(63) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '短信 API 的模板编号',
  `channel_id` bigint(20) NOT NULL COMMENT '短信渠道编号',
  `channel_code` varchar(63) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '短信渠道编码',
  `creator` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updater` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '短信模板' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_task_retry
-- ----------------------------
DROP TABLE IF EXISTS `sys_task_retry`;
CREATE TABLE `sys_task_retry`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_id` varchar(40) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `task_data` varchar(800) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT '',
  `retry_times` int(11) NOT NULL DEFAULT 0,
  `next_time` int(11) NOT NULL DEFAULT 0 COMMENT '下次执行时间',
  `is_ok` tinyint(1) NOT NULL DEFAULT 0,
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `task_id`(`task_id`) USING BTREE,
  INDEX `is_ok`(`is_ok`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户类型：玩家（1）、代理（2）、管理员（3）',
  `promote_code` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '推广码',
  `account_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '帐号类型：0  游客; 1 普通账号 ; 2 手机账号; 3  微信账号;4 机器人',
  `account` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '帐户名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登陆密码',
  `salt` char(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '-1:未激活，0：正常，1：禁用',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `pu`(`account`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 199583 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_user_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_log`;
CREATE TABLE `sys_user_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `log_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '日志类型',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `account` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '账号名',
  `request_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '请求ID',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '操作结果',
  `create_time` int(11) NOT NULL DEFAULT 0 COMMENT '添加时间',
  `result` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '操作结果',
  `descriptor` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '操作描述',
  `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '模块名',
  `ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'ip',
  `params` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '参数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `account`(`account`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9469 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '记录系统日志' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_user_profile
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_profile`;
CREATE TABLE `sys_user_profile`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `nickname` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `mobile` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '绑定的手机号',
  `trade_password` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '安全码',
  `balance` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '余额',
  `full_name` varchar(12) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `reg_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '注册时间',
  `reg_ip` char(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '注册IP',
  `last_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登陆时间',
  `last_ip` char(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '最后登陆IP',
  `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门ID',
  `post_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '岗位编号数组',
  `qq` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `email` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `sex` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '2:未知 1:男 0: 女',
  `avatar` varchar(400) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `vip` smallint(5) UNSIGNED NOT NULL DEFAULT 1 COMMENT '会员VIP等级',
  `login_times` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登陆次数',
  `remark` varchar(800) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `last_ip`(`last_ip`) USING BTREE,
  INDEX `last_time`(`last_time`) USING BTREE,
  INDEX `reg_time`(`reg_time`) USING BTREE,
  INDEX `reg_ip`(`reg_ip`) USING BTREE,
  INDEX `mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 159 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户扩展属性表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `role_id` varchar(25) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for sys_user_third
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_third`;
CREATE TABLE `sys_user_third`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '帐号id：member_id or user_id',
  `socialite` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '应用渠道即带有SDK的渠道如 Google、FaceBook',
  `openid` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '三方开放的帐号ID 如 vivo、wx 等',
  `open_account` varchar(200) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '',
  `unionid` varchar(128) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL DEFAULT '' COMMENT '联合ID',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ao`(`socialite`, `openid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci COMMENT = '三方用户' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
