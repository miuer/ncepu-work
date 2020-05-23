
CREATE TABLE IF NOT EXISTS students(
		id           CHAR(12)  NOT NULL,
		sname        VARCHAR(20) NOT NULL,
		class        VARCHAR(20) NOT NULL,
		PRIMARY KEY(id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


INSERT INTO students(id,sname,class) VALUES("201709000403","miuer","软件1701");
INSERT INTO students(id,sname,class) VALUES("201709000512","张","软件1702");
INSERT INTO students(id,sname,class) VALUES("201709000401","曹","计科实1701");
INSERT INTO students(id,sname,class) VALUES("201709000523","李","信安1701");
INSERT INTO students(id,sname,class) VALUES("201709001007","苏","网络1701");

