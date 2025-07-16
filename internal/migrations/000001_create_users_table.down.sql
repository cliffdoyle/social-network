-- Drop the trigger first
DROP TRIGGER IF EXISTS update_users_updated_at;

-- Drop indexes
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_nickname;
DROP INDEX IF EXISTS idx_users_created_at;
DROP INDEX IF EXISTS idx_users_is_private;

-- Drop the users table
DROP TABLE IF EXISTS users;