CREATE TABLE "clients" (
  "id" uuid UNIQUE PRIMARY KEY,
  "name" varchar(255),
  "cpf" varchar(11),
  "email" varchar(255),
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE "product_categories" (
  "id" uuid UNIQUE PRIMARY KEY,
  "description" varchar(255)
);

CREATE TABLE "products" (
  "id" uuid UNIQUE PRIMARY KEY,
  "id_category" uuid,
  "description" varchar(255),
  "ingredients" text,
  "price" decimal(10,2),
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE "order_status" (
  "status" varchar(20) PRIMARY KEY,
  "description" varchar(255)
);

CREATE TABLE "orders" (
  "id" uuid UNIQUE PRIMARY KEY,
  "id_client" uuid,
  "status" varchar(20),
  "notification_attempts" int,
  "notified_at" timestamptz,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

CREATE TABLE "order_products" (
  "id_order" uuid,
  "id_product" uuid
);

CREATE TABLE "users" (
  "id" uuid UNIQUE PRIMARY KEY,
  "username" varchar(255),
  "password" text
);

-- ALTER TABLE "product_categories" ADD FOREIGN KEY ("id") REFERENCES "products" ("id_category");

-- ALTER TABLE "clients" ADD FOREIGN KEY ("id") REFERENCES "orders" ("id_client");

-- ALTER TABLE "order_status" ADD FOREIGN KEY ("status") REFERENCES "orders" ("status");

-- ALTER TABLE "orders" ADD FOREIGN KEY ("id") REFERENCES "order_products" ("id_order");

-- ALTER TABLE "products" ADD FOREIGN KEY ("id") REFERENCES "order_products" ("id_product");
