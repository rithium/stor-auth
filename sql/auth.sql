CREATE TABLE `stor`.`apikey` (
  `id` int(11) NOT NULL,
  `key` varchar(64) NOT NULL,
  `active` tinyint(4) NOT NULL,
  `created` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `apikey`
  ADD PRIMARY KEY (`id`);
