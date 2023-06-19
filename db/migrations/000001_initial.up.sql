BEGIN;

CREATE TABLE "user" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "username" varchar(20),
  "first_name" varchar(25),
  "last_name" varchar(25),
  "email" varchar(100),
  "user_type" varchar(20),
  "push_token" varchar(200)
);

CREATE TABLE "product" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_at" timestamp,
  "updated_by" uuid NOT NULL,
  "deleted_at" timestamp,
  "active" boolean NOT NULL,
  "name" varchar(50) NOT NULL,
  "category" varchar(25) NOT NULL,
  "quantity" varchar(10) NOT NULL,
  "unit" varchar(15) NOT NULL,
  "price" float NOT NULL,
  "description" text,
  "manufacturer" varchar(50),
  "in_stock" boolean NOT NULL
);

CREATE TABLE "sale" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" uuid NOT NULL,
  "updated_at" timestamp,
  "active" boolean NOT NULL,
  "product_id" uuid NOT NULL,
  "quantity" varchar(50) NOT NULL,
  "unit" varchar(50) NOT NULL,
  "price" float NOT NULL,
  "sold_by" uuid NOT NULL
);

CREATE TABLE "user_pin" (
  "id" uuid PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "created_by" uuid,
  "updated_at" timestamp,
  "updated_by" uuid,
  "active" boolean NOT NULL,
  "valid_from" timestamp NOT NULL,
  "valid_to" timestamp NOT NULL,
  "hashed_pin" text NOT NULL,
  "salt" text NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "user_otp" (
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

ALTER TABLE "product" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("updated_by") REFERENCES "user" ("id");

ALTER TABLE "sale" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "sale" ADD FOREIGN KEY ("sold_by") REFERENCES "user" ("id");

ALTER TABLE "user_pin" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_otp" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

COMMIT;