CREATE TABLE news
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    title   TEXT      NOT NULL,
    content TEXT      NOT NULL
);

CREATE TABLE newscategories
(
    id         BIGSERIAL                   NOT NULL PRIMARY KEY,
    newsid     BIGINT REFERENCES news (id) NOT NULL,
    categoryid BIGINT                      NOT NULL
);
