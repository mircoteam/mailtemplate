报警状态：[{{.Status}}]
报警级别：[P{{.Priority}}]
报警主机：[{{.Endpoint}}]
监控指标：[{{.Metric}}[{{.Note}}]]
监控标签：[{{.PushedTags}}]
报警阀值：[{{.Func}} {{.LeftValue}} {{.Operator}} {{.RightValue}}]
报警次数：当前第 {{.CurrentStep}} 次，最大 {{.MaxStep}} 次
报警时间：[{{.FormattedTime}}]
规则配置：[{{ .Link }}]