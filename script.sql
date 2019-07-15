CREATE TABLE IF NOT EXISTS "user" (
  id SERIAL PRIMARY KEY,
  username varchar(50) UNIQUE NOT NULL,
  password varchar(255) NOT NULL,
  email    varchar(255) UNIQUE NOT NULL,
  created_at timestamp,
  updated_at timestamp
);


CREATE TABLE IF NOT EXISTS "session" (
  sessionid varchar(256) PRIMARY KEY UNIQUE NOT NULL,
  username varchar(50),
  agent   varchar(255),
  expires_at timestamp,
  created_at timestamp,
  updated_at timestamp
);