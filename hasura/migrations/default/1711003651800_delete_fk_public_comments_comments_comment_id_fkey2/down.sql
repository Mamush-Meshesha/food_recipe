alter table "public"."comments"
  add constraint "comments_comment_id_fkey2"
  foreign key ("comment_id")
  references "public"."recipe"
  ("id") on update restrict on delete restrict;
