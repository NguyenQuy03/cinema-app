-- Procedure help generate `created_at`, `updated_at`, `deleted_at`
-- And create trgger when update

CREATE PROCEDURE CreateIntIdPKTable
    @TableName NVARCHAR(255),
    @AdditionalFields NVARCHAR(MAX) -- Accept additional fields as a string
AS
BEGIN
    DECLARE @SQL NVARCHAR(MAX);
    DECLARE @TriggerSQL NVARCHAR(MAX);
    DECLARE @TriggerName NVARCHAR(255);

    -- Create the table with timestamp fields and additional fields
    SET @SQL = 'CREATE TABLE ' + QUOTENAME(@TableName) + ' (
        id INT PRIMARY KEY IDENTITY(1,1),
        created_at DATETIME DEFAULT SYSDATETIME(),
        updated_at DATETIME DEFAULT SYSDATETIME(),
        deleted_at DATETIME NULL, ' + @AdditionalFields + ');';
    
    -- Execute the table creation
    EXEC sp_executesql @SQL;

    -- Construct the trigger name properly, replacing spaces in table names with underscores
    SET @TriggerName = 'trg_Update_' + REPLACE(@TableName, ' ', '_');

    -- Create the update trigger
    SET @TriggerSQL = 'CREATE TRIGGER ' + @TriggerName + '
    ON ' + QUOTENAME(@TableName) + '
    AFTER UPDATE
    AS
    BEGIN
        UPDATE ' + QUOTENAME(@TableName) + '
        SET updated_at = SYSDATETIME()
        WHERE Id IN (SELECT Id FROM inserted);
    END;';

    -- Execute the trigger creation
    EXEC sp_executesql @TriggerSQL;
END;