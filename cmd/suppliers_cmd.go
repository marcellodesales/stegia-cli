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

var suppliersViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a cached supplier by id (reads examples/suppliers/<id>.toon)",
	RunE: func(cmd *cobra.Command, args []string) error {
		log := logger.New()

		id, _ := cmd.Flags().GetString("id")
		format, _ := cmd.Flags().GetString("format")

		env := util.LoadTotvsEnv()
		log.Info("loaded env", "envFile", env.EnvFile, "hostname", env.Hostname)

		cli := factory.ClientFactory{Log: log}.New(env)
		sf := factory.ServiceFactory{Log: log}
		suppliersSvc := sf.SuppliersService(cli)

		ctrl := &suppliers.Controller{
			Service: suppliersSvc,
			Builder: suppliers.Builder{},
			Log:     log,
		}
		return ctrl.ViewFromCache(id, format)
	},
}

func init() {
	totvsCmd.AddCommand(suppliersCmd)
	suppliersCmd.AddCommand(suppliersAddCmd)

	suppliersAddCmd.Flags().StringP("file", "f", "", "Path to .toon file (TOON format)")
	_ = suppliersAddCmd.MarkFlagRequired("file")
	suppliersAddCmd.Flags().String("company-id", "", "CompanyId header value (optional; auto-selected if omitted)")

	suppliersViewCmd.Flags().String("id", "", "Supplier ID (e.g., SUP-902341)")
    suppliersViewCmd.Flags().StringP("format", "f", "toon", "Output format: toon or json")
	_ = suppliersViewCmd.MarkFlagRequired("id")
    suppliersCmd.AddCommand(suppliersViewCmd)
}
