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
    "id_customer"           uuid references "customers" ("id"),
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

-- Inserção dos registros na ordem correta

-- insert order_status
INSERT INTO public.order_status VALUES
    ('Pending', 'Order has been created but not processed yet'),
    ('Received', 'Order has been received and is ready to prepare'),
    ('InPreparation', 'Order is currently being prepared'),
    ('Ready', 'Order is ready for pickup or delivery'),
    ('Completed', 'Order has been completed and delivered'),
    ('Cancelled', 'Order has been cancelled');

-- insert customers
INSERT INTO public.customers VALUES
    ('1bf11979-6b27-495c-a604-1e9db32a86bd', 'Lord Voldemort', '11111111111', 'lord.voldermort@gmail.com', now(), now()),
    ('2bf11979-6b27-495c-a604-1e9db32a86bd', 'Ned Flanders', '22222222222', 'ned.flanders@gmail.com', now(), now()),
    ('3bf11979-6b27-495c-a604-1e9db32a86bd', 'Vladimir Putin', '66666666666', 'vladimir.putin@gmail.com', now(), now());

-- insert product_categories
INSERT INTO public.product_categories VALUES
    ('83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Lanche'),
    ('83a7f6c5-b717-4159-b2b2-675bccb5cd82', 'Acompanhamento'),
    ('83a7f6c5-b717-4159-b2b2-675bccb5cd83', 'Bebida'),
    ('83a7f6c5-b717-4159-b2b2-675bccb5cd84', 'Sobremesa');

-- insert products
INSERT INTO public.products (id, id_category, description, ingredients, price, created_at, updated_at) VALUES
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Hamburguer simples', 'Pão e hamburguer', 7.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Cheese burguer', 'Pão, hamburguer e queijo', 9.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Cheese egg burguer', 'Pão, ovo, hamburguer e queijo', 10.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd81', 'Cheese bacon', 'Pão, bacon, hamburguer e queijo', 10.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd82', 'Batatas fritas', 'Batatas', 4.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd82', 'Tiras de frango', 'Tiras de peito de frango frito', 5.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd83', 'Cerveja Heineken', 'Água, malte, lúpulo e levedura', 10.99, now(), now()),
    (gen_random_uuid(), '83a7f6c5-b717-4159-b2b2-675bccb5cd83', 'Coca-cola', 'Matéria escura ', 9.75, now(), now()),
    ('92e63fba-795e-4cdc-b8e2-96529703e265', '83a7f6c5-b717-4159-b2b2-675bccb5cd84', 'Pudim', 'leite condensado, ovos e açucar', 5.99, now(), now());

-- insert orders
INSERT INTO public.orders VALUES
    ('109186d7-7b46-4c67-863e-70a08b1e7315', '3bf11979-6b27-495c-a604-1e9db32a86bd', 'Pending', 0, now(), now(), now());

-- insert order_products
INSERT INTO public.order_products VALUES
    ('109186d7-7b46-4c67-863e-70a08b1e7315', '92e63fba-795e-4cdc-b8e2-96529703e265');
