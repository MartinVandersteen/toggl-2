CREATE TABLE IF NOT EXISTS cards(
   id uuid DEFAULT uuid_generate_v4(),
   code VARCHAR(25),
   value VARCHAR(25),
   suit VARCHAR(25)
);
