# mailtemplate
增强对open-falcon提供的alarm模块邮件模板支持。  
本插件对falcon-alarm的源码有一定的侵入性。  
集成到falcon-alarm后，进行统一编译，并生成新的falcon-alarm二进制包。  
本功能只是本人在使用falcon-alarm功能时，感觉告警邮件内容有些不可定制化，于是做了这样一个基于模板的插件功能。  
由于是因为falcon才使本人学习GO语言，在此感谢falcon，此功能也当做是一种练习。  
如有问题，请及时联系。

# 自定义模板样式
```
报警状态：[PROBLEM]
报警级别：[P0]
报警主机：[localhost]
监控指标：[agent.alive[客户端存活]]
监控标签：[]
报警阀值：[all(#1) 1 == 1]
报警次数：当前第 7 次，最大 7 次
报警时间：[2016-05-04 14:11:00]
规则配置：[ http://127.0.0.1:8899/template/view/1 ]
```

# 使用说明：
  * 下载插件代码  
      `[root@localhost work]#git clone https://github.com/mircoteam/mailtemplate.git`
  * 将下载后的alarm文件夹中的文件复制到falcon-alarm文件夹下  
      `[root@localhost alarm]#cd mailtemplate`  
      `[root@localhost alarm]#cp alarm $WORKSPACE/alarm/`  
  * 修改falcon-alarm源码中cron/builder.go文件中BuildCommonMailContent方法的代码：  
      >增加导入包  
        `"github.com/open-falcon/alarm/extends/mailtemplate"`  
      >去除导入包  
        `//"github.com/open-falcon/alarm/g"`  
      >修改前:  

        func BuildCommonMailContent(event *model.Event) string {
	        link := g.Link(event)
	        return fmt.Sprintf(
      		"%s\r\nP%d\r\nEndpoint:%s\r\nMetric:%s\r\nTags:%s\r\n%s: %s%s%s\r\nNote:%s\r\nMax:%d, Current:%d\r\nTimestamp:%s\r\n%s\r\n",
      		event.Status,
      		event.Priority(),
      		event.Endpoint,
      		event.Metric(),
      		utils.SortedTags(event.PushedTags),
      		event.Func(),
      		utils.ReadableFloat(event.LeftValue),
      		event.Operator(),
      		utils.ReadableFloat(event.RightValue()),
      		event.Note(),
      		event.MaxStep(),
      		event.CurrentStep,
      		event.FormattedTime(),
      		link,
      	  )
        }

      >修改后:

          func BuildCommonMailContent(event *model.Event) string {
          	return mailtemplate.BuildCommonMailContent(event)
          }
  * 编译生成二进制文件
    `./control build`

# 扩展说明
  如果想做其它指标的展现，可以对extends/mailtemplate.go文件中MailTmpl结构进行扩展。然后在mail.tpl中进行展现即可
