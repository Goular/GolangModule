package com.goular.demo;

public class Demo01 {

	public static void main(String[] args) {
		new Thread(new Test_Thread_1()).start();
		new Thread(new Test_Thread_2()).start();
	}
}

class Test_Thread_1 implements Runnable {

	@Override
	public void run() {
		for (int i = 0; i < 100; i++) {
			System.out.println("Thread 1 : " + i);
		}
	}

}

class Test_Thread_2 implements Runnable {

	@Override
	public void run() {
		for (int i = 100; i < 200; i++) {
			System.out.println("Thread 2 : " + i);
		}
	}

}
