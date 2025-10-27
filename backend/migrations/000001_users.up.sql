CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
	id          UUID         PRIMARY KEY,
	login       VARCHAR(255) NOT NULL UNIQUE,
	name        VARCHAR(100) NOT NULL,
	sex         VARCHAR(5)   NOT NULL CHECK(sex IN('male', 'female')),
	age         INTEGER      NOT NULL CHECK(age > 0),
	city        VARCHAR(50)  NOT NULL DEFAULT '',
	
	telegram_id BIGINT       NOT NULL DEFAULT 0,
	
	created_at  TIMESTAMP    NOT NULL DEFAULT now()
);