CREATE TABLE if not exists "booking_table"(
    "id" serial PRIMARY KEY,
    "room_id" INTEGER NOT NULL REFERENCES rooms(id)ON DELETE CASCADE,
    "from" TIMESTAMP WITH TIME ZONE,
    "to" TIMESTAMP WITH TIME ZONE
);
