// go query

// NewDocumentFromNode *Document: 传入*html.Node对象，也就是根结点
// NewDocument(url string) *Document, error 传入URL，内部用http.Get获取网页
// NewDocumentFromReader 传入io.Reader，内部从reader中读取内容并解析
// NewDocumentFromResponse 传入HTTP响应，内部拿到res.Body

// Find(), First(), Last(), Eq(), Slice(start, end int)

// 获取一个标签的内容 Text() 返回字符串，可能前后有很多空格，可以视情况做清除

// 获取属性值也很容易，有两个方法
	// Attr 返回属性值和该属性是否存在，类型从map中取值
	// AttrOr 和上一个方法类似，区别在于如果属性不存在，则返回给定的默认值

// goquery 提供了三个用于迭代的方法
	// Each(f func(int, *Selection)) *Selection f第一个参数是当前的下标，第二个是当前的节点
	// EachWithBreak(f func(int, *Selection) bool) *Selection和Each类似，增加了中途退出循环的能力
	// Map

// 删除节点 调用Remove() 方法可以把一个节点删掉 在节点后插入html AfterHtml("<br>")

// 在标签尾部append一段内容 AppendSelection