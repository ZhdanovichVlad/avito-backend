-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION save_tender_version() RETURNS TRIGGER AS $$
BEGIN
INSERT INTO tender_history (tender_id, name, description,serviceType, version, updated_at)
VALUES (OLD.id, OLD.name, OLD.description, OLD.serviceType, OLD.version, NOW());

NEW.version := OLD.version + 1;
    NEW.updated_at := NOW();

RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS save_tender_version();
-- +goose StatementEnd
