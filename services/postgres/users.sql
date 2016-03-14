--
-- Users
--

create type user_role as enum ('admin', 'help');

create table users (
    id bigserial primary key,
    created timestamp default now()  not null,
    active boolean default true not null,
    login varchar(30) unique not null,
    email citext unique not null,
    full_name varchar(100),
    organization varchar(100),
    website text,
    about text,

    hash bytea,
    hash_version integer,
    admin boolean default false not null,
    role varchar(30) default null
);

create view users_view as
    select
        id, created, active, login, email, full_name, organization, website, about
    from users;

create view users_active as
    select * from users_view where active = true;

create table sessions (
    id bigserial primary key,
    created timestamp default now() not null,
    user_id bigint not null references users on delete cascade,
    active boolean default true  not null,
    host text
);

create table password_requests (
    id bigserial primary key,
    created timestamp default now()  not null,
    user_id bigint not null references users on delete cascade,
    active boolean default true  not null,
    url varchar(40) unique not null
);

create type quota_plan as enum ('free', 'small', 'medium', 'large', 'custom');

create table quotas (
    user_id bigint primary key references users on delete cascade,
    created timestamp default now()  not null,
    plan quota_plan default 'free'  not null
);

create table quotas_custom (
    user_id bigint primary key  references users on delete cascade,
    created timestamp default now()  not null,
    instances bigint default 0  not null,
    forms bigint default 0  not null,
    roles bigint default 0  not null,
    files bigint default 0  not null
);

create table quotas_bonus (
    id bigserial primary key,
    created timestamp default now()  not null,
    user_id bigint not null references users on delete cascade,
    description text not null,
    instances bigint default 0  not null,
    forms bigint default 0  not null,
    roles bigint default 0  not null,
    files bigint default 0  not null
);

create table usages (
    user_id bigint primary key references users on delete cascade,
    created timestamp default now() not null,
    instances bigint default 0  not null,
    forms bigint default 0  not null,
    roles bigint default 0  not null,
    files bigint default 0  not null
);

create function user_new(login varchar, email citext)
returns void as $$
declare
  user_id bigint;
begin
    insert into users (login, email)
    values (login, email);

    user_id := lastval();

    insert into quotas (user_id) values (user_id);

    insert into usages (user_id) values (user_id);
end;
$$ language plpgsql;
