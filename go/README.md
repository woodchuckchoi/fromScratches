# Test in progress...

## things to remember

test MySQL container running on localhost:3306\

	SET GLOBAL time_zone='Asia/Seoul';
	// Global timezone for the server
	SET time_zone='Asia/Seoul';
	// Each client has its own session timezone

## Front

GitHub like table structure that indicates the overall blood sugar flow
Click on one particle -> go on

## DB

* user
	* id INT PRIMARY KEY
	* name VARCHAR(20) NOT NULL
	* uuid VARCHAR(36) NOT NULL
	* low SMALLINT
	* high SMALLINT

* health
	* id INT PRIMARY KEY
	* user\_id FOREIGN KEY REFERENCES user(id)
	* blood\_sugar SMALLINT NOT NULL
	* ts TIMESTAMP NOT NULL

* link
	* id INT PRIMARY KEY
	* user\_id FOREIGN KEY REFERENCES user(id)
	* link VARCHAR(30) UNIQUE NOT NULL
