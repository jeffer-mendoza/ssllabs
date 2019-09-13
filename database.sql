

-- Host (id, hostname)

-- HostInfo(sslGrade, previous, logo, title, isDown, hostId, timestamp)

-- Server (address, sslGrade, country, owner, checksum)

CREATE TABLE servers (
	id INT8 NOT NULL,
	address VARCHAR NULL,
	ssl_grade VARCHAR NULL,
	country VARCHAR NULL,
	owner VARCHAR NULL,
	checksum VARCHAR,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, address, ssl_grade, country, owner)
)

CREATE TABLE host_info (
    id INT8 NOT NULL,
    ssl_grade VARCHAR,
    previous_grade VARCHAR,
    logo VARCHAR,
    title VARCHAR,
    is_down BOOL,
    host_id INT8,
    timestamp TIMESTAMPTZ,
    checksum VARCHAR
)

CREATE TABLE host (
    id INT8 NOT NULL,
    name VARCHAR NOT NULL
)

