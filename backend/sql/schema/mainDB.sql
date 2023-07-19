CREATE DATABASE IF NOT EXISTS `main_db`
USE `main_db`;


CREATE TABLE IF NOT EXISTS `Bank` (
  `UserID` int(10) unsigned NOT NULL,
  `Coin` bigint(20) unsigned zerofill NOT NULL DEFAULT 00000000000000000000 COMMENT '遊戲貨幣 (物品買賣 & 任務達成獲得)',
  `Faith` bigint(20) unsigned zerofill NOT NULL DEFAULT 00000000000000000000 COMMENT '遊戲貨幣 (成就 & 任務達成獲得)',
  `Gems` bigint(20) unsigned zerofill NOT NULL DEFAULT 00000000000000000000 COMMENT '商城貨幣',
  `Treasure` bigint(20) unsigned zerofill NOT NULL DEFAULT 00000000000000000000 COMMENT '商城貨幣 (課金)',
  UNIQUE KEY `UserID` (`UserID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='玩家貨幣資料';


CREATE TABLE IF NOT EXISTS `Users` (
  `UserID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `Token` varchar(50) NOT NULL DEFAULT '0',
  `Name` varchar(50) NOT NULL DEFAULT '0',
  `Mail` varchar(50) NOT NULL DEFAULT '0',
  `CreateTime` int(11) NOT NULL COMMENT '帳號建立時間',
  `UpdateTime` int(11) NOT NULL COMMENT '最後登入時間',
  UNIQUE KEY `UserID` (`UserID`),
  UNIQUE KEY `Token` (`Token`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='玩家帳號資料';

