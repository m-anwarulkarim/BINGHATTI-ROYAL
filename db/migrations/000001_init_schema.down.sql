-- Drop Tables in reverse order of foreign key dependencies
DROP TABLE IF EXISTS leads;
DROP TABLE IF EXISTS units;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS users;

-- Drop Enum Types
DROP TYPE IF EXISTS lead_status;
DROP TYPE IF EXISTS unit_status;
DROP TYPE IF EXISTS user_role;
