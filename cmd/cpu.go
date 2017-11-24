package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	cAdvisor "github.com/google/cadvisor/client/v2"
	v1 "github.com/google/cadvisor/info/v1"
	v2 "github.com/google/cadvisor/info/v2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "cpu",
	//Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := cAdvisor.NewClient(cAdvisorUrl)
		if err != nil {
			log.Fatal(err)
		}
		request := v2.RequestOptions{
			IdType:    "docker",
			Count:     1,
			Recursive: false,
		}
		containers, err := client.Stats(containerID, &request)
		if err != nil {
			log.Fatal(err)
		}
		for _, container := range containers {
			switch metrics {
			case "all":
				allCpuMetrics(container.Stats[0].Cpu)
			case "total":
				fmt.Println(container.Stats[0].Cpu.Usage.Total)
			case "user":
				fmt.Println(container.Stats[0].Cpu.Usage.User)
			case "system":
				fmt.Println(container.Stats[0].Cpu.Usage.System)
			case "load-average":
				fmt.Println(container.Stats[0].Cpu.LoadAverage)
			default:
				allCpuMetrics(container.Stats[0].Cpu)
			}
		}
	},
}

func allCpuMetrics(c *v1.CpuStats) {
	data := [][]string{
		[]string{"total", strconv.FormatUint(c.Usage.Total, 10)},
		[]string{"user", strconv.FormatUint(c.Usage.User, 10)},
		[]string{"system", strconv.FormatUint(c.Usage.System, 10)},
		[]string{"load-average", strconv.FormatInt(int64(c.LoadAverage), 10)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Time"})
	table.SetBorders(
		tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false},
	)
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
}

func init() {
	resourceCmd.AddCommand(cpuCmd)
	cpuCmd.Flags().StringVarP(&containerID, "container-id", "c", "", "Container ID")
	cpuCmd.Flags().StringVarP(&metrics, "metrics", "m", "all", "Metrics")
}
