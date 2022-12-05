CREATE TABLE if not exists "hotel_images"(
    "id" serial PRIMARY KEY,
    "hotel_id" INTEGER REFERENCES hotels(id)ON DELETE CASCADE,
    "image_url" VARCHAR(255) NOT NULL,
    "sequence_number" INTEGER NOT NULL
);
