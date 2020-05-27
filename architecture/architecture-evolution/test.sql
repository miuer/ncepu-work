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

CREATE TABLE IF NOT EXISTS grade(
        id        BIGINT NOT NULL AUTO_INCREMENT,
		`sid`     CHAR(12)  NOT NULL,
        course    INT NOT NULL,
        score     INT NOT NULL,
		PRIMARY KEY(id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO grade(sid,course,score)VALUES("201709000403",10001,90);
INSERT INTO grade(sid,course,score)VALUES("201709000403",10002,85);
INSERT INTO grade(sid,course,score)VALUES("201709000403",10003,79);
INSERT INTO grade(sid,course,score)VALUES("201709000401",10001,95);
INSERT INTO grade(sid,course,score)VALUES("201709000401",10002,90);
INSERT INTO grade(sid,course,score)VALUES("201709000401",10003,92);


