create table if not exists nacbank (
    id integer primary key,
    cur_id bigint,
    abbreviation text,
    name text,
    scale int,
    rate float
);