CREATE TABLE if not exists "hotels"(
    "id" serial PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255),
    "address" VARCHAR(255) NOT NULL,
    "rating" DECIMAL(8, 2) default 0,
    "phone_number" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "price" DECIMAL(18,2) NOT NULL default 0,
    "rooms_count" integer default 0
);
