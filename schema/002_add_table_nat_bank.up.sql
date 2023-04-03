create table if not exists nat_bank (
    id integer primary key,
    nat_id int,
    abbreviation text,
    name text,
    scale int,
    official_rate float,
    requested_at datetime not null default current_timestamp
);