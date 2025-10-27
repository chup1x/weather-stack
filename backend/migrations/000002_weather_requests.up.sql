CREATE TABLE IF NOT EXISTS weather_requests(
	id          SERIAL       PRIMARY KEY,
	user_id     UUID         REFERENCES users,
	temperature NUMERIC(5,2) NOT NULL,
	clothing    VARCHAR(50)  NOT NULL,
	
	created_at  TIMESTAMP    NOT NULL DEFAULT now()
);