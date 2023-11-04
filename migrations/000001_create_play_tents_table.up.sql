CREATE TABLE IF NOT EXISTS play_tents (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    description text NOT NULL,
    color text NOT NULL,
    material text NOT NULL,
    weight numeric(6,3) NOT NULL,
    size text NOT NULL,
    version integer NOT NULL DEFAULT 1
);