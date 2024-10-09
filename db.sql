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

-- insere registros de testes
INSERT INTO product_categories values('83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'bebidas');
INSERT INTO public.products (id, id_category, description, ingredients, price, created_at, updated_at)
VALUES(gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'cerveja', 'agua, malte, lupulo e levedura', 10.99, now(), now()),
      (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'coca-cola', '?', 9.75, now(), now());