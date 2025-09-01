-- Drop constraints first
ALTER TABLE IF EXISTS "entries" DROP CONSTRAINT IF EXISTS entries_account_id_fkey;
ALTER TABLE IF EXISTS "transfers" DROP CONSTRAINT IF EXISTS transfers_from_account_id_fkey;
ALTER TABLE IF EXISTS "transfers" DROP CONSTRAINT IF EXISTS transfers_to_account_id_fkey;

-- Drop indexes
DROP INDEX IF EXISTS accounts_owner_idx;
DROP INDEX IF EXISTS entries_account_id_idx;
DROP INDEX IF EXISTS transfers_from_account_id_idx;
DROP INDEX IF EXISTS transfers_to_account_id_idx;
DROP INDEX IF EXISTS transfers_from_account_id_to_account_id_idx;

-- Drop comments
COMMENT ON COLUMN "entries"."amount" IS NULL;
COMMENT ON COLUMN "transfers"."amount" IS NULL;

-- Drop tables in reverse order
DROP TABLE IF EXISTS "transfers";
DROP TABLE IF EXISTS "entries";
DROP TABLE IF EXISTS "accounts";
