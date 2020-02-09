drop type if exists branch_status;
create type branch_status as enum(
    'active',
    'inactive'
);

drop table if exists branches;
create table branches(
    id uuid,
    name varchar(100) not null,
    create_at timestamp not null default current_timestamp,
    status branch_status not null default 'active',
    primary key(id),
    unique(name)
);

insert into branches(id, name) values(
    '6c2cfb13-82a8-4fc8-85cb-82aedc9b982d',
    'HQ'
);
