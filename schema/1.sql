create table if not exists accounts (
  id varchar(128) primary key,
  fullname text,
  email text not null unique,
  username text
);

create table if not exists tokens (
  token varchar(128) primary key,
  account_id varchar(128)
);
