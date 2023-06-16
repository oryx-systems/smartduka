BEGIN;

ALTER TABLE "product" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("updated_by") REFERENCES "user" ("id");

ALTER TABLE "sale" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "sale" ADD FOREIGN KEY ("sold_by") REFERENCES "user" ("id");

ALTER TABLE "user_pin" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_otp" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

DROP TABLE IF EXISTS "sale";

DROP TABLE IF EXISTS "product";

DROP TABLE IF EXISTS "user_pin";

DROP TABLE IF EXISTS "user_otp";

DROP TABLE IF EXISTS "user";

COMMIT;