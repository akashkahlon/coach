create table account_organisations (
	id serial primary key,
	account_id int not null,
	organisation_id int not null,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp
);