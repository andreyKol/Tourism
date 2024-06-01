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
                                                                        ('Байкальский ледовый марафон', 'Экстремальный марафон по льду Байкала, в котором участвуют бегуны со всего мира. Уникальная возможность испытать себя и насладиться красотой зимнего Байкала.', 1, '2024-02-25 09:00:00+08', 'https://sport-marafon.ru/upload/iblock/9a4/%D0%A4%D0%BE%D1%82%D0%BE-%D0%BE%D0%B1%D0%BB%D0%BE%D0%B6%D0%BA%D0%B0.jpg'),
                                                                        ('Московский Международный кинофестиваль', 'Один из старейших кинофестивалей в мире, на котором представляются новые фильмы и встречаются мировые звезды кино.', 1, '2024-04-15 19:00:00+03', 'https://www.m24.ru/b/d/nAgWUB67nUJkzJ7jP6aO_pb2kY3t-dvtg32Qhv2YqzbZJC3PU2mcy3ou4cNb8QfTuNGV_CeILNx_SL-4hSMhMXXfqDgJqgjJnSrxUBPiuy9YKH1Ypyom2ybiaByXQz5QEreaOfg=ON3895A9omGm-Z7TpwaEwg.jpg');

INSERT INTO events (name, description, country_id, date, image) VALUES
                                                                        ('Фестиваль цветения сакуры', 'Один из самых знаменитых и красивых фестивалей Японии, который проходит весной. Люди собираются в парках, чтобы насладиться видом цветущих вишневых деревьев.', 2, '2024-04-01 10:00:00+09', 'https://vsegda-pomnim.com/uploads/posts/2022-04/1650624931_7-vsegda-pomnim-com-p-yaponiya-prazdnik-tsveteniya-sakuri-foto-8.jpg'),
                                                                        ('Гион Мацури в Киото', 'Один из самых известных традиционных фестивалей Японии, который проводится в июле. Парады, традиционные костюмы и танцы привлекают множество туристов.', 2, '2024-07-17 18:00:00+09', 'https://www.mitrey.ru/wp-content/uploads/2018/07/gion-macuri-glavnyj-prazdnik-kioto.jpg'),
                                                                        ('Фестиваль снежных скульптур в Саппоро', 'Знаменитый зимний фестиваль, на котором представлены ледяные и снежные скульптуры. Мероприятие привлекает множество туристов и художников.', 2, '2024-02-05 10:00:00+09', 'https://kartinki.pibig.info/uploads/posts/2023-04/1682450895_kartinki-pibig-info-p-snezhnie-skulpturi-kartinki-arti-krasivo-57.jpg');

-- Казахстан
INSERT INTO events (name, description, country_id, date, image) VALUES
('Наурыз', 'Традиционный весенний праздник, отмечаемый в Казахстане с древних времен. Праздник символизирует обновление природы и начало нового года.', 3, '2024-03-21 09:00:00+06', 'https://kartinki.pics/uploads/posts/2022-02/thumbs/1644887636_4-kartinkin-net-p-nauriz-kartinki-4.jpg'),
('Международный фестиваль куртов', 'Фестиваль, посвященный традиционному казахскому блюду курт. Гостей ждут дегустации, мастер-классы и конкурсы.', 3, '2024-06-15 10:00:00+06', 'https://www.m24.ru/b/d/nAgWUB67nUJkzJ7jP6aI_Jn-mYHt-dvtg32Qhv2YqzbZJC3PU2mcy3ou4cNb8QfTuNGV_CeILNx_SL-4hSMhMXXfqDgJqgjJnSrxUBPiuy9YKH1Ypyom2ybiaByXQz5QEreaOfg=aY3ts0Zi8kneKwVwNr8d9Q.jpg'),
('Астана День города', 'Празднование дня столицы Казахстана Астаны. Включает концерты, фейерверки и культурные мероприятия по всему городу.', 3, '2024-07-06 18:00:00+06', 'https://sun9-13.userapi.com/impf/c631724/v631724750/3ba26/vUWTBRyjV6I.jpg?size=1200x704&quality=96&sign=c41863633da124a9ae8b5af3f99b4077&c_uniq_tag=3a3b9d30B5D1dAwXqV8G8cSsAB1NXTVz0ih7-3Y_rvE&type=album');

-- Узбекистан
INSERT INTO events (name, description, country_id, date, image) VALUES
('Шелковый и пряный фестиваль', 'Фестиваль, посвященный богатому наследию Великого шелкового пути. Включает выставки, концерты и мастер-классы.', 4, '2024-04-25 10:00:00+05', 'https://uzholidays.com/thumb/2/vZIY5ikXv8NHWwHl8PyGgQ/r/d/thumbs_b_c_2d02c7bdd2f7def1f85b4cc12774e148.jpg'),
('День Независимости Узбекистана', 'Главный государственный праздник, отмечаемый с парадами, концертами и фейерверками.', 4, '2024-09-01 09:00:00+05', 'https://img.redzhina.ru/img/dd/b2/ddb245fa4a330d586585a55f7a559310.jpg'),
('Международный фестиваль народного искусства в Хиве', 'Фестиваль, на котором представлены народные ремесла и искусства со всего мира. Проходит в древнем городе Хива.', 4, '2024-10-15 10:00:00+05', 'https://e-cis.info/upload/iblock/3ee/3eedbdec69e211ec03fb49c0cce96bab.jpg');

-- Иран
INSERT INTO events (name, description, country_id, date, image) VALUES
('Фестиваль цветов в Ширазе', 'Весенний фестиваль, посвященный цветам и розам, которые цветут в этом регионе.', 5, '2024-05-05 10:00:00+03:30', 'https://cdn.fishki.net/upload/post/2022/07/22/4192594/4735ef06764eee18f449c856152ef245.jpg'),
('Фестиваль искусств в Тегеране', 'Ежегодный фестиваль, который включает в себя выставки, театральные постановки и концерты.', 5, '2024-09-20 18:00:00+03:30', 'https://cdn.iz.ru/sites/default/files/inline/2022-07-03%2012.48.41%20%281%29.JPG'),
('Фестиваль огня Чахаршанбе-Сури', 'Традиционный иранский праздник, предшествующий Новрузу, с кострами и фейерверками.', 5, '2024-03-18 18:00:00+03:30', 'https://www.hitehranhostel.com/wp-content/uploads/2017/03/Fire-Jumping-in-Charshange-Suri.jpg');

-- ОАЭ
INSERT INTO events (name, description, country_id, date, image) VALUES
('Дубай Шопинг Фестиваль', 'Один из крупнейших в мире торговых фестивалей с огромными скидками, развлекательными программами и розыгрышами призов.', 6, '2024-01-10 10:00:00+04', 'https://go-travel.ru/upload/iblock/460/m7xlx65sn5iu7z5b4u52v9kcxemv8f2v/zimniy_shoping_festival_v_dubae_startuet_8_dekabrya.jpg'),
('Формула 1 Абу-Даби', 'Заключительный этап чемпионата мира Формулы 1, проходящий на трассе Яс Марина в Абу-Даби.', 6, '2024-11-22 15:00:00+04', 'https://grand-prixf1.ru/uploads/posts/2022-11/jzg5iinstpq.jpg'),
('Фестиваль вкусов в Дубае', 'Гастрономический фестиваль, который включает мастер-классы от шеф-поваров, дегустации и кулинарные шоу.', 6, '2024-03-05 10:00:00+04', 'https://dorognoe.ru/upload/editor/01/2020/57/03/5e34306499970_86dc3e2ac9cb1fb5baa49f1be4130ff4.jpg');

-- Кыргызстан
INSERT INTO events (name, description, country_id, date, image) VALUES
('Фестиваль искусств в Бишкеке', 'Международный фестиваль, на котором представлены театральные постановки, концерты и художественные выставки.', 7, '2024-06-10 10:00:00+06', 'https://turktoday.info/cdn/2022/06/DOS_0837-scaled.jpg'),
('День Независимости Кыргызстана', 'Празднование дня независимости страны с парадами, концертами и фейерверками.', 7, '2024-08-31 09:00:00+06', 'https://cdnuploads.aa.com.tr/uploads/VideoGallery/2021/08/31/17df48277a4def2f1c0cb03c163fc5fd.jpg'),
('Фестиваль "Игры кочевников"', 'Уникальный фестиваль, включающий в себя традиционные спортивные игры, такие как борьба, скачки и стрельба из лука.', 7, '2024-09-02 10:00:00+06', 'https://avatars.dzeninfra.ru/get-zen_doc/9663006/pub_642175258d46566563492e92_6421786265530d6563f258cd/scale_1200');

-- Китай
INSERT INTO events (name, description, country_id, date, image) VALUES
('Праздник Весны', 'Один из самых важных праздников в Китае, известный также как Китайский Новый год. Включает парады, фейерверки и семейные собрания.', 8, '2024-01-29 09:00:00+08', 'https://www.susu.ru/sites/default/files/field/image/chinese-new-year_1.jpg'),
('Фестиваль фонарей', 'Красочный фестиваль, завершающий празднование Китайского Нового года, с тысячами фонарей, украшающих улицы и парки.', 8, '2024-02-11 18:00:00+08', 'https://bestvietnam.ru/wp-content/uploads/2020/06/%D0%9F%D1%80%D0%B0%D0%B7%D0%B4%D0%BD%D0%B8%D0%BA-%D1%84%D0%BE%D0%BD%D0%B0%D1%80%D0%B5%D0%B9-%D0%B2-%D0%9A%D0%B8%D1%82%D0%B0%D0%B5.jpg'),
('Международный кинофестиваль в Пекине', 'Крупнейший кинофестиваль Китая, привлекающий кинематографистов и любителей кино со всего мира.', 8, '2024-04-16 10:00:00+08', 'https://sun9-28.userapi.com/impg/B1yVScPRuUx9lSgTJ1wbX-P4buwGmuo7yKZhwQ/7ZwwG1r-5CU.jpg?size=1280x825&quality=95&sign=75f858b5afd3c5c16c754a4e4771c011&c_uniq_tag=fnu4xkIALruguNA8IkOaQ-FOyIHJ_NWUYOqXCAu9JRE&type=album');

-- Турция
INSERT INTO events (name, description, country_id, date, image) VALUES
('Международный кинофестиваль в Анталии', 'Один из старейших кинофестивалей Турции, представляющий фильмы со всего мира.', 9, '2024-10-10 10:00:00+03', 'https://sun9-13.userapi.com/impg/rzxF44VaCMq3v2azG-w8_L7Xr7UrNanoHiPiBA/VkJjcBB2CXU.jpg?size=807x538&quality=95&sign=61288c52efcf8c122ada06a95387aa10&c_uniq_tag=ruhZ0Q5VyUi_VdanXqxt9yG04I-WZYC-ji8adlp08UA&type=album'),
('Фестиваль тюльпанов в Стамбуле', 'Весенний фестиваль, на котором тысячи тюльпанов украшают парки и сады города.', 9, '2024-04-01 10:00:00+03', 'https://avatars.dzeninfra.ru/get-zen_doc/1716636/pub_5ffc9ff6aeef3c7829c2956e_5ffca7fe7cd87011f01ca36a/scale_1200'),
('Международный фестиваль джаза в Стамбуле', 'Известный джазовый фестиваль, привлекающий музыкантов и любителей джаза со всего мира.', 9, '2024-07-01 18:00:00+03', 'https://cdn-st4.rtr-vesti.ru/p/o_1680255.jpg');

-- Египет
INSERT INTO events (name, description, country_id, date, image) VALUES
('Фестиваль солнцестояния в Абу-Симбеле', 'Уникальное событие, когда солнечные лучи освещают внутренние помещения храма Рамзеса II.', 10, '2024-10-22 06:00:00+02', 'https://mykaleidoscope.ru/x/uploads/posts/2022-09/1663422091_17-mykaleidoscope-ru-p-khram-ramzesa-v-abu-simbele-instagram-18.jpg'),
('Каирский международный кинофестиваль', 'Крупнейший кинофестиваль на Ближнем Востоке, представляющий фильмы из арабского мира и других стран.', 10, '2024-11-20 10:00:00+02', 'https://newsaf.cgtn.com/news/2022-04-13/2022-Cairo-International-Film-Festival-scheduled-for-November-19bdFFB3UnS/img/f21b9f0f18a34c52b483a36dddecab02/f21b9f0f18a34c52b483a36dddecab02.jpeg'),
('Фестиваль в Луксоре', 'Культурный фестиваль с фольклорными выступлениями, выставками и театральными постановками.', 10, '2024-02-10 10:00:00+02', 'https://english.ahram.org.eg/Media/News/2021/11/25/2021-637734764090290931-29.jpg');

-- Словакия
INSERT INTO events (name, description, country_id, date, image) VALUES
('Братиславский культурный фестиваль', 'Крупный фестиваль, на котором представлены музыка, танцы и театральные постановки со всего мира.', 11, '2024-07-15 10:00:00+02', 'https://mykaleidoscope.ru/x/uploads/posts/2022-09/1663422091_17-mykaleidoscope-ru-p-khram-ramzesa-v-abu-simbele-instagram-18.jpg'),
('Фестиваль вина в Модре', 'Традиционный фестиваль, посвященный виноделию, с дегустациями местных вин и культурными мероприятиями.', 11, '2024-09-10 10:00:00+02', 'https://www.visitbratislava.com/wp-content/uploads/2016/02/BPV1-e1555073268213.jpg'),
('Фестиваль цветов Братиславы', 'С наступлением весенней поры город становится цветущим ботаническим садом.', 11, '2024-04-10 15:00:00+02', 'https://www.mos.ru/upload/newsfeed/articles/DSC_7105.jpg');

-- Греция
INSERT INTO events (name, description, country_id, date, image) VALUES
('Афинский Фотофестиваль', 'Ежегодный Афинский фотофестиваль, основанный в 1987 году, организуется Греческим центром фотографии и представляет собой крупнейший международный фестиваль фотографии в регионе.', 12, '2024-06-22 16:00:00+02', 'https://cdn-v2.theculturetrip.com/610x407/wp-content/uploads/2017/01/27276120346_f471f38aa1_b.webp'),
('Международный Фестиваль комиксов', 'Международный фестиваль комиксов в Афинах, иногда называемый Вавилонским фестивалем, проводится с 1996 года.', 12, '2024-11-16 10:00:00+02', 'https://cdn-v2.theculturetrip.com/610x406/wp-content/uploads/2017/01/kik_4650_.webp'),
('Фестиваль Rockwave', 'Rockwave Festival - это музыкальный фестиваль, который проводится каждый июнь в Малакасе, в 40 километрах (25 милях) к северу от Афин. ', 12, '2024-03-16 10:00:00+02', 'https://cdn-v2.theculturetrip.com/610x458/wp-content/uploads/2017/01/4786006515_6bbabe429a_b.webp');
