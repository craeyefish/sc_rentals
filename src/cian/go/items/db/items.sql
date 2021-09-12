DROP TABLE IF EXISTS `items`;
CREATE TABLE `items` (
  `name` varchar(32) NOT NULL,
  `code` int NOT NULL AUTO_INCREMENT,
  `type` varchar(45) NOT NULL,
  `value` varchar(45) DEFAULT NULL,
  `init_date` datetime DEFAULT NULL,
  PRIMARY KEY (`code`),
  UNIQUE KEY `code_UNIQUE` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;