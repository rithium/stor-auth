CREATE TABLE `stor`.`apiKey` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `key` varchar(64) NOT NULL,
  `active` tinyint(4) NOT NULL,
  `created` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;