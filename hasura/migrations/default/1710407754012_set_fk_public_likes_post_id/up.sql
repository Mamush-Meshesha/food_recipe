alter table "public"."likes"
  add constraint "likes_post_id_fkey"
  foreign key ("post_id")
  references "public"."recipe"
  ("id") on update restrict on delete restrict;
