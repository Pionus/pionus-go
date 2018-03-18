CREATE TABLE dreams(
    id int(10) not null AUTO_INCREMENT,
    author varchar(120) null DEFAULT null,
    content text,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
