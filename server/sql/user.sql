
	-- USER
EXEC CreateIntIdPKTable
	'user',
	'
		[email]			VARCHAR(255) UNIQUE NOT NULL,
		[password]		VARCHAR(255) NOT NULL,
		[full_name]		VARCHAR(100),
		[phone_number]	VARCHAR(10),
		[role_code]		VARCHAR(10) DEFAULT ''8001'',
		CONSTRAINT FK_user_role FOREIGN KEY ([role_code]) REFERENCES [user_role](role_code)
	'
    
	-- USER ROLE
CREATE TABLE [user_role] (
	role_code	VARCHAR(10) PRIMARY KEY,
	role_value	NVARCHAR(20)
)

INSERT INTO [user_role] (role_code, role_value) VALUES
('8001', 'customer'),
('8002', 'cashier'),
('8003', 'manager')