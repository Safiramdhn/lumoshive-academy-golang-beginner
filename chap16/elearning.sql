
CREATE TYPE status_enum AS enum ('active', 'deleted');
create type attendance_status_enum as enum ('ontime', 'late');
create  type assignment_status_enum as enum ('not_started', 'submited','missing');

CREATE TABLE users (
    id INT PRIMARY KEY NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);



create table admins (
	id int primary key not null unique,
	first_name varchar(255) not null,
	last_name varchar(255) not null,
	user_id int references users(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table student (
	id int primary key not null unique,
	first_name varchar (255) not null,
	last_name varchar (255) not null,
	user_id int references users(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table mentor (
 	id int primary key not null unique,
	first_name varchar (255) not null,
	last_name varchar (255) not null,
	user_id int references users(id) not null,
	added_by int references admins(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table material (
	id int primary key not null unique,
	title varchar(255),
	description varchar,
	media_url varchar,
	added_by int references admins(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table schedule (
	id int primary key not null unique,
	date date not null,
	time time not null,
	added_by int references admins(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table classes (
	id int primary key not null unique,
	title varchar(255),
	description varchar,
	added_by int references admins(id) not null,
	mentor_id int references mentor(id) not null,
	schedule_id int references schedule(id) not null,
	material_id int references material(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table announcement (
	id int primary key not null unique,
	title varchar(255),
	description varchar,
	added_by int references admins(id) not null,
	class_id int references classes(id),
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table student_class (
	id int primary key not null unique,
	student_id int references student(id) not null,
	class_id int references classes(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table student_attendance (
	id int primary key not null unique,
	student_id int references student(id) not null,
	student_class_id int references student_class(id) not null,
	datetime timestamp,
	attendance_status attendance_status_enum,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

create table mentor_attendance (
	id int primary key not null unique,
	mentor_id int references mentor(id) not null,
	class_id int references classes(id) not null,
	datetime timestamp,
	attendance_status attendance_status_enum,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)


create table assignments (
	id int primary key not null unique,
	class_id int references classes(id) not null,
	deadline timestamp not null,
	description varchar,
	title varchar(255) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

create table student_assignment (
	id int primary key not null unique,
	assignment_id int references assignments(id) not null,
	student_id int references student(id) not null,
	submit_date timestamp,
	assignment_status assignment_status_enum default 'not_started',
	score real,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)

create table leaderboard (
	id int primary key not null unique,
	student_id int references student(id) not null,
	class_id int references classes(id) not null,
	total_score real,
	total_attendance int,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)