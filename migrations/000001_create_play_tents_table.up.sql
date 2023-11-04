CREATE TABLE IF NOT EXISTS play_tents (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL DEFAULT 'PLAY TENT DEFAULT',
    description text NOT NULL,
    color text NOT NULL DEFAULT 'NONE',
    material text NOT NULL DEFAULT 'NONE',
    weight numeric(6,3) NOT NULL,
    size text NOT NULL,
    version integer NOT NULL DEFAULT 1
);