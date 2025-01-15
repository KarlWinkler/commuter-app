CREATE TABLE mode (
  id   BIGSERIAL PRIMARY KEY,
  name text
);

CREATE TABLE trip (
  id       BIGSERIAL PRIMARY KEY,
  name     text,
  distance int NOT NULL,
  mode     int,
  FOREIGN KEY (mode) REFERENCES mode (id)
);
