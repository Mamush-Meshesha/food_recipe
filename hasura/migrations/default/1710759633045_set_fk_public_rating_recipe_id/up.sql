alter table "public"."rating" drop constraint "rating_post_id_fkey",
  add constraint "rating_recipe_id_fkey"
  foreign key ("recipe_id")
  references "public"."recipe"
  ("id") on update restrict on delete no action;
