CREATE TABLE IF NOT EXISTS decks(
   id uuid DEFAULT uuid_generate_v4(),
   shuffled boolean NOT NULL,
   remaining NUMERIC NOT NULL DEFAULT 0,
   cardset_id uuid REFERENCES cardsets,
   cards VARCHAR(25)[]
);
