module github.com/99graciamanel/mlw/worm/main

go 1.17

replace github.com/99graciamanel/mlw/worm/infection => ../infection

require (
	github.com/99graciamanel/mlw/worm/ddos v0.0.0-00010101000000-000000000000
	github.com/99graciamanel/mlw/worm/infection v0.0.0-00010101000000-000000000000
	github.com/99graciamanel/mlw/worm/scan v0.0.0-00010101000000-000000000000
)

require (
	github.com/jasonlvhit/gocron v0.0.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)

replace github.com/99graciamanel/mlw/worm/scan => ../scan

replace github.com/99graciamanel/mlw/worm/ddos => ../ddos
