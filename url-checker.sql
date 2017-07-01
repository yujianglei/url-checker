CREATE TABLE `instance_url_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `item` varchar(1000) DEFAULT NULL,
  `instance_name` varchar(100) DEFAULT NULL,
  `url_type` varchar(30) DEFAULT NULL,
  `timeout` int(11) DEFAULT '10',
  `keyword` varchar(100) DEFAULT NULL,
  `remark` varchar(2000) DEFAULT NULL,
  `maintainer` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=164 DEFAULT CHARSET=utf8;
