alter table "public"."recipe" drop constraint "recipe_catagory_id_fkey",
  add constraint "recipe_catagory_id_fkey"
  foreign key ("catagory_id")
  references "public"."catagories"
  ("id") on update restrict on delete cascade;
