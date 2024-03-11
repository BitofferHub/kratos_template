create table `t_user_info` (
                                   `id`                  bigint(20)       not null AUTO_INCREMENT comment '用户ID',
                                   `user_name` varchar(65)  not null                comment '用户名字',
                                   `pwd` varchar(65)  not null                comment '用户密码',
                                   `sex`                 int(11)      not null                comment '性别',
                                   `age`                 int(11)      not null                comment '年龄',
                                   `email`               varchar(128)                         comment '邮箱',
                                   `contact`               varchar(128)                         comment '联系地址',
                                   `mobile`              varchar(64)  not null                comment '手机号',
                                   `id_card`             varchar(64)  not null                comment '证件号',
                                   `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                                   `modify_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `idx_username` (`user_name`),
                                   UNIQUE KEY `idx_mobile` (`mobile`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '用户信息表' ;

insert into t_user_info(user_name, pwd, age, sex, mobile, id_card) values("niuge", "123321", 18, 1, "18676662555", "5106811222223333");