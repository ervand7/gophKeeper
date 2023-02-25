-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table if not exists "user"
(
    "id"       uuid default gen_random_uuid() primary key,
    "login"    varchar(100)  not null
        unique,
    "password" varchar(100)  not null,
    "token"    varchar(1000) not null
        unique
);

create table if not exists "credentials"
(
    "id"         uuid                                 default gen_random_uuid() primary key,
    "user_id"    uuid                        not null,
    "title"      varchar(1000)               not null,
    "login"      varchar(1000)               not null,
    "password"   varchar(1000)               not null,
    "created_at" timestamp without time zone not null default now(),
    "updated_at" timestamp without time zone not null default now(),
    "meta"       varchar(1000),

    unique ("user_id", "title")
);
alter table only "public"."credentials"
    add constraint "fk_user_id" foreign key ("user_id")
        references "public"."user" ("id") on delete cascade;

create table if not exists "text"
(
    "id"         uuid                                 default gen_random_uuid() primary key,
    "user_id"    uuid                        not null,
    "title"      varchar(1000)               not null,
    "content"    varchar(1000)               not null,
    "created_at" timestamp without time zone not null default now(),
    "updated_at" timestamp without time zone not null default now(),
    "meta"       varchar(1000),

    unique ("user_id", "title")
);
alter table only "public"."text"
    add constraint "fk_user_id" foreign key ("user_id")
        references "public"."user" ("id") on delete cascade;

create table if not exists "binary_data"
(
    "id"         uuid                                 default gen_random_uuid() primary key,
    "user_id"    uuid                        not null,
    "title"      varchar(1000)               not null,
    "content"    bytea                       not null,
    "created_at" timestamp without time zone not null default now(),
    "updated_at" timestamp without time zone not null default now(),
    "meta"       varchar(1000),

    unique ("user_id", "title")
);
alter table only "public"."binary_data"
    add constraint "fk_user_id" foreign key ("user_id")
        references "public"."user" ("id") on delete cascade;

create table if not exists "bank_card"
(
    "id"          uuid                                 default gen_random_uuid() primary key,
    "user_id"     uuid                        not null,
    "title"       varchar(1000)               not null,
    "card_holder" varchar(1000)               not null,
    "card_number" varchar(1000)               not null,
    "card_expire" varchar(1000)               not null,
    "card_cvv"    varchar(1000)               not null,
    "created_at"  timestamp without time zone not null default now(),
    "updated_at"  timestamp without time zone not null default now(),
    "meta"        varchar(1000),

    unique ("user_id", "title")
);
alter table only "public"."bank_card"
    add constraint "fk_user_id" foreign key ("user_id")
        references "public"."user" ("id") on delete cascade;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists "user" cascade;
drop table if exists "credentials" cascade;
drop table if exists "text" cascade;
drop table if exists "binary_data" cascade;
drop table if exists "bank_card" cascade;
