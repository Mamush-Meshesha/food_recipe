alter table "public"."blog"
  add constraint "blog_blog_id_fkey"
  foreign key ("blog_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;
