BEGIN;

CREATE TABLE "smartduka_user" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "username" varchar(20),
  "first_name" varchar(25),
  "last_name" varchar(25),
  "email" varchar(100),
  "user_type" varchar(20),
  "push_token" varchar(200)
);

CREATE TABLE "smartduka_product" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "name" varchar(50) NOT NULL,
  "category" varchar(25) NOT NULL,
  "quantity" float NOT NULL,
  "unit" varchar(15) NOT NULL,
  "price" float NOT NULL,
  "vat" float NOT NULL,
  "description" text,
  "manufacturer" varchar(50),
  "in_stock" boolean NOT NULL
);

CREATE TABLE "smartduka_sale" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "product_id" uuid NOT NULL,
  "quantity" float NOT NULL,
  "unit" varchar(50) NOT NULL,
  "price" float NOT NULL
);

CREATE TABLE "smartduka_user_pin" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "valid_from" timestamp NOT NULL,
  "valid_to" timestamp NOT NULL,
  "hashed_pin" text NOT NULL,
  "salt" text NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "smartduka_user_otp" (
  "id" uuid PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "is_valid" boolean NOT NULL,
  "valid_until" timestamp NOT NULL,
  "phone_number" varchar(20) NOT NULL,
  "otp" varchar(10) NOT NULL,
  "flavour" varchar(10) NOT NULL,
  "user_id" uuid NOT NULL
);

ALTER TABLE "smartduka_product" ADD FOREIGN KEY ("created_by") REFERENCES "smartduka_user" ("id");

ALTER TABLE "smartduka_product" ADD FOREIGN KEY ("updated_by") REFERENCES "smartduka_user" ("id");

ALTER TABLE "smartduka_sale" ADD FOREIGN KEY ("product_id") REFERENCES "smartduka_product" ("id");

ALTER TABLE "smartduka_sale" ADD FOREIGN KEY ("created_by") REFERENCES "smartduka_user" ("id");

ALTER TABLE "smartduka_user_pin" ADD FOREIGN KEY ("user_id") REFERENCES "smartduka_user" ("id");

ALTER TABLE "smartduka_user_otp" ADD FOREIGN KEY ("user_id") REFERENCES "smartduka_user" ("id");

COMMIT;