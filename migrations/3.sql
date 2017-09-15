ALTER TABLE `members` 
ADD COLUMN `password` TEXT NOT NULL AFTER `email`,
ADD COLUMN `name` VARCHAR(45) NOT NULL AFTER `password`;
