-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tenders (
            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
            name VARCHAR(100) NOT NULL,
            description VARCHAR(500) NOT NULL,
            serviceType service_type NOT NULL,
            status tender_status NOT NULL,
            organizationId UUID REFERENCES organization(id) ON DELETE CASCADE,
            creatorUsername VARCHAR(50) REFERENCES employee(username),
            version INT DEFAULT 1 NOT NULL,
            createdAt TIMESTAMP NOT NULL DEFAULT NOW()
        );
CREATE TABLE IF NOT EXISTS tender_history (
    id SERIAL PRIMARY KEY,
    tender_id UUID REFERENCES tenders(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    serviceType service_type NOT NULL, -- Предполагается, что тип service_type уже существует
    version INT NOT NULL,
    updated_at TIMESTAMP NOT NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tenders;
DROP TABLE IF EXISTS tender_history;
-- +goose StatementEnd
