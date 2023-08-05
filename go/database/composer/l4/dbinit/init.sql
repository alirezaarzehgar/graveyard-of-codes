CREATE TABLE products (
	`id` int AUTO_INCREMENT NOT NULL,
	`name` VARCHAR(255) NOT NULL,
	`price` int NOT NULL,
	`rate` int DEFAULT 0,
	PRIMARY KEY(id)
);