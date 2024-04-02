alter table "public"."rating" drop constraint "rating_value_key";
alter table "public"."rating" add constraint "rating_user_id_value_key" unique ("user_id", "value");
