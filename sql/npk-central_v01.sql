drop function if exists hq_id();
create function hq_id() returns uuid as $$
begin
    return '6c2cfb13-82a8-4fc8-85cb-82aedc9b982d';
end;
$$ language plpgsql;

drop type if exists branch_status cascade;
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
    hq_id(),
    'HQ'
);

drop function if exists delete_branch();
create function delete_branch() returns trigger as $$
begin
    if old.id = hq_id() then
        raise exception 'cannot delete HQ branch';
    else
        return old;
    end if;
end;
$$ language plpgsql;

create trigger before_branch_delete
    before delete on branches
    for each row
        execute function delete_branch();

