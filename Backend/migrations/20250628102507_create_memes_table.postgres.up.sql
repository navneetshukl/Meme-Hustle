DROP TABLE IF EXISTS memes;

CREATE TABLE memes (
    meme_id SERIAL,
    title TEXT NOT NULL,
    image_url VARCHAR(255) NOT NULL DEFAULT 'https://picsum.photos/200',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON memes
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

-- Insert
INSERT INTO memes(title) VALUES('Best Meme');