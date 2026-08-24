module github.com/tylland/transition/commoninterface

go 1.26.5

require (
	github.com/ledongthuc/pdf v0.0.0-20250511090121-5959a4027728
	github.com/lestrrat-go/helium v0.7.0
)

require (
	github.com/dlclark/regexp2 v1.12.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

require (
	github.com/google/uuid v1.6.0
	github.com/tylland/transition/specs v0.0.0-00010101000000-000000000000
)

replace github.com/tylland/transition/specs => ../specs
