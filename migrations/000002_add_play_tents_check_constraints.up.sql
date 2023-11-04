ALTER TABLE play_tents ADD CONSTRAINT play_tents_weight_check CHECK (weight > 0);
ALTER TABLE play_tents ADD CONSTRAINT play_tents_color_check CHECK (length(color) > 0 AND length(color) < 64);