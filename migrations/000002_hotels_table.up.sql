CREATE TABLE if not exists "hotels"(
    "id" serial PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255),
    "address" VARCHAR(255) NOT NULL,
    "rating" DECIMAL(8, 2),
    "rooms_count" integer default 0
);
