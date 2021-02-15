package postgres

const createCalculationTableSQL = `
create table if not exists "calculation" (
id serial primary key,
expression text,
result float,
err varchar(64)
)`

const insertCalculationSQL = `
insert into calculation(expression)
values($1)`

const updateCalculationResultSQL = `
update calculation
set result = $2, err = $3
where id = $1`

const selectLastIDSQL = `
select id from calculation
order by id desc limit 1`

const selectByIDSQL = `
select expression from calculation
where id = $1`

const selectResultByIDSQL = `
select result, err from calculation
where id = $1`
