-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION restore_tender_to_version(tender_id UUID, version_to_restore INT) RETURNS VOID AS $$
DECLARE
target_record RECORD;
BEGIN
SELECT * INTO target_record
FROM tender_history
WHERE tender_id = tender_id AND version = version_to_restore
    LIMIT 1;

IF target_record IS NOT NULL THEN
UPDATE tenders
SET name = target_record.name,
    description = target_record.description,
    serviceType = target_record.serviceType,
    version = version + 1,
    updated_at = NOW()
WHERE id = tender_id;

ELSE
        RAISE NOTICE 'Запись с указанной версией % для tender_id % не найдена', version_to_restore, tender_id;
END IF;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS restore_tender_to_version();
-- +goose StatementEnd
