/*
 Source Server         : pgku
 Source Server Type    : PostgreSQL
 Source Server Version : 140005
 Source Host           : localhost:5432
 Source Catalog        : bedbo
 Source Schema         : public
 
 Target Server Type    : PostgreSQL
 Target Server Version : 140005
 File Encoding         : 65001
 
 Date: 08/01/2023 22:52:51
 */
-- ----------------------------
-- Sequence structure for trx_order_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "trx_order_id_seq";
CREATE SEQUENCE "trx_order_id_seq" INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;
-- ----------------------------
-- Table structure for mst_customer
-- ----------------------------
DROP TABLE IF EXISTS "mst_customer";
CREATE TABLE "mst_customer" (
  "id" uuid NOT NULL,
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "address" varchar(255) COLLATE "pg_catalog"."default",
  "no_handphone" varchar(13) COLLATE "pg_catalog"."default",
  "gender" char(1) COLLATE "pg_catalog"."default"
);
-- ----------------------------
-- Records of mst_customer
-- ----------------------------
BEGIN;
INSERT INTO "mst_customer" (
    "id",
    "name",
    "address",
    "no_handphone",
    "gender"
  )
VALUES (
    '49f49589-d899-4b85-b67c-d6c588ad3f10',
    'ase',
    'asd',
    'asd',
    'L'
  );
INSERT INTO "mst_customer" (
    "id",
    "name",
    "address",
    "no_handphone",
    "gender"
  )
VALUES (
    'c8a9741c-4225-4fe2-a5ba-fc8a26f8c40e',
    'ase',
    'asd',
    'asd',
    'P'
  );
INSERT INTO "mst_customer" (
    "id",
    "name",
    "address",
    "no_handphone",
    "gender"
  )
VALUES (
    'aa16320f-748d-40e5-a05c-d1c45ea73fb8',
    'Fajar',
    'asd',
    'asd',
    'L'
  );
INSERT INTO "mst_customer" (
    "id",
    "name",
    "address",
    "no_handphone",
    "gender"
  )
VALUES (
    '80d5ee31-4e31-40d9-9a7f-5c187c9dbfe6',
    'Rizki',
    'asd',
    'asd',
    'L'
  );
INSERT INTO "mst_customer" (
    "id",
    "name",
    "address",
    "no_handphone",
    "gender"
  )
VALUES (
    '044e05a4-ed66-439f-975b-fab9fc9ef678',
    'ase',
    'asd',
    'asd',
    'P'
  );
COMMIT;
-- ----------------------------
-- Table structure for mst_product
-- ----------------------------
DROP TABLE IF EXISTS "mst_product";
CREATE TABLE "mst_product" (
  "id" uuid NOT NULL,
  "name" varchar(150) COLLATE "pg_catalog"."default",
  "description" varchar(255) COLLATE "pg_catalog"."default",
  "unit" varchar(100) COLLATE "pg_catalog"."default",
  "price" numeric(10, 2),
  "quantity" int8,
  "category" varchar(50) COLLATE "pg_catalog"."default"
);
-- ----------------------------
-- Records of mst_product
-- ----------------------------
BEGIN;
INSERT INTO "mst_product" (
    "id",
    "name",
    "description",
    "unit",
    "price",
    "quantity",
    "category"
  )
VALUES (
    '85068593-aea7-47ca-a6b9-5b37622f0872',
    'Aqua',
    '',
    '1 liter',
    5000.00,
    100,
    'minuman'
  );
INSERT INTO "mst_product" (
    "id",
    "name",
    "description",
    "unit",
    "price",
    "quantity",
    "category"
  )
VALUES (
    '06c2df2b-4bea-4c1b-932b-ea389aad48a7',
    'Minyak Sajen',
    '',
    '1 liter',
    28000.00,
    10,
    'minuman'
  );
INSERT INTO "mst_product" (
    "id",
    "name",
    "description",
    "unit",
    "price",
    "quantity",
    "category"
  )
VALUES (
    'c575b364-1962-480f-a9b9-6583a6457c24',
    'Anggur',
    NULL,
    '1 kg',
    50000.00,
    20,
    'buah'
  );
INSERT INTO "mst_product" (
    "id",
    "name",
    "description",
    "unit",
    "price",
    "quantity",
    "category"
  )
VALUES (
    'ed110c8e-6491-48aa-b2a9-4240a4bedf35',
    'Jeruk',
    NULL,
    '1 kg',
    18000.13,
    30,
    'buah'
  );
