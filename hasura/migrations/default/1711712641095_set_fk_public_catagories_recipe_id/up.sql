alter table "public"."catagories"
  add constraint "catagories_recipe_id_fkey"
  foreign key ("recipe_id")
  references "public"."recipe"
  ("id") on update restrict on delete restrict;
