
/* TICKET TYPE */
EXEC CreateIntIdPKTable
	'ticket_type',
	'
		[ticket_name]		NVARCHAR(50),
		[slug]				VARCHAR(50),
		[ticket_surcharge]	INT
	'

/* SEAT TYPE */
EXEC CreateIntIdPKTable
	'seat_type',
	'
		[type_name]			NVARCHAR(50),
		[slug]				VARCHAR(50),
		[seat_surcharge]	INT
	'

/* BOOKING */
EXEC CreateIntIdPKTable
	'booking',
	'
		[show_id]		INT,
		CONSTRAINT FK_booking_showing FOREIGN KEY ([show_id]) REFERENCES showing_time(id),
	'

/* BOOKING_TICKET */
CREATE TABLE booking_ticket (
    [booking_id]		INT,
	[ticket_type_id]	INT,
	[ticket_quanity]	INT,
	CONSTRAINT PK_booking_ticket PRIMARY KEY([booking_id], [ticket_type_id]),
    CONSTRAINT FK_bt_tickettype FOREIGN KEY ([ticket_type_id]) REFERENCES ticket_type(id),
	CONSTRAINT FK_bt_booking FOREIGN KEY ([booking_id]) REFERENCES booking(id),
);

/* SEAT */
EXEC CreateIntIdPKTable
	'seat',
	'
		[seat_type_id]	INT,
		[theater_id]	INT,
		[seat_location] VARCHAR(3),
		CONSTRAINT FK_seat_seattype FOREIGN KEY ([seat_type_id]) REFERENCES seat_type(id),
		CONSTRAINT FK_seat_theater FOREIGN KEY ([theater_id]) REFERENCES theater(id),
	'

/* BOOKING_SEAT */
CREATE TABLE booking_seat (
    [booking_id]	INT,
	[seat_id]		INT,
	CONSTRAINT PK_booking_seat PRIMARY KEY([booking_id], [seat_id]),
    CONSTRAINT FK_bs_booking FOREIGN KEY ([booking_id]) REFERENCES booking(id),
	CONSTRAINT FK_bs_seat FOREIGN KEY ([seat_id]) REFERENCES seat(id),
);
