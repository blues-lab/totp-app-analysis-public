module main

go 1.18

replace github.com/blues-lab/totp-app-analysis-public/utils => ../../../utils

require gopkg.in/square/go-jose.v2 v2.6.0

require (
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
)
