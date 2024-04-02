alter table "public"."comments" alter column "food_id" drop not null;
alter table "public"."comments" add column "food_id" int4;
