alter table "public"."rating"
  add constraint "rating_recipe_id_fkey"
  foreign key ("recipe_id")
  references "public"."food"
  ("id") on update restrict on delete restrict;
