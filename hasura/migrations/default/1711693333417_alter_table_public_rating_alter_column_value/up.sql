alter table "public"."rating" alter column "value" set not null;
alter table "public"."rating" rename column "value" to "rating_value";
