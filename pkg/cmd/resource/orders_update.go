package resource

import (
	"net/http"

	"github.com/spf13/cobra"

	"github.com/stripe/stripe-cli/pkg/config"
)

// OrdersUpdateCmd represents the order update API operation command. This
// command is manually defined because of quirkiness with the OpenAPI spec autogenerated command and required parameters.
type OrdersUpdateCmd struct {
	opCmd *OperationCmd
}

func (ouc *OrdersUpdateCmd) runOrdersUpdateCmd(cmd *cobra.Command, args []string) error {
	return ouc.opCmd.runOperationCmd(cmd, args)
}

// NewOrdersUpdateCmd creates a new orders creation sub command.
func NewOrdersUpdateCmd(parentCmd *cobra.Command, cfg *config.Config) *OrdersUpdateCmd {
	ordersUpdateCmd := &OrdersUpdateCmd{
		opCmd: NewOperationCmd(parentCmd, "update", "/v1/orders/{id}", http.MethodPost, map[string]string{
			"currency":               "string",
			"line_items[][product]":  "string",
			"line_items[][quantity]": "integer",
		}, cfg),
	}

	ordersUpdateCmd.opCmd.Cmd.RunE = ordersUpdateCmd.runOrdersUpdateCmd

	return ordersUpdateCmd
}