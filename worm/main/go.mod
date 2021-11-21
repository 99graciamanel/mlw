module github.com/99graciamanel/mlw/worm/main

go 1.17

replace github.com/99graciamanel/mlw/worm/infection => ../infection

require (
	github.com/99graciamanel/mlw/worm/ddos v0.0.0-00010101000000-000000000000
	github.com/99graciamanel/mlw/worm/scan v0.0.0-00010101000000-000000000000
)

replace github.com/99graciamanel/mlw/worm/scan => ../scan

replace github.com/99graciamanel/mlw/worm/ddos => ../ddos
