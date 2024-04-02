alter table "public"."rating" drop constraint "rating_user_id_recipe_id_key";
alter table "public"."rating" add constraint "rating_user_id_key" unique ("user_id");
