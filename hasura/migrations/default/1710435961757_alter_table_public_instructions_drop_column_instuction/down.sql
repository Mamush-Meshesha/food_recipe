alter table "public"."instructions" alter column "instuction" drop not null;
alter table "public"."instructions" add column "instuction" jsonb;
