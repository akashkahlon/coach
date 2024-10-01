create table repositories (
	id serial primary key,
	organisation_id int not null,
	name varchar(255) not null,
	full_name varchar(255) not null,
	private boolean not null default true,
	html_url varchar(255) not null,
	github_data_all jsonb not null,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp
);