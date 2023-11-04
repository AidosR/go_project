CREATE INDEX IF NOT EXISTS play_tents_title_idx ON play_tents USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS play_tents_color_idx ON play_tents USING GIN (to_tsvector('simple', color));
CREATE INDEX IF NOT EXISTS play_tents_material_idx ON play_tents USING GIN (to_tsvector('simple', material));