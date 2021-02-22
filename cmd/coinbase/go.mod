module github.com/jsirianni/coinbase-cli/cmd/coinbase

go 1.16

require (
	github.com/jedib0t/go-pretty/v6 v6.1.0
	github.com/jsirianni/coinbase-cli/internal v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.2.0
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
)

replace github.com/jsirianni/coinbase-cli/internal => ../../internal
