alter table "public"."food"
  add constraint "food_recipe_id_fkey"
  foreign key ("recipe_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;
