CREATE TABLE IF NOT EXISTS "users"(
    "id" VARCHAR PRIMARY KEY,
    "first_name" VARCHAR NOT NULL,
    "last_name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL, 
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "posts"(
    "id" VARCHAR PRIMARY KEY,
    "user_id" VARCHAR NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    "title" VARCHAR NOT NULL,
    "body" TEXT NOT NULL,
    "published" BOOLEAN DEFAULT FALSE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
