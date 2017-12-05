CREATE DATABASE priceservice;

USE priceservice;

CREATE TABLE `channels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE `price` (
  `id` varchar(24) NOT NULL,
  `fk_channel` int(11) NOT NULL,
  `price` double(16,2) NOT NULL DEFAULT '0.00',
  `special_price` double(16,2) DEFAULT NULL,
  `special_from_date` datetime DEFAULT NULL,
  `special_to_date` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `key` (`id`,`fk_channel`),
  KEY `productChannel` (`id`,`fk_channel`) USING BTREE,
  KEY `fk_price_channel_idx` (`fk_channel`),
  CONSTRAINT `fk_price_channel` FOREIGN KEY (`fk_channel`) REFERENCES `channels` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

docker run -d --hostname price-rabbit -e RABBITMQ_DEFAULT_USER=root -e RABBITMQ_DEFAULT_PASS=root rabbitmq:3-management