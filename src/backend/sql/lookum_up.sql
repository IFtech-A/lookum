CREATE TABLE "countries" (
  "code" int PRIMARY KEY,
  "name" varchar,
  "continent_name" varchar
);

CREATE TABLE "address" (
  "id" SERIAL PRIMARY KEY,
  "default" bool NOT NULL DEFAULT false,
  "user_id" int NOT NULL,
  "line_1" varchar NOT NULL DEFAULT '',
  "line_2" varchar NOT NULL DEFAULT '',
  "city" varchar NOT NULL DEFAULT '',
  "province" varchar NOT NULL DEFAULT '',
  "country" varchar NOT NULL DEFAULT '',
  "created_at" varchar NOT NULL DEFAULT 'now()',
  "updated_at" varchar DEFAULT null
);

CREATE TABLE "transaction" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "order_id" int NOT NULL,
  "code" varchar NOT NULL,
  "type" int NOT NULL DEFAULT 0,
  "mode" int NOT NULL DEFAULT 0,
  "status" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp DEFAULT null,
  "content" text NOT NULL DEFAULT ''
);

CREATE TABLE "order_item" (
  "id" SERIAL PRIMARY KEY,
  "product_id" int NOT NULL,
  "order_id" int NOT NULL,
  "sku" varchar NOT NULL,
  "price" float4 NOT NULL DEFAULT 0,
  "discount" float4 NOT NULL DEFAULT 0,
  "quantity" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp DEFAULT null
);

CREATE TABLE "order" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "address_id" int NOT NULL,
  "token" varchar NOT NULL,
  "status" int NOT NULL DEFAULT 0,
  "sub_total" float8 NOT NULL DEFAULT 0,
  "item_discount" float8 NOT NULL DEFAULT 0,
  "tax" float4 NOT NULL DEFAULT 0,
  "shipping" float4 NOT NULL DEFAULT 0,
  "total" float8 NOT NULL DEFAULT 0,
  "promo" varchar NOT NULL DEFAULT '',
  "total_discount" float4 NOT NULL DEFAULT 0,
  "grand_total" float8 NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp DEFAULT null
);

CREATE TABLE "cart_item" (
  "id" SERIAL PRIMARY KEY,
  "product_id" int NOT NULL,
  "cart_id" int NOT NULL,
  "sku" varchar NOT NULL,
  "price" float4 NOT NULL DEFAULT 0,
  "discount" float4 NOT NULL DEFAULT 0,
  "quantity" int NOT NULL DEFAULT 0,
  "active" bool NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp DEFAULT null,
  "content" text NOT NULL DEFAULT ''
);

CREATE TABLE "cart" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "address_id" int NOT NULL,
  "token" varchar NOT NULL,
  "status" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp DEFAULT null,
  "content" text NOT NULL DEFAULT ''
);

CREATE TABLE "image" (
  "id" SERIAL PRIMARY KEY,
  "product_id" int,
  "filename" varchar NOT NULL DEFAULT '',
  "file_uri" varchar NOT NULL DEFAULT '',
  "main" bool NOT NULL DEFAULT false
);

CREATE TABLE "product_category" (
  "product_id" int NOT NULL,
  "category_id" int NOT NULL,
  PRIMARY KEY ("product_id", "category_id")
);

CREATE TABLE "category" (
  "id" SERIAL PRIMARY KEY,
  "parent_id" int DEFAULT null,
  "title" varchar NOT NULL,
  "meta_title" varchar NOT NULL DEFAULT '',
  "slug" varchar NOT NULL,
  "content" text NOT NULL DEFAULT ''
);

CREATE TABLE "product" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "title" varchar NOT NULL,
  "meta_title" varchar NOT NULL DEFAULT '',
  "slug" varchar NOT NULL,
  "summary" text NOT NULL DEFAULT '',
  "type" int NOT NULL DEFAULT 0,
  "sku" varchar NOT NULL DEFAULT '',
  "price" float4 NOT NULL DEFAULT 0,
  "discount" float4 NOT NULL DEFAULT 0,
  "quantity" int NOT NULL DEFAULT 0,
  "available" bool NOT NULL DEFAULT false,
  "content" text NOT NULL DEFAULT '',
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp DEFAULT 'now()',
  "published_at" timestamp DEFAULT null,
  "starts_at" timestamp DEFAULT null,
  "ends_at" timestamp DEFAULT null
);

CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar NOT NULL DEFAULT '',
  "middle_name" varchar NOT NULL DEFAULT '',
  "last_name" varchar NOT NULL DEFAULT '',
  "mobile" varchar NOT NULL DEFAULT '',
  "email" varchar NOT NULL DEFAULT '',
  "password" varchar NOT NULL DEFAULT '',
  "admin" bool NOT NULL DEFAULT false,
  "vendor" bool NOT NULL DEFAULT false,
  "registered_at" timestamp NOT NULL DEFAULT 'now()',
  "last_login" timestamp DEFAULT null,
  "intro" text NOT NULL DEFAULT '',
  "profile" text NOT NULL DEFAULT ''
);

ALTER TABLE "address" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "transaction" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "order_item" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "order_item" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "cart_item" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "cart_item" ADD FOREIGN KEY ("cart_id") REFERENCES "cart" ("id");

ALTER TABLE "cart" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "cart" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");

ALTER TABLE "image" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_category" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_category" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "category" ADD FOREIGN KEY ("parent_id") REFERENCES "category" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

CREATE INDEX "idx_address_user" ON "address" ("user_id");

CREATE INDEX "idx_address_user_default" ON "address" ("user_id", "default");

CREATE INDEX "idx_transaction_user" ON "transaction" ("user_id");

CREATE INDEX "idx_transaction_order" ON "transaction" ("order_id");

CREATE INDEX "idx_order_item_product" ON "order_item" ("product_id");

CREATE INDEX "idx_order_item_order" ON "order_item" ("order_id");

CREATE INDEX "idx_order_user" ON "order" ("user_id");

CREATE INDEX "idx_cart_item_product" ON "cart_item" ("product_id");

CREATE INDEX "idx_cart_item_cart" ON "cart_item" ("cart_id");

CREATE INDEX "idx_cart_user" ON "cart" ("user_id");

CREATE INDEX "idx_image_product" ON "image" ("product_id");

CREATE INDEX "idx_pc_category" ON "product_category" ("category_id");

CREATE INDEX "idx_pc_product" ON "product_category" ("product_id");

CREATE INDEX "idx_category_parent" ON "category" ("parent_id");

CREATE UNIQUE INDEX "uq_slug" ON "product" ("slug");

CREATE INDEX "idx_product_user" ON "product" ("user_id");

CREATE UNIQUE INDEX "uq_mobile" ON "user" ("mobile");

CREATE UNIQUE INDEX "uq_email" ON "user" ("email");

COMMENT ON COLUMN "transaction"."code" IS 'The payment id provided by the payment gateway';

COMMENT ON COLUMN "transaction"."type" IS 'The type of order transaction can be either Credit or Debit, etc.';

COMMENT ON COLUMN "transaction"."mode" IS 'The mode of the order transaction can be Offline, Cash On Delivery, Cheque, Draft, Wired, and Online';

COMMENT ON COLUMN "transaction"."status" IS 'The status of the order transaction can be New, Cancelled, Failed, Pending, Declined, Rejected, and Success';

COMMENT ON COLUMN "order"."sub_total" IS 'The total price of the Order Items';

COMMENT ON COLUMN "order"."item_discount" IS 'The total discount of the Order Items';

COMMENT ON COLUMN "order"."tax" IS 'The tax on the Order Items';

COMMENT ON COLUMN "order"."shipping" IS 'The shipping charges of the Order Items';

COMMENT ON COLUMN "order"."total" IS 'The total price of the Order including tax and shipping. It excludes the items discount';

COMMENT ON COLUMN "order"."promo" IS 'The Promo code of the Order';

COMMENT ON COLUMN "order"."total_discount" IS 'The total discount of the Order based on the promo code or store discount';

COMMENT ON COLUMN "order"."grand_total" IS 'The grand total of the order to be paid by the buyer';
