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
