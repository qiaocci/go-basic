什么是goroutine泄露?

通俗来说，开启了一个goroutine，用完后，我们要正确让其结束。如果它没用了，还没结束，那就是goroutine leak。
泄漏的goroutine占用一部分cpu，还可能占着一些其他资源，从而影响主协程效率，有时甚至产生异常。