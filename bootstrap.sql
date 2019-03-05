CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;

CREATE TABLE IF NOT EXISTS node_measurement
(
  time      TIMESTAMP WITHOUT TIME ZONE not null, -- UTC timestamp
  nodename  varchar(255)                not null,
  timeslice float                       not null,
  cpu       float                       not null,
  mem       float                       not null
);

SELECT create_hypertable('node_measurement', 'time', if_not_exists=> true);

CREATE TABLE IF NOT EXISTS process_measurement
(
  time         TIMESTAMP WITHOUT TIME ZONE not null, -- UTC timestamp
  nodename     varchar(255)                not null,
  process_name varchar(255)                not null,
  timeslice    float                       not null,
  cpu          float                       not null,
  mem          float                       not null
);

SELECT create_hypertable('process_measurement', 'time', if_not_exists=> true);
