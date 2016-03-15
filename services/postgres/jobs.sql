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

create table jobs (
    id bigserial primary key,
    created timestamp default now() not null,
    group_id bigint not null references groups on delete cascade,
    task_id bigint not null references tasks on delete cascade,
    customer_id bigint not null references customers on delete cascade,
    user_id bigint not null references users on delete cascade,
    day date default now() not null,
    duration integer not null,
    comment text
);

create table guests (
    id bigserial primary key,
    created timestamp default now() not null,
    active boolean default true not null,
    group_id bigint not null references groups on delete cascade,
    user_id bigint not null references users on delete cascade,
    rate int default 0 not null,
    admin boolean default false not null,
    email citext unique not null,
    message text,
    unique(group_id, user_id)
);
