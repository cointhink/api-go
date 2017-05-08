create table if not exists accounts (
  id varchar(128) primary key,
  fullname text,
  email text not null unique,
  username text unique
);

