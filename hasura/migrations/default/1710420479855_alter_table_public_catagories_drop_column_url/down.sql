alter table "public"."catagories" alter column "url" drop not null;
alter table "public"."catagories" add column "url" text;
