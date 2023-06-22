BEGIN;

ALTER TABLE "smartduka_product" DROP FOREIGN KEY ("created_by") REFERENCES "smartduka_user" ("id");

ALTER TABLE "smartduka_product" DROP FOREIGN KEY ("updated_by") REFERENCES "smartduka_user" ("id");

ALTER TABLE "smartduka_sale" DROP FOREIGN KEY ("product_id") REFERENCES "smartduka_product" ("id");

ALTER TABLE "smartduka_sale" DROP FOREIGN KEY ("sold_by") REFERENCES "smartduka_user" ("id");

ALTER TABLE "smartduka_user_pin" DROP FOREIGN KEY ("user_id") REFERENCES "smartduka_user" ("id");

ALTER TABLE "smartduka_user_otp" DROP FOREIGN KEY ("user_id") REFERENCES "smartduka_user" ("id");

DROP TABLE IF EXISTS "smartduka_sale";

DROP TABLE IF EXISTS "smartduka_product";

DROP TABLE IF EXISTS "smartduka_user_pin";

DROP TABLE IF EXISTS "smartduka_user_otp";

DROP TABLE IF EXISTS "smartduka_user";

COMMIT;