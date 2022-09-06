-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

create table if not exists public.geolocations
(
    id            serial
        constraint geolocations_pk primary key,
    ip_address    inet        not null,
    country_code  varchar,
    country       varchar,
    city          varchar,
    coordinates   point       not null,
    mystery_value varchar,
    created_at    timestamptz DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

drop table if exists public.geolocations;
-- +goose StatementEnd
