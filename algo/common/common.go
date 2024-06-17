package common

// 这是个接口
type DecoderInterface interface {
	Validate() error // 方法声明
	Decode() error
	GetCoverImage() []byte
	GetAudioData() []byte
	GetAudioExt() string
	GetMeta() MetaInterface
}

type MetaInterface interface {
	GetArtists() []string
	GetTitle() string
	GetAlbum() string
}


// type ExampleInterface interface {
//     DoWork() string   // 方法声明
//     GetData() int      // 方法声明
// }
// 要实现接口，一个类型只需提供接口中所有声明的方法的具体实现。例如，假设我们有一个结构体 ExampleStruct：
// type ExampleStruct struct {
//     // 结构体字段
// }
// 然后，我们为这个结构体提供接口中声明的所有方法的实现：

// func (e *ExampleStruct) DoWork() string {
//     // 方法实现
//     return "Work done"
// }

// func (e *ExampleStruct) GetData() int {
//     // 方法实现
//     return 42
// }

// 一旦一个类型实现了接口，你可以使用该接口类型的变量来存储任何实现了该接口的类型的值，这允许你编写更通用的代码：

// func workWithInterface(e ExampleInterface) {
//     fmt.Println(e.DoWork())
//     fmt.Println(e.GetData())
// }
// 这个函数可以接受任何实现了 ExampleInterface 的类型的值。