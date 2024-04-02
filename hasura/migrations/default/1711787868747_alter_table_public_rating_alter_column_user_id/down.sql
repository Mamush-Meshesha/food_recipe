alter table "public"."rating" drop constraint "rating_user_id_key";
alter table "public"."rating" alter column "user_id" drop not null;
