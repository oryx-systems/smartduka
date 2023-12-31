BEGIN;

CREATE TABLE IF NOT EXISTS "smartduka_contact" (
  "id" uuid PRIMARY KEY NOT NULL,
  "active" boolean NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "contact_type" varchar(16) NOT NULL,
  "contact_value" text NOT NULL,
  "flavour" text NOT NULL,
  "user_id" uuid NOT NULL
);

ALTER TABLE "smartduka_contact" ADD FOREIGN KEY ("user_id") REFERENCES "smartduka_user" ("id");

COMMIT;