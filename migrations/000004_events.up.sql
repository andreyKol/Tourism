CREATE TABLE IF NOT EXISTS events
(
    id           bigserial primary key,
    name         text   not null,
    description  text   not null,
    country_id   int,
    date         timestamptz,
    image        text,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);
INSERT INTO events (name, description, country_id, date, image) VALUES
                                                                        ('Белые ночи в Санкт-Петербурге', 'Фестиваль, который проходит летом, когда солнце практически не заходит за горизонт. Концерты, спектакли и уличные представления проходят по всему городу.', 1, '2024-06-21 00:00:00+03', 'https://vsegda-pomnim.com/uploads/posts/2022-02/1645919144_58-vsegda-pomnim-com-p-belie-nochi-foto-66.jpg'),
                                                                        ('Байкальский ледовый марафон', 'Экстремальный марафон по льду Байкала, в котором участвуют бегуны со всего мира. Уникальная возможность испытать себя и насладиться красотой зимнего Байкала.', 1, '2024-02-25 09:00:00+08', 'https://baikal-marathon.org/userfiles/image/3S8A2929.jpg'),
                                                                        ('Московский Международный кинофестиваль', 'Один из старейших кинофестивалей в мире, на котором представляются новые фильмы и встречаются мировые звезды кино.', 1, '2024-04-15 19:00:00+03', 'https://www.m24.ru/b/d/nAgWUB67nUJkzJ7jP6aO_pb2kY3t-dvtg32Qhv2YqzbZJC3PU2mcy3ou4cNb8QfTuNGV_CeILNx_SL-4hSMhMXXfqDgJqgjJnSrxUBPiuy9YKH1Ypyom2ybiaByXQz5QEreaOfg=ON3895A9omGm-Z7TpwaEwg.jpg');

INSERT INTO events (name, description, country_id, date, image) VALUES
                                                                        ('Фестиваль цветения сакуры', 'Один из самых знаменитых и красивых фестивалей Японии, который проходит весной. Люди собираются в парках, чтобы насладиться видом цветущих вишневых деревьев.', 2, '2024-04-01 10:00:00+09', 'https://vsegda-pomnim.com/uploads/posts/2022-04/1650624931_7-vsegda-pomnim-com-p-yaponiya-prazdnik-tsveteniya-sakuri-foto-8.jpg'),
                                                                        ('Гион Мацури в Киото', 'Один из самых известных традиционных фестивалей Японии, который проводится в июле. Парады, традиционные костюмы и танцы привлекают множество туристов.', 2, '2024-07-17 18:00:00+09', 'https://www.mitrey.ru/wp-content/uploads/2018/07/gion-macuri-glavnyj-prazdnik-kioto.jpg'),
                                                                        ('Фестиваль снежных скульптур в Саппоро', 'Знаменитый зимний фестиваль, на котором представлены ледяные и снежные скульптуры. Мероприятие привлекает множество туристов и художников.', 2, '2024-02-05 10:00:00+09', 'https://kartinki.pibig.info/uploads/posts/2023-04/1682450895_kartinki-pibig-info-p-snezhnie-skulpturi-kartinki-arti-krasivo-57.jpg');
