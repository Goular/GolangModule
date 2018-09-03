Go������ʱ�ĵ�ǰ��ƣ��ٶ�����Ա�Լ��������ʱ��ֹһ��goroutine�Լ���ʱ��ֹ�ó��� ����ͨ������os.Exit���main()�����ķ�������������ʽ��ֹ���򡣶���ʱ��������Ҫ����ʹ������������һ�С�

ʹ��sync.WaitGroup
	һֱ�ȴ�ֱ��WaitGroup����0

		package main

		import "sync"

		func main() {
			var wg sync.WaitGroup
			wg.Add(1)
			wg.Wait()
		}

��select
	select{}��һ��û���κ�case��select������һֱ����

		package main

		func main() {
			select{}
		}
		
��ѭ��
	��Ȼ������������100%ռ��һ��cpu��������ʹ��

		package main

		func main() {
			for {}
		}
		
��sync.Mutex
	һ���Ѿ����˵���������һ�λ�һֱ���������������ʹ��

		package main

		import "sync"

		func main() {
			var m sync.Mutex
			m.Lock()
			m.Lock()
		}
		
os.Signal
	ϵͳ�ź�������go����Ҳ�Ǹ�channel�����յ��ض�����Ϣ֮ǰһֱ����

		package main

		import (
			"os"
			"syscall"
			"os/signal"
		)

		func main() {
			sig := make(chan os.Signal, 2)
			signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
			<-sig
		}
		
��channel����nil channel
	channel��һֱ����ֱ���յ���Ϣ��nil channel��Զ������

		package main

		func main() {
			c := make(chan struct{})
			<-c
		}
		package main

		func main() {
			var c chan struct{} //nil channel
			<-c
		}
�ܽ�
	ע������д�ĵĴ���󲿷ֲ���ֱ�����У�����panic����ʾ��all goroutines are asleep - deadlock!������Ϊgo��runtime���������е�goroutine����ס�ˣ� û��һ��Ҫִ�С����������������ǰ�����һ���������Լ�ҵ���߼���goroutine�������Ͳ���deadlock�ˡ�

�ο�
	http://pliutau.com/different-ways-to-block-go-runtime-forever/