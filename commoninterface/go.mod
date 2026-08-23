module github.com/tylland/transition/commoninterface

go 1.26.5

require github.com/ledongthuc/pdf v0.0.0-20250511090121-5959a4027728

require (
	github.com/lestrrat-go/libxml2 v0.0.0-20260709020957-3f0f7bd60b63 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/tylland/transition/specs v0.0.0-00010101000000-000000000000
)

replace github.com/tylland/transition/specs => ../specs
