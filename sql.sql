drop database if exists hang;
create database hang;
use hang;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

drop table if exists owner;
create table owner
(
    id    int         not null auto_increment comment '业主编号' primary key,
    name  varchar(50) not null comment '业主姓名',
    work  varchar(50) not null comment '业主工作',
    phone varchar(50) not null comment '业主电话'
) comment '业主信息'
    AUTO_INCREMENT = 10
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

drop table if exists room;
create table room
(
    id         int         not null auto_increment comment '房屋编号' primary key,
    owner_name varchar(50) default null comment '业主姓名',
    area       int(10)     not null comment '房屋面积',
    number     varchar(50) not null comment '房号'
) comment '房屋信息'
    AUTO_INCREMENT = 100
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

drop table if exists room_info;
create table room_info
(
    id          int          not null auto_increment comment '费用编号' primary key,
    number      varchar(50)  not null comment '房号',
    water       float(10, 2) not null comment '用水量',
    electricity float(10, 2) not null comment '用电量',
    year        int          not null comment '年份',
    month       int          not null comment '月份',
    fee         float(10, 2) default null comment '总费用'
) comment '房屋费用'
    AUTO_INCREMENT = 1000
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

drop table if exists department;
create table department
(
    id           int         not null auto_increment comment '部门编号' primary key,
    name         varchar(50) not null comment '部门姓名',
    phone        varchar(50) not null comment '部门电话',
    manager_name varchar(50) not null comment '部门主管'
) comment '部门信息'
    AUTO_INCREMENT = 10000
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

drop table if exists employee;
create table employee
(
    id              int         not null auto_increment comment '员工编号' primary key,
    department_name varchar(50) not null comment '部门姓名',
    username        varchar(50) not null comment '员工姓名',
    password        varchar(50) not null comment '员工密码',
    sex             varchar(50) not null comment '员工性别',
    phone           varchar(50) not null comment '员工电话'
) comment '员工信息'
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

INSERT INTO hang.department (name, phone, manager_name)
VALUES ('department1', '15387113436', 'hang');
INSERT INTO hang.department (name, phone, manager_name)
VALUES ('department2', '15387113436', 'ljj');

INSERT INTO hang.employee (username, password, sex, phone, department_name)
VALUES ('hang', '123456', 'male', '15387113436', 'department1');
INSERT INTO hang.employee (username, password, sex, phone, department_name)
VALUES ('lp', '123456', 'famale', '15387113436', 'department1');
INSERT INTO hang.employee (username, password, sex, phone, department_name)
VALUES ('ljj', '123456', 'male', '15387113436', 'department2');

INSERT INTO hang.owner (name, work, phone)
VALUES ('hang', 'work1', '15387113436');
INSERT INTO hang.owner (name, work, phone)
VALUES ('lp', 'work2', '15387113436');

INSERT INTO hang.room (owner_name, area, number)
VALUES ('hang', 100, '2-303');
INSERT INTO hang.room (owner_name, area, number)
VALUES ('hang', 200, '1-202');
INSERT INTO hang.room (owner_name, area, number)
VALUES ('lp', 100, '4-512');
INSERT INTO hang.room (owner_name, area, number)
VALUES ('ljj', 100, '2-301');
INSERT INTO hang.room (owner_name, area, number)
VALUES ('ljj', 120, '3-301');
INSERT INTO hang.room (owner_name, area, number)
VALUES ('ljj', 100, '4-301');

INSERT INTO hang.room_info (number, water, electricity, year, month)
VALUES ('2-303', 100, 100, 2022, 12);
INSERT INTO hang.room_info (number, water, electricity, year, month)
VALUES ('2-303', 50, 50, 2023, 1);
INSERT INTO hang.room_info (number, water, electricity, year, month)
VALUES ('1-202', 100, 100, 2023, 1);
INSERT INTO hang.room_info (number, water, electricity, year, month)
VALUES ('4-512', 50, 50, 2023, 1);
