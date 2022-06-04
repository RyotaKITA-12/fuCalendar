drop table sessions;
drop table users;
drop table todos;
drop table events;

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

create table todos (
  id         serial primary key,
  content    text,
  user_id    integer references users(id),
  created_at timestamp not null
);

create table events (
  id         serial primary key,
  content    text,
  location   text,
  user_id    text,
  start_time timestamp not null,
  end_time   timestamp not null,
  created_at timestamp not null
)

create table invitations(
  id         serial primary key,
  event_id   integer references events(id),
  host_id    integer references users(id),
  gest_id    integer references users(id),
  created_at timestamp not null
);
