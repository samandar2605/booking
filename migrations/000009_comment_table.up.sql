CREATE TABLE if not exists "comments"(
    "id" serial PRIMARY KEY,
    "hotel_id" INTEGER REFERENCES hotels(id)ON DELETE CASCADE,
    "user_id" INTEGER REFERENCES users(id)ON DELETE CASCADE,
    "description" VARCHAR(255),
    "created_at" TIMESTAMP WITH TIME ZONE default current_timestamp,
    "updated_at" TIMESTAMP WITH TIME ZONE
);
