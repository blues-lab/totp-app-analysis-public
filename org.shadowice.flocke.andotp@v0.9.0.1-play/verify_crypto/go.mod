module main

go 1.18

replace github.com/blues-lab/totp-app-analysis-public/utils => ../../utils

require (
	github.com/blues-lab/totp-app-analysis-public/utils v0.0.0
	golang.org/x/crypto v0.0.0-20221010152910-d6f0a8c073c2
)
