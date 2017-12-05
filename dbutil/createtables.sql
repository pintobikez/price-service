CREATE DATABASE priceservice;

USE priceservice;

CREATE TABLE IF NOT EXISTS `price` (
  `id` varchar(24) NOT NULL,
  `fk_channel` int(11) NOT NULL,
  `price` double(16,2) NOT NULL DEFAULT '0',
  `special_price` double(16,2),
  `special_from_date` datetime,
  `special_to_date` datetime,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  KEY `productChannel` (`id`,`fk_channel`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `channels` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(45) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE INDEX idx_name ON channels (name);
ALTER TABLE `priceservice`.`price` ADD INDEX `fk_price_channel_idx` (`fk_channel` ASC);
ALTER TABLE `priceservice`.`price` ADD CONSTRAINT `fk_price_channel`  FOREIGN KEY (`fk_channel`)
  REFERENCES `priceservice`.`channels` (`id`)  ON DELETE NO ACTION  ON UPDATE NO ACTION;


docker run -d --hostname price-rabbit -e RABBITMQ_DEFAULT_USER=root -e RABBITMQ_DEFAULT_PASS=root rabbitmq:3-management