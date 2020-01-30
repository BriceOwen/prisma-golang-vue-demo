GCC=go
GCMD=run
GOPATH=src/main.go

PRISMADIR=cd prisma
DEPLOY=prisma deploy
GENERATE=prisma generate

run:
	${GCC} ${GCMD} ${GOPATH}

build_db:
	${PRISMADIR} && ${DEPLOY} && ${GENERATE}

install:
	make install_routes

install_routes:
	go get -u github.com/gorrilla/mux