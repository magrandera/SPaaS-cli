package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a app from the server",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serverDefined()
		tokenDefined()
		token := viper.GetString("token")
		client := &http.Client{Transport: tr}
		url := viper.GetString("url") + "/api/app/" + args[0]
		req, _ := http.NewRequest("DELETE", url, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		isLoggedIn(res)
		scanner := bufio.NewScanner(res.Body)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := Application{}
			err := json.Unmarshal([]byte(scanner.Text()), &line)
			if err != nil {
				color.Red(err.Error())
			}
			switch line.Type {
			case "info":
				color.Yellow(fmt.Sprintf("%-14v", "Info:") + line.Message)
			case "error":
				color.Red(fmt.Sprintf("%-14v", "Error:") + line.Message)
			case "success":
				color.Green(fmt.Sprintf("%-14v", "Success:") + line.Message)
				if line.Extended != nil {
					for _, value := range line.Extended {
						color.Green("    " + fmt.Sprintf("%-14v", value.Key) + value.Value.(string))
					}
				}
			default:
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
