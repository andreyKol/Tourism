CREATE TABLE IF NOT EXISTS room_clients (
    room_id   bigint,
    client_id bigint,
    PRIMARY KEY (room_id, client_id),
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);


INSERT INTO countries (name, description) VALUES
                                              ('USA', 'United States of America, located in North America'),
                                              ('France', 'A country in Western Europe'),
                                              ('Japan', 'An island nation in East Asia'),
                                              ('Brazil', 'The largest country in South America'),
                                              ('Australia', 'A country and continent surrounded by the Indian and Pacific oceans');


