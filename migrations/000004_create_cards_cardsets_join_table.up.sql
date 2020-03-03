CREATE TABLE IF NOT EXISTS cards_cardsets(
   cardset_id uuid REFERENCES cardsets,
   card_id uuid REFERENCES cards
);
