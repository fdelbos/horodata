
create extension citext;

--
-- Users
--

create table users (
    id bigserial primary key,
    created timestamp default now()  not null,
    active boolean default true not null,
    email citext unique not null,
    full_name varchar(50) not null,
    hash bytea,
    hash_version integer
);

create view users_view as
    select
        id, created, active, email, full_name
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

create function user_new(email citext, full_name text)
returns void as $$
declare
  user_id bigint;
begin
    insert into users (email, full_name) values (email, full_name);

    user_id := lastval();

    insert into quotas (user_id) values (user_id);

    insert into usages (user_id) values (user_id);
end;
$$ language plpgsql;


--
-- Billing
--

create table addresses (
    id bigserial primary key,
    created timestamp default now()  not null,
    user_id bigint not null references users on delete cascade,
    full_name varchar(100) not null,
    email citext not null,
    company varchar(100),
    vat varchar(100),
    address1 text not null,
    address2 text,
    city varchar(150) not null,
    zip varchar(15) not null
);

create table address_current (
    user_id bigint primary key references users on delete cascade,
    address_id bigint references addresses on delete cascade
);

create function addresses_keep_current() returns trigger as  $keep_current$
begin

    insert into
        address_current (user_id, address_id)
    values
        (new.user_id, new.id)
    on conflict (user_id) do update set
        address_id = new.id;

    delete from addresses
    where
        id not in (
            select address_id
            from address_current
            where user_id = new.user_id)
        and id not in (
            select address_id
            from invoices
            where
                user_id = new.user_id
                and address_id = new.id);

    return new;
end;
$keep_current$ language plpgsql;

create trigger addresses_keep_current_trigger
after insert on addresses for each row execute procedure addresses_keep_current();

create table subscriber (
    user_id bigint primary key references users on delete cascade,
    created timestamp default now()  not null,
    stripe_id text not null
);

create table cards (
    user_id bigint primary key references users on delete cascade,
    created timestamp default now()  not null,
    stripe_id text not null,
    last4 char(4) not null,
    brand text not null,
    expiration date not null
);

create table stripe_subscriptions (
    user_id bigint primary key references users on delete cascade,
    stripe_id text not null,
    active boolean default true not null,
    plan quota_plan not null,
    tax_percent int,
    end_date timestamp
);

create table invoices (
    id bigserial primary key,
    created timestamp default now()  not null,
    user_id bigint references users on delete restrict,
    address_id bigint references addresses on delete restrict,
    plan quota_plan not null,
    start_date timestamp not null,
    end_date timestamp not null,
    subtotal bigint not null,
    total bigint not null,
    tax bigint not null,
    tax_percent numeric(4 ,2),
    paid bool default false not null,
    sent bool default false not null,
    charge text
);

create table invoice_items (
    id bigserial primary key,
    invoice_id  bigint references users on delete restrict,
    amount bigint not null,
    unit_price bigint not null,
    quantity bigint not null,
    start_date timestamp not null,
    end_date timestamp not null,
    description text,
    title text
);


--
-- Jobs
--

create table groups (
    id bigserial primary key,
    created timestamp default now() not null,
    active boolean default true not null,
    owner_id bigint not null references users on delete cascade,
    name varchar(30) not null,
    url varchar(200) unique not null
);

create table tasks (
    id bigserial primary key,
    created timestamp default now() not null,
    active boolean default true not null,
    group_id bigint not null references groups on delete cascade,
    name citext not null,
    comment_mandatory boolean default false not null,
    unique(group_id, name)
);

create table customers (
    id bigserial primary key,
    created timestamp default now() not null,
    active boolean default true not null,
    group_id bigint not null references groups on delete cascade,
    name citext not null,
    unique(group_id, name)
);

create table guests (
    id bigserial primary key,
    created timestamp default now() not null,
    active boolean default true not null,
    group_id bigint not null references groups on delete cascade,
    user_id bigint,
    rate int default 0 not null,
    admin boolean default false not null,
    email citext not null,
    unique(group_id, user_id)
);

create table jobs (
    id bigserial primary key,
    created timestamp default now() not null,
    group_id bigint not null references groups on delete cascade,
    task_id bigint not null references tasks on delete restrict,
    customer_id bigint not null references customers on delete restrict,
    creator_id bigint not null references guests on delete restrict,
    duration bigint not null,
    comment text,
    updated timestamp,
    updater_id bigint references users on delete restrict
);
