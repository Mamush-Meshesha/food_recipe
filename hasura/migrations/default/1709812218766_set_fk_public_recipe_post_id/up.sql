alter table "public"."recipe"
  add constraint "recipe_post_id_fkey"
  foreign key ("post_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;
