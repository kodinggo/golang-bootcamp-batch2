
-- +migrate Up
ALTER TABLE articles
ADD category_id INT NULL;

-- +migrate Down
ALTER TABLE articles
DROP COLUMN category_id;