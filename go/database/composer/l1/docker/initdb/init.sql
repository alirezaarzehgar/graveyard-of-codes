CREATE TABLE book
(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`create_time` timestamp DEFAULT NULL,
	`update_time` timestamp DEFAULT NULL,
	`name` VARCHAR(255) DEFAULT NULL,
	`author_name` VARCHAR(255) DEFAULT NULL,
	PRIMARY KEY(id)
);