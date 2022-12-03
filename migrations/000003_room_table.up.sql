CREATE TABLE if not exists "rooms"(
    "id" serial PRIMARY KEY,
    "hotel_id" INTEGER REFERENCES hotels(id)ON DELETE CASCADE,
    "image_url" VARCHAR(255),
    "is_active" BOOLEAN NOT NULL,
    "room_type" VARCHAR(255),
    "price_for_children" DECIMAL(18, 2) NOT NULL,
    "price_for_adults" DECIMAL(18, 2) NOT NULL
);
