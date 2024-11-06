-- +goose Up
-- +goose StatementBegin
DO $do$
BEGIN
    -- Проверка, существует ли тип tender_status
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_status') THEN
CREATE TYPE tender_status AS ENUM ('Created', 'Published', 'Closed');
END IF;

    -- Проверка, существует ли тип service_type
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'service_type') THEN
CREATE TYPE service_type AS ENUM ('Construction', 'Delivery', 'Manufacture');
END IF;
END $do$;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DO $do$
BEGIN
    -- Удаление типов (если они не используются)
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tender_status') THEN
DROP TYPE tender_status;
END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'service_type') THEN
DROP TYPE service_type;
END IF;
END $do$;
-- +goose StatementEnd
