create table products(
    `id` int not null primary key auto_increment,
    `description` text not null,
    expiration_rate float not null,
    freezing_rate float not null,
    height float not null,
    lenght float not null,
    netweight float not null,
    product_code text not null,
    recommended_freezing_temperature float not null,
    width float not null,
    id_product_type int not null,
    id_seller int not null
);
create table employees(
    `id` int not null primary key auto_increment,
    card_number_id text not null,
    first_name text not null,
    last_name text not null,
    warehouse_id int not null
);
create table warehouses(
    `id` int not null primary key auto_increment,
    `address` text null,
    telephone text null,
    warehouse_code text null,
    minimum_capacity int null,
    minimum_temperature int null
);
create table sections(
    `id` int not null primary key auto_increment,
    section_number int not null,
    current_temperature int not null,
    minimum_temperature int not null,
    current_capacity int not null,
    minimum_capacity int not null,
    maximum_capacity int not null,
    warehouse_id int not null,
    id_product_type int not null
);
create table sellers(
    `id` int not null primary key auto_increment,
    cid int not null,
    company_name text not null,
    `address` text not null,
    telephone varchar(15) not null
);
create table buyers(
    `id` int not null primary key auto_increment,
    card_number_id text not null,
    first_name text not null,
    last_name text not null
);