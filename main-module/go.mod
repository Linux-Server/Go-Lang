module sample.com/main-module

go 1.22.4

replace sample.com/greet-module => ../greet-module

require (
	sample.com/exec-module v0.0.0-00010101000000-000000000000
	sample.com/greet-module v0.0.0-00010101000000-000000000000
)

replace sample.com/exec-module => ../exec-module
