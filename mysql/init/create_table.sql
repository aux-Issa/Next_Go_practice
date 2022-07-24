create table if not exists tasks (
  id int(11) not null auto_increment,
  title varchar(255) not null,
  content text not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update CURRENT_TIMESTAMP
  primary key (id)
);

insert into tasks (title, content)values("title","content1");