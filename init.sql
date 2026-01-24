create table if not exists wrestler_alignment (
	alignment TEXT unique primary key
);


insert into wrestler_alignment(alignment) values ('face'), ('heel'), ('tweener');

select * from wrestler_alignment wa 

create table if not exists wrestlers (
	wrestler_id SERIAL primary key,
	ringname VARCHAR(64) unique not null,
	alignment TEXT references wrestler_alignment(alignment),
	signature_move VARCHAR(64) not null
);

insert into wrestlers(ringname, alignment, signature_move) values
('King Mystery', 'face', 'Area Codes'),
('Jonathan C. Nah', 'face', 'Fix Your Face'),
('HANS', 'heel', 'Piledriver');

select * from wrestlers;