BEGIN;

CREATE TABLE IF NOT EXISTS "contact" (
  "id" uuid PRIMARY KEY NOT NULL,
  "active" boolean NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "contact_type" varchar(16) NOT NULL,
  "contact_value" text NOT NULL,
  "user_id" uuid NOT NULL
);

ALTER TABLE "contact" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

COMMIT;