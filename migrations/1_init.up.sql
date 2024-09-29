DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'biomolecule_type') THEN
        CREATE TYPE biomolecule_type AS ENUM ('rna', 'dna', 'protein');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS biomolecule (
    id SERIAL PRIMARY KEY,
    type biomolecule_type NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    sequence TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON biomolecule
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TABLE IF NOT EXISTS biomolecule_files (
    id SERIAL PRIMARY KEY,
    biomolecule_id INT REFERENCES biomolecule(id) ON DELETE CASCADE,
    file_name TEXT NOT NULL,
    file_type VARCHAR(50),
    uploaded_at TIMESTAMP DEFAULT NOW()
);