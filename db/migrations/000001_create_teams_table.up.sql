BEGIN;

CREATE TABLE IF NOT EXISTS teams (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    members TEXT[] NOT NULL,
    image TEXT NOT NULL
);

COMMIT;