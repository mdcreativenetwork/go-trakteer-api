MIG := @migrate
NODEMON := @nodemon
SIGNAL := SIGTERM
DSN := postgres://restuwahyu13:restuwahyu13@localhost:5432/postgres?sslmode=disable
MIGDIR := databases/migrations/

start:
	${NODEMON} -V -x go run main.go --signal ${SIGNAL}

migmake:
ifdef name
		${MIG} -verbose create -ext sql -dir ${MIGDIR} ${name}
endif

migup:
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose up

migdown:
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose down -all

migupspec:
ifdef target
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose down ${target}
endif

migdownspec:
ifdef target
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose down ${target}
endif

migdrop:
		${MIG} -database ${DSN} -path ${MIGDIR} -verbose drop -f