
	-- PLACE
EXEC CreateIntIdPKTable
	'place',
	'
		[place_name]	NVARCHAR(50),
		[place_slug]	VARCHAR(50) UNIQUE,
	';

	-- CINEMA
EXEC CreateIntIdPKTable
	'cinema',
	'
		[cinema_name]	NVARCHAR(50),
		[cinema_slug]	VARCHAR(50) UNIQUE,
		[place_id]		INT,
		CONSTRAINT FK_cienma_place FOREIGN KEY (place_id) REFERENCES place(id),
	';

	-- ACCESSIBILITY
EXEC CreateIntIdPKTable
	'accessibility',
	'
		[acc_feature]	NVARCHAR(50) UNIQUE,
		[shorten]		VARCHAR(10),
	';

	-- EXPERIENCE
EXEC CreateIntIdPKTable
	'experience',
	'
		[exp_feature]	NVARCHAR(50) UNIQUE,
		[description]	NVARCHAR(100),
		CONSTRAINT uniq_feature UNIQUE ([exp_feature])
	';

	-- THEATER
EXEC CreateIntIdPKTable
	'theater',
	'
		theater_num      INT,
		cinema_id        INT,
		acc_id           INT,
		exp_id           INT,
		CONSTRAINT FK_thea_cinema FOREIGN KEY (cinema_id) REFERENCES cinema(id),
		CONSTRAINT FK_thea_accessibility FOREIGN KEY (acc_id) REFERENCES accessibility(id),
		CONSTRAINT FK_thea_experience FOREIGN KEY (exp_id) REFERENCES experience(id),
		CONSTRAINT UQ_theater_num_cinema UNIQUE (theater_num, cinema_id)
	';

	-- SHOWING_TIME
EXEC CreateIntIdPKTable
	'showing_time',
	'
		[movie_id]      INT,
		[theater_id]    INT,
		[show_date]     DATETIME DEFAULT SYSDATETIME(),
		[base_price]    INT,
		CONSTRAINT FK_show_movie FOREIGN KEY ([movie_id]) REFERENCES movie(id),
		CONSTRAINT FK_show_theater FOREIGN KEY ([theater_id]) REFERENCES theater(id)
	';