-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER update_tender_version
BEFORE UPDATE ON tenders
FOR EACH ROW
WHEN (OLD.name IS DISTINCT FROM NEW.name OR
      OLD.description IS DISTINCT FROM NEW.description OR
      OLD.serviceType IS DISTINCT FROM NEW.serviceType)
EXECUTE FUNCTION save_tender_version();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_tender_version ON tender;
-- +goose StatementEnd
