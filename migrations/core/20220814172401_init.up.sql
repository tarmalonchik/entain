create table if not exists "user"
(
    id                text constraint user_pk not null primary key,
    current_amount      bigint                   not null default 0,
    created_at          timestamptz not null default now(),
    updated_at          timestamptz
);

insert into "user" (id, current_amount) values ('1', 0);
insert into "user" (id,current_amount) values ('2',0);
insert into "user" (id,current_amount) values ('3', 0);

create table if not exists transaction
(
    id                  serial constraint transaction_pk primary key,
    user_id             text not null,
    external_id         text not null,
    source_type         text not null,
    amount              bigint                   not null default 0,
    created_at          timestamptz not null default now()
);
create index if not exists idx_transaction_external_id on transaction(external_id);

alter table transaction
    add constraint fk_transaction_user_id foreign key (user_id) references "user" (id);
alter table transaction add constraint unique_transaction_external_id unique (external_id);