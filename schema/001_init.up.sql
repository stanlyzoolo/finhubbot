create table if not exists commercials (
    id integer primary key,
    bank text,
    usd_in float,
    usd_out float,
    euro_in float,
    euro_out float,
    rub_in float,
    rub_out float,
    conv_usd_to_euro_in float,
    conv_usd_to_euro_out float,
    date datetime
);