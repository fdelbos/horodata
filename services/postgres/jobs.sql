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
    creator_id bigint not null references users on delete restrict,
    duration bigint not null,
    comment text,
    updated timestamp,
    updater_id bigint references users on delete restrict
);
