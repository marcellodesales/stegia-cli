package cmd

import (
	"github.com/spf13/cobra"

	"stegia/internal/logger"
	"stegia/internal/totvs/companies"
	"stegia/internal/totvs/factory"
	"stegia/internal/util"
)

var companyStatus string

var companiesCmd = &cobra.Command{
	Use:   "companies",
	Short: "Company operations",
}

var companiesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List companies (GET /api/btb/v1/companies)",
	RunE: func(cmd *cobra.Command, args []string) error {
        log := logger.New(logger.Level(logLevel))

		env := util.LoadTotvsEnv()
		log.Info("loaded env", "envFile", env.EnvFile, "hostname", env.Hostname)

		cli := factory.ClientFactory{Log: log}.New(env)
		svc := factory.ServiceFactory{Log: log}.CompaniesService(cli)

		ctrl := &companies.Controller{
			Service: svc,
			Builder: companies.Builder{},
			Log:     log,
		}
        return ctrl.ListAndPrint(companyStatus)
	},
}

func init() {
	totvsCmd.AddCommand(companiesCmd)
	companiesCmd.AddCommand(companiesListCmd)

	companiesListCmd.Flags().StringVar(
		&companyStatus,
		"status",
		"",
		"Filter by company status (e.g. ACTIVE, INACTIVE)",
	)
}
