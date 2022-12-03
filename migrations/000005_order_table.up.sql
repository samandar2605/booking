CREATE TABLE if not exists "order"(
    "id" serial primary KEY ,
    "hotel_id" INTEGER REFERENCES hotels(id)ON DELETE CASCADE,
    "room_id" INTEGER REFERENCES rooms(id)ON DELETE CASCADE,
    "full_name" VARCHAR(255) NOT NULL,
    "date_first" TIMESTAMP WITH TIME ZONE,
    "date_last" TIMESTAMP WITH TIME ZONE,
    "adults_count" INTEGER NOT NULL default 0,
    "children_count" INTEGER NOT NULL default 0,
    "user_id" INTEGER REFERENCES users(id)ON DELETE CASCADE
);
 
 
 
 
 
