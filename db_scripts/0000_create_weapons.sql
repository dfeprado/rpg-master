create table weapons(
    id int auto_increment primary key,
    name varchar(32) not null,
    weight real not null,
    pa int not null,
    damage varchar(32) not null,
    two_handed int not null,
    distance real not null,
    effect varchar(32) not null,
    size varchar(1) not null,
    enhancement varchar(32) not null
);