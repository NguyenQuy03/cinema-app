
	-- DIRECTOR
EXEC CreateIntIdPKTable
	'director',
	'
		[director_name] NVARCHAR(100)
	';

	-- MOVIE
EXEC CreateIntIdPKTable
	'movie', 
	'
		[director_id]	INT, 
		[title]			NVARCHAR(255), 
		[status]		NVARCHAR(20), 
		[age_rating]	VARCHAR(3), 
		[duration]		INT, 
		[release_date]	DATETIME, 
		[trailer_link]	VARCHAR(255), 
		[description]	NVARCHAR(2000), 
		[poster_img]	VARCHAR(255), 
		[header_img]	VARCHAR(255),
		FOREIGN KEY ([director_id]) REFERENCES director([id])
	'

-- CAST_MEMBER
EXEC CreateIntIdPKTable
	'cast_member',
	'
		[cast_name] NVARCHAR(100),
		[dob]		DATETIME DEFAULT NULL,
		[bio]		NVARCHAR(MAX) DEFAULT NULL
	';

-- GENRE
EXEC CreateIntIdPKTable
	'genre',
	'
		[genre_slug]	VARCHAR(20) UNIQUE,
		[genre_name]	NVARCHAR(20),
	';


-- MOVIE_GENRE
CREATE TABLE movie_genre (
	[movie_id] INT REFERENCES movie([id]),
	[genre_id] INT REFERENCES genre([id]),
	CONSTRAINT PK_movie_genre PRIMARY KEY([movie_id], [genre_id])
)

-- MOVIE_CAST
CREATE TABLE movie_cast (
	[movie_id]	INT REFERENCES movie([id]),
	[cast_id]	INT REFERENCES cast_member([id])	
	CONSTRAINT PK_movie_cast PRIMARY KEY([movie_id], [cast_id])
)
