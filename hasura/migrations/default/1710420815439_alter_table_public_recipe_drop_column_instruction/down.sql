alter table "public"."recipe" alter column "instruction" drop not null;
alter table "public"."recipe" add column "instruction" jsonb;
