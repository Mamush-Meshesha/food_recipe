alter table "public"."recipe" alter column "instraction" drop not null;
alter table "public"."recipe" add column "instraction" text;
