package cmd

import (
	"github.com/spf13/cobra"

	"stegia/internal/logger"
	"stegia/internal/totvs/companies"
	"stegia/internal/totvs/factory"
	"stegia/internal/totvs/suppliers"
	"stegia/internal/util"
)

var suppliersCmd = &cobra.Command{
	Use:   "suppliers",
	Short: "Supplier (fornecedor) operations",
}

var suppliersAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a supplier from a TOON file (lists companies first)",
	RunE: func(cmd *cobra.Command, args []string) error {
		log := logger.New()

		toonPath, _ := cmd.Flags().GetString("file")
		companyId, _ := cmd.Flags().GetString("company-id")

		env := util.LoadTotvsEnv()
		log.Info("loaded env", "envFile", env.EnvFile, "hostname", env.Hostname)

		cli := factory.ClientFactory{Log: log}.New(env)

		sf := factory.ServiceFactory{Log: log}
		companiesSvc := sf.CompaniesService(cli)
		suppliersSvc := sf.SuppliersService(cli)

		companiesCtrl := &companies.Controller{
			Service: companiesSvc,
			Builder: companies.Builder{},
			Log:     log,
		}

		suppliersCtrl := &suppliers.Controller{
			Service:   suppliersSvc,
			Builder:   suppliers.Builder{},
			Companies: companiesCtrl,
			Log:       log,
		}

		return suppliersCtrl.AddFromTOON(toonPath, companyId)
	},
}

func init() {
	totvsCmd.AddCommand(suppliersCmd)
	suppliersCmd.AddCommand(suppliersAddCmd)

	suppliersAddCmd.Flags().StringP("file", "f", "", "Path to .toon file (TOON format)")
	_ = suppliersAddCmd.MarkFlagRequired("file")
	suppliersAddCmd.Flags().String("company-id", "", "CompanyId header value (optional; auto-selected if omitted)")
}
