ALTER TABLE messages ADD COLUMN file_id TEXT;
ALTER TABLE messages ADD COLUMN log_index INTEGER;
ALTER TABLE messages ADD COLUMN log_term INTEGER;