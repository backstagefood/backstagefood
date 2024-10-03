CREATE TABLE "customers"
(
    "id"         uuid UNIQUE PRIMARY KEY,
    "name"       varchar(255),
    "cpf"        varchar(11),
    "email"      varchar(255),
    "created_at" timestamptz default now(),
    "updated_at" timestamptz default now()
);

CREATE TABLE "product_categories"
(
    "id"          uuid UNIQUE PRIMARY KEY,
    "description" varchar(255)
);

CREATE TABLE "products"
(
    "id"          uuid UNIQUE PRIMARY KEY,
    "id_category" uuid references "product_categories" ("id"),
    "description" varchar(255),
    "ingredients" text,
    "price"       decimal(10, 2),
    "created_at"  timestamptz default now(),
    "updated_at"  timestamptz default now()
);

CREATE TABLE "order_status"
(
    "status"      varchar(20) PRIMARY KEY,
    "description" varchar(255)
);

CREATE TABLE "orders"
(
    "id"                    uuid UNIQUE PRIMARY KEY,
    "id_customer"             uuid references "customers" ("id"),
    "status"                varchar(20) references "order_status" ("status"),
    "notification_attempts" int,
    "notified_at"           timestamptz,
    "created_at"            timestamptz default now(),
    "updated_at"            timestamptz default now()
);

CREATE TABLE "order_products"
(
    "id_order"   uuid references "orders" ("id"),
    "id_product" uuid references "products" ("id")
);

CREATE TABLE "users"
(
    "id"       uuid UNIQUE PRIMARY KEY,
    "username" varchar(255),
    "password" text
);
