create sequence table_name_id_seq;

alter sequence table_name_id_seq owner to postgres;

create sequence table_name_id_chat_seq;

alter sequence table_name_id_chat_seq owner to postgres;

create table messages
(
    id      bigint default nextval('table_name_id_seq'::regclass)      not null
        constraint messages_pk
            primary key,
    time    timestamp                                                  not null,
    body    varchar,
    id_chat bigint default nextval('table_name_id_chat_seq'::regclass) not null
);

comment on column messages.id is 'Id сообщения';

comment on column messages.time is 'Временная метка';

comment on column messages.body is 'Тело сообщения';

comment on column messages.id_chat is 'Id чата';

alter table messages
    owner to postgres;

alter sequence table_name_id_seq owned by messages.id;

alter sequence table_name_id_chat_seq owned by messages.id_chat;

create table users
(
    id    bigserial
        constraint users_pk
            primary key,
    name  varchar,
    login varchar not null
);

comment on column users.id is 'Id пользователя';

comment on column users.name is 'Имя пользователя';

comment on column users.login is 'Логин пользователя';

alter table users
    owner to postgres;

create table chats
(
    id   bigserial
        constraint chats_pk
            primary key,
    name integer
);

comment on column chats.id is 'Id чата';

comment on column chats.name is 'Имя чата';

alter table chats
    owner to postgres;

create table users_chats
(
    id_user bigserial,
    id_chat bigserial
);

comment on column users_chats.id_user is 'Id пользователя';

comment on column users_chats.id_chat is 'Id чата';

alter table users_chats
    owner to postgres;