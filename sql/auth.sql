create database stor;

CREATE TABLE `stor`.`apikey` ( `id` INT NOT NULL , `nodeId` INT NOT NULL , `key` INT NOT NULL , `active` INT NOT NULL , `created` INT NOT NULL ) ENGINE = InnoDB;