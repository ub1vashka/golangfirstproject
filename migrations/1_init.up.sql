CREATE TABLE IF NOT EXISTS users (
  uid varchar(36) NOT NULL PRIMARY KEY,
  name test NOT NULL,
  email test NOT NULL,
  pass  test NOT NULL,
  age INTEGER,
  RegisteredAt DATATIME
);

CREATE TABLE IF NOT EXISTS books (
  bid varchar(36) NOT NULL PRIMARY KEY,
  label test NOT NULL,
  author test NOT NULL,
  desc  test NOT NULL,
  WrittenAt DATATIME
);