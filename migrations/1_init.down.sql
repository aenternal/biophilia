CREATE TABLE IF NOT EXISTS biomolecule (
                             id SERIAL PRIMARY KEY,
                             type VARCHAR(255),
                             name VARCHAR(255),
                             sequence TEXT,
                             description TEXT,
                             created_at TIMESTAMP DEFAULT NOW(),
                             updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS biomolecule_files (
                                  id SERIAL PRIMARY KEY,
                                  biomolecule_id INT REFERENCES biomolecule(id),
                                  file_url TEXT,
                                  file_type VARCHAR(50),
                                  uploaded_at TIMESTAMP DEFAULT NOW()
);