alter table "public"."rating" rename column "rating_value" to "value";
alter table "public"."rating" alter column "value" drop not null;
