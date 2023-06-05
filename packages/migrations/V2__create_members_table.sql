CREATE TABLE IF NOT EXISTS members
(
  member_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  member_alias_id VARCHAR(10) UNIQUE NOT NULL,
  family_name VARCHAR (50) UNIQUE NOT NULL
);