CREATE TABLE IF NOT EXISTS weather_requests(
	city_id     VARCHAR(30)  PRIMARY KEY,
	path 		VARCHAR(100) NOT NULL,
	
	created_at  TIMESTAMP    NOT NULL DEFAULT now()
);
