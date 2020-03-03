CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS cardsets(
   id uuid DEFAULT uuid_generate_v4(),
   name VARCHAR (50) UNIQUE NOT NULL
);
