show databases;
create database AddressBook;
use AddressBook;
create table contact(
address_id int not null primary key,
firstname varchar(250) not null,
lastname varchar(250) not null,
address varchar(250) not null,
city varchar(50) not null,
state varchar(50) not null,
phoneno int not null,
email varchar(100)not null );
desc contact;

insert into contact
(address_id,firstname,lastname,address,city,state,phoneno,email)
values
(1,"Rose","Carter","2nd North Street","Pune","MAHA",987586241,"rose@mail.co"),
(2,"Amy","Owen","Northhill","Nagpur","MAHA",988636541,"oamy@gmail.co"),
(3,"Lucian","Stokes","F City","Jaipur","Rajasthan",85436541,"lstokes@gmail.com"),
(4,"Arjun","Reddy","Moon City","Belgaon","Karnataka",8532331,"reddys@gmail.com") ;
select * from contact;
show tables;

select * from contact
where city="Pune";

update contact
set phoneno=775432188
where state="MAHA";
update contact
set address="Gorebandar"
where city="Pune";
select * from contact;
