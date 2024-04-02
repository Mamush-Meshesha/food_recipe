alter table "public"."bookmarks" add constraint "bookmarks_user_id_recipe_id_key" unique ("user_id", "recipe_id");
