CREATE TABLE if not exists "likes"(
    "id" serial PRIMARY KEY,
    "hotel_id" INTEGER REFERENCES hotels(id)ON DELETE CASCADE,
    "user_id" INTEGER REFERENCES users(id)ON DELETE CASCADE,
    "status" BOOLEAN NOT NULL,
    UNIQUE(hotel_id, user_id)
);
