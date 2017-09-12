CREATE SCHEMA `friend_management`;


CREATE TABLE `members` (
  `entity_id` char(36) NOT NULL,
  `email` varchar(45) NOT NULL,
  `created` timestamp NOT NULL,
  `updated` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`entity_id`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `connections` (
  `member_entity_id` char(36) NOT NULL,
  `friend_entity_id` char(36) NOT NULL,
  `created` timestamp NOT NULL,
  `updated` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`member_entity_id`,`friend_entity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `updates` (
  `member_entity_id` char(36) NOT NULL,
  `target_entity_id` char(36) NOT NULL,
  `is_blocked` int(1) DEFAULT '0',
  `created` timestamp NOT NULL,
  `updated` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`member_entity_id`,`target_entity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
