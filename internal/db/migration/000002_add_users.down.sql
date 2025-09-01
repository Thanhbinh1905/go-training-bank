ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

DROP INDEX IF EXISTS "accounts_owner_idx";

ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

DROP INDEX IF EXISTS "user_email_idx";

DROP TABLE IF EXISTS "users";
