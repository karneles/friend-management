ALTER TABLE `connections` 
DROP PRIMARY KEY,
ADD PRIMARY KEY (`member_entity_id`, `friend_entity_id`);
