CREATE DATABASE IF NOT EXISTS `main_db`;
USE `main_db`;


CREATE TABLE IF NOT EXISTS `Users` (
  `UserID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `Name` varchar(50) NOT NULL DEFAULT '0',
  `Mail` varchar(50) NOT NULL DEFAULT '0',
  `Password` varchar(50) NOT NULL,
  `CreateTime` int(11) NOT NULL COMMENT '帳號建立時間',
  `UpdateTime` int(11) NOT NULL COMMENT '最後登入時間',
  UNIQUE KEY `UserID` (`UserID`),
  UNIQUE KEY `Mail` (`Mail`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='玩家帳號資料';
