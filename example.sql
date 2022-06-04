drop table sessions;
drop table events;
drop table groups;
drop table users;

create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);

create table groups (
  id         serial primary key,
  name       text,
  host_id    integer references users(id),
  gest_id    integer references users(id),
  created_at timestamp not null
);

create table events (
  id         serial primary key,
  content    text,
  location   text,
  start_time timestamp not null,
  end_time   timestamp not null,
  host_id    integer references users(id),
  group_id   integer references groups(id),
  created_at timestamp not null
);
