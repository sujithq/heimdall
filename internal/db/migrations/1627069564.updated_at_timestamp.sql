-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION trigger_set_updated_at();
