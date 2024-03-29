# 工具

> 简单介绍每个工具及其用法

# 测试代码生成工具

`go run . -h`
`go run . -n [question number] [-i] [-f]`

根据指定参数与路径定义规则，获取并解析对应文件中的第一个函数，为其生成测试代码；也可以用来生成符合规则的目录结构

关键步骤实现说明：

1. 整体实现采用填充模板的方式，根据待测试函数的关键信息填充模板，然后将模板写入文件
    1. 工具也因此受到模板的制约：只能为符合模板的函数，正确生成测试代码
2. 按行读取目标文件，使用正则匹配待测试函数，正则表达式：`^func (\w+)\((.*)\)(.*){$`
3. 在解析函数声明的时候，使用了AST的概念和FSM技术
    1. AST：abstract syntax tree，抽象语法树
    2. 本工具仅使用AST部分概念
    3. FSM：finite state machine，有限状态机
