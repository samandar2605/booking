CREATE TABLE if not exists "room_images"(
    "id" serial PRIMARY KEY,
    "room_id" INTEGER  REFERENCES rooms(id)ON DELETE CASCADE,
    "image_url" VARCHAR(255) NOT NULL,
    "sequence_number" INTEGER NOT NULL
);
