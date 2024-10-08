
/* USER */
EXEC CreateIntIdPKTable
	'user',
	'
		[email] VARCHAR(255) UNIQUE NOT NULL,
		[password] VARCHAR(255) NOT NULL,
		[full_name] VARCHAR(100),
		[phone_number] VARCHAR(10),
		[role]	VARCHAR(10) NOT NULL DEFAULT ''customer'',
		CONSTRAINT chk_role CHECK (role IN (''customer'', ''cashier'', ''manager''))
	'
    