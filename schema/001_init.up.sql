create table if not exists commercials (
    id integer primary key,
    bank text,
    usd_in float,
    usd_out float,
    euro_in float,
    euro_out float,
    rub_in float,
    rub_out float,
    requested_at datetime not null default current_timestamp
);