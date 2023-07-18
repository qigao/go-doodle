create table if not exists users
(
    id         bigint unsigned auto_increment primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,
    username   longtext    not null,
    email      longtext    not null,
    password   longtext    not null,
    bio        longtext    null,
    image      longtext    null
);

create table if not exists articles
(
    id          bigint unsigned auto_increment primary key,
    created_at  datetime(3)     null,
    updated_at  datetime(3)     null,
    deleted_at  datetime(3)     null,
    slug        longtext        not null,
    title       longtext        not null,
    description longtext        null,
    body        longtext        null,
    author_id   bigint unsigned null,
    constraint fk_articles_author
        foreign key (author_id) references users (id)
);
create table if not exists tags
(
    id         bigint unsigned auto_increment primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,
    tag        longtext    null
);
create table if not exists article_tags
(
    tag_id     bigint unsigned not null,
    article_id bigint unsigned not null,
    primary key (tag_id, article_id),
    constraint fk_article_tags_article
        foreign key (article_id) references articles (id),
    constraint fk_article_tags_tag
        foreign key (tag_id) references tags (id)
);

create table if not exists comments
(
    id         bigint unsigned auto_increment primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    article_id bigint unsigned null,
    user_id    bigint unsigned null,
    body       longtext        null,
    constraint fk_articles_comments
        foreign key (article_id) references articles (id),
    constraint fk_comments_user
        foreign key (user_id) references users (id)
);

create table if not exists favorites
(
    article_id bigint unsigned not null,
    user_id    bigint unsigned not null,
    primary key (article_id, user_id),
    constraint fk_favorites_article
        foreign key (article_id) references articles (id),
    constraint fk_favorites_user
        foreign key (user_id) references users (id)
);

create table if not exists follows
(
    follower_id  bigint unsigned not null,
    following_id bigint unsigned not null,
    primary key (follower_id, following_id),
    constraint fk_users_followers
        foreign key (following_id) references users (id),
    constraint fk_users_followings
        foreign key (follower_id) references users (id)
);
