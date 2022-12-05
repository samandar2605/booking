CREATE TABLE if not exists "orders"(
    "id" serial primary KEY ,
    "hotel_id" INTEGER  REFERENCES hotels(id)ON DELETE CASCADE,
    "room_id" INTEGER  REFERENCES rooms(id)ON DELETE CASCADE,
    "date_first" date not NULL,
    "date_last" date not NULL,
    "adults_count" INTEGER NOT NULL default 0,
    "children_count" INTEGER NOT NULL default 0,
    "user_id" INTEGER REFERENCES users(id)ON DELETE CASCADE,
    "summa" decimal(18,2)
);
 
 
 
 
 
