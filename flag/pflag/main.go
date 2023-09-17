package main

// pflag is a drop-in replacement of Go's native flag package.
// If you import pflag under the name "flag" then all
// code should continue to function with no changes.
import (
	"fmt"
	"strings"

	flag "github.com/spf13/pflag" // 导包这里直接使用 flag 的名称，可无痛接入 go 标准库的 flag
)

func main() {
	// demo_1: 声明一个整数flag “-flagname” ，存储在指针 ip 中，类型为 *int。
	var ip *int = flag.Int("flagname_1", 1234, "some help message for flagname_1")
	fmt.Println(*ip) // 1234

	// demo_2: 如果您愿意，还可以使用 Var() 函数将标记与变量绑定。
	var flagvar int
	flag.IntVar(&flagvar, "flagname_2", 5678, "some help message for flagname_2")
	fmt.Println(flagvar) // 5678

	// 这样就可以直接使用flag。如果使用flag本身，它们都是指针；如果绑定到变量，它们都是值。
	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

	// demo_3: 您也可以创建满足 Value 接口（带有指针接收器）的自定义标记，并通过以下方式将它们与标记解析结合起来
	// 参考这里：https://juejin.cn/post/6947349909783707662
	var varprovite arrayFlags
	flag.Var(&varprovite, "这里是varprovite的默认值", "some help message for flagname")
	fmt.Println(varprovite)

	// 定义完所有 flag 后，调用 flag.Parse() 将命令行解析为已定义的 flag。
	flag.Parse()

	// ---
	// 如果您有一个 FlagSet，但发现很难跟上代码中的所有指针，
	// 那么可以使用辅助函数来获取存储在 Flag 中的值。
	// 如果您有一个 pflag.FlagSet，其中有一个名为 "flagname "的 int 类型的标志，
	// 您可以使用 GetInt() 来获取 int 值。但请注意，"flagname "必须存在，而且必须是 int。
	// GetString("flame")将失效。
	myflagset := flag.NewFlagSet("myflagset", flag.ContinueOnError)
	var intValue int
	myflagset.IntVar(&intValue, "flagname_4", 7890, "some help message for flagname_4")
	i, err := myflagset.GetInt("flagname_4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("i has value ", i)

	//
	// Parsed Arguments	Resulting Value
	// --flagname_5=1357	lookupval=6666
	// --flagname_5	lookupval=4321
	// [nothing]	lookupval=6666
	var lookupval = flag.IntP("flagname_5", "f", 6666, "help message")
	flag.Lookup("flagname_5").NoOptDefVal = "4321"
	fmt.Println("lookupval has value ", *lookupval)

}

// arrayFlags implements pflag.Value interface.
// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
//
//	type Value interface {
//		String() string
//		Set(string) error
//		Type() string
//	}
type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ", ")
}
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *arrayFlags) Type() string {
	return "slice"
}
