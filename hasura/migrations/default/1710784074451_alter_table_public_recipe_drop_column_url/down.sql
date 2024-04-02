alter table "public"."recipe" alter column "url" drop not null;
alter table "public"."recipe" add column "url" jsonb;
