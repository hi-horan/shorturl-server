

DROP TABLE short_urls IF EXISTS;
CREATE TABLE short_urls(
   short_key VARCHAR(8) NOT NULL,
   original_url VARCHAR(1024) NOT NULL,

   expire_time DATETIME DEFAULT NULL,

   create_time DATETIME NOT NULL,
   update_time DATETIME NOT NULL,
   PRIMARY KEY (short_key)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='key与长链映射表';
