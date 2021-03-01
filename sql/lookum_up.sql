
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'products_status') THEN
		CREATE TYPE "products_status" AS ENUM (
  			'out_of_stock',
  			'in_stock',
  			'running_low'
		);
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS "users" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar DEFAULT '',
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp DEFAULT now(),
  "country_code" int DEFAULT 0
);

CREATE TABLE IF NOT EXISTS  "countries" (
  "code" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "continent_name" varchar NOT NULL
);

CREATE TABLE  IF NOT EXISTS "images" (
  "id" SERIAL PRIMARY KEY,
  "product_id" int NOT NULL,
  "filename" varchar NOT NULL,
  "file_uri" varchar NOT NULL
);

CREATE TABLE  IF NOT EXISTS "order_items" (
  "order_id" int NOT NULL,
  "product_id" int NOT NULL,
  "quantity" int DEFAULT 1
);

CREATE TABLE  IF NOT EXISTS "orders" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int UNIQUE NOT NULL,
  "status" varchar DEFAULT '',
  "created_at" varchar DEFAULT ''
);

CREATE TABLE  IF NOT EXISTS "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar DEFAULT '',
  "description" varchar DEFAULT ''
);

CREATE TABLE  IF NOT EXISTS "products" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "desc" varchar,
  "price" float4 DEFAULT 0,
  "discount" float4 DEFAULT 0,
  "likes" int DEFAULT 0,
  "status" products_status DEFAULT 'out_of_stock',
  "created_at" timestamp DEFAULT (now()),
  "category_id" int DEFAULT 0
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'images_product_id_fkey') THEN
        ALTER TABLE images
            ADD CONSTRAINT images_product_id_fkey
            FOREIGN KEY ("product_id") REFERENCES "products" ("id");
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'users_country_code_fkey') THEN
        ALTER TABLE users
            ADD CONSTRAINT users_country_code_fkey
            FOREIGN KEY ("country_code") REFERENCES "countries" ("code");
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'order_items_order_id_fkey') THEN
        ALTER TABLE order_items
            ADD CONSTRAINT order_items_order_id_fkey
            FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'order_items_product_id_fkey') THEN
        ALTER TABLE order_items
            ADD CONSTRAINT order_items_product_id_fkey
            FOREIGN KEY ("product_id") REFERENCES "products" ("id");
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'products_category_id_fkey') THEN
        ALTER TABLE products
            ADD CONSTRAINT products_category_id_fkey
            FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
    END IF;
END;
$$;

CREATE INDEX  IF NOT EXISTS "users_email" ON "users" ("email");

CREATE INDEX  IF NOT EXISTS "images_product_id" ON "images" ("product_id");

CREATE INDEX  IF NOT EXISTS "order_items_order_id" ON "order_items" ("order_id");

CREATE INDEX  IF NOT EXISTS "order_items_product_id" ON "order_items" ("product_id");

CREATE INDEX  IF NOT EXISTS "orders_status" ON "orders" ("status");

CREATE INDEX  IF NOT EXISTS "orders_user_id" ON "orders" ("user_id");

CREATE UNIQUE INDEX  IF NOT EXISTS "products_id" ON "products" ("id");

CREATE INDEX  IF NOT EXISTS "products_category_id" ON "products" ("category_id");

CREATE INDEX  IF NOT EXISTS "products_created_at" ON "products" ("created_at");

CREATE INDEX  IF NOT EXISTS "products_price" ON "products" ("price");

CREATE INDEX  IF NOT EXISTS "products_discount" ON "products" ("discount");

COMMENT ON COLUMN "orders"."created_at" IS 'When order created';
