-- PostgreSQL 스키마

create table users
(
  id serial primary key,
  name varchar(255) not null,
  github_id varchar(255) unique,

  edited_time timestamp,
  created_time timestamp default CURRENT_TIMESTAMP
);

create table links
(
  id serial primary key,

  url text,
  tags character varying[],
  comment text,
  user_id int references users(id) not null,

  edited_time timestamp,
  created_time timestamp default CURRENT_TIMESTAMP
);

