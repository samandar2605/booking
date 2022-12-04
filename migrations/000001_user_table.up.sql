CREATE TABLE if not exists "users"(
    "id" serial PRIMARY KEY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255),
    "phone_number" VARCHAR(255),
    "email" VARCHAR(255) NOT NULL,
    "username" VARCHAR(255),
    "password" VARCHAR(255) NOT NULL,
    "profile_image" VARCHAR(255),
    "type" VARCHAR(255)CHECK("type" IN('superadmin','admin','user')) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE default current_timestamp
);