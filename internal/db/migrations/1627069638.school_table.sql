-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "school" (
   "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
   "name" TEXT NOT NULL
);

CREATE TRIGGER "set_updated_at"
    BEFORE UPDATE ON "school"
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

-- +migrate Down
DROP TRIGGER "set_updated_at" ON "school";
DROP TABLE "school";