COMMIT;
-- ----------------------------
-- Table structure for trx_order
-- ----------------------------
DROP TABLE IF EXISTS "trx_order";
CREATE TABLE "trx_order" (
  "id" int4 NOT NULL DEFAULT nextval('trx_order_id_seq'::regclass),
  "order_date" date,
  "total" numeric(10, 2),
  "status" varchar(10) COLLATE "pg_catalog"."default",
  "customer_id" uuid
);
-- ----------------------------
-- Records of trx_order
-- ----------------------------
BEGIN;
INSERT INTO "trx_order" (
    "id",
    "order_date",
    "total",
    "status",
    "customer_id"
  )
VALUES (
    9,
    '2023-01-08',
    20000.00,
    'UNPAID',
    'c8a9741c-4225-4fe2-a5ba-fc8a26f8c40e'
  );
INSERT INTO "trx_order" (
    "id",
    "order_date",
    "total",
    "status",
    "customer_id"
  )
VALUES (
    8,
    '2023-01-08',
    20000.00,
    'PAID',
    '044e05a4-ed66-439f-975b-fab9fc9ef678'
  );
COMMIT;
-- ----------------------------
-- Table structure for trx_order_detail
-- ----------------------------
DROP TABLE IF EXISTS "trx_order_detail";
CREATE TABLE "trx_order_detail" (
  "id" uuid NOT NULL,
  "product_id" uuid,
  "unit_price" numeric(10, 2),
  "qty" int8,
  "total" numeric(10, 2),
  "order_id" int8
);
-- ----------------------------
-- Records of trx_order_detail
-- ----------------------------
BEGIN;
INSERT INTO "trx_order_detail" (
    "id",
    "product_id",
    "unit_price",
    "qty",
    "total",
    "order_id"
  )
VALUES (
    '33023c94-e98a-4b1d-ba30-713bd31cc5ad',
    '85068593-aea7-47ca-a6b9-5b37622f0872',
    5000.00,
    4,
    20000.00,
    9
  );
COMMIT;
-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
  "id" uuid NOT NULL,
  "email" varchar(100) COLLATE "pg_catalog"."default",
  "password" varchar(100) COLLATE "pg_catalog"."default",
  "role" varchar(20) COLLATE "pg_catalog"."default"
);
-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO "user" ("id", "email", "password", "role")
VALUES (
    '0928f75f-5b5e-4d3d-ae0e-e5f6fa8f78d8',
    're@gmail.com',
    '$2a$10$HQkz2qMGzyE9e0QzKd//y.TX75rB6cjk8CJMvytGBWQJInZufxFAu',
    'admin'
  );
COMMIT;
-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "trx_order_id_seq" OWNED BY "trx_order"."id";
SELECT setval('"trx_order_id_seq"', 9, true);
-- ----------------------------
-- Primary Key structure for table mst_customer
-- ----------------------------
ALTER TABLE "mst_customer"
ADD CONSTRAINT "customer_pkey" PRIMARY KEY ("id");
-- ----------------------------
-- Primary Key structure for table mst_product
-- ----------------------------
ALTER TABLE "mst_product"
ADD CONSTRAINT "mst_product_pkey" PRIMARY KEY ("id");
-- ----------------------------
-- Primary Key structure for table trx_order
-- ----------------------------
ALTER TABLE "trx_order"
ADD CONSTRAINT "trx_order_pkey" PRIMARY KEY ("id");
-- ----------------------------
-- Primary Key structure for table trx_order_detail
-- ----------------------------
ALTER TABLE "trx_order_detail"
ADD CONSTRAINT "trx_order_detail_pkey1" PRIMARY KEY ("id");
-- ----------------------------
-- Uniques structure for table user
-- ----------------------------
ALTER TABLE "user"
ADD CONSTRAINT "email" UNIQUE ("email");
-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "user"
ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");
-- ----------------------------
-- Foreign Keys structure for table trx_order
-- ----------------------------
ALTER TABLE "trx_order"
ADD CONSTRAINT "customer_id" FOREIGN KEY ("customer_id") REFERENCES "mst_customer" ("id") ON DELETE CASCADE ON UPDATE RESTRICT;
-- ----------------------------
-- Foreign Keys structure for table trx_order_detail
-- ----------------------------
ALTER TABLE "trx_order_detail"
ADD CONSTRAINT "order_id" FOREIGN KEY ("order_id") REFERENCES "trx_order" ("id") ON DELETE CASCADE ON UPDATE RESTRICT;
ALTER TABLE "trx_order_detail"
ADD CONSTRAINT "product_id" FOREIGN KEY ("product_id") REFERENCES "mst_product" ("id") ON DELETE CASCADE ON UPDATE RESTRICT;