package data

type table struct {
	name   string
	create string
}

var billsTable = table{
	name: "bills",
	create: `
		create table if not exists bills (
			id integer primary key autoincrement,
			paid boolean default false,
			name varchar(20) not null,
			day_of_month integer
		);
`,
}

var tables = []table{billsTable}
