alter table "public"."recipe" alter column "ingradient" drop not null;
alter table "public"."recipe" add column "ingradient" text;
