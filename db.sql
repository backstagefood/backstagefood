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
INSERT INTO product_categories VALUES
    ('83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Lanche'),
    ('83a7f6c5-b717-4159-b2b2-675bccb5cd82', 'Acompanhamento'),
    ('83a7f6c5-b717-4159-b2b2-675bccb5cd83', 'Bebida'),
    ('83a7f6c5-b717-4159-b2b2-675bccb5cd84', 'Sobremesa');
INSERT INTO public.products (id, id_category, description, ingredients, price, created_at, updated_at) VALUES
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Hamburguer simples', 'Pão e hamburguer', 7.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Cheese burguer', 'Pão, hamburguer e queijo', 9.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Cheese egg burguer', 'Pão, ovo, hamburguer e queijo', 10.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Cheese bacon', 'Pão, bacon, hamburguer e queijo', 10.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd82', 'Batatas fritas', 'Batatas', 4.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd82', 'Tiras de frango', 'Tiras de peito de frango frito', 5.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd83', 'Cerveja Heineken', 'Água, malte, lúpulo e levedura', 10.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd83', 'Coca-cola', 'Matéria escura ', 9.75, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd84', 'Pudim', 'leite condensado, ovos e açucar', 5.99, now(), now());

