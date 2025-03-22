package main

import (
	"my-go-project/cmd"
)

func main() {
	// بدء Publisher في Go Routine منفصل
	go cmd.Publisher()

	// بدء Subscriber في Go Routine منفصل
	go cmd.Subscriber()

	// إضافة تأخير (اختياري) لإعطاء الوقت للجميع ليتم التنفيذ قبل إنهاء البرنامج
	// في حالة عدم وجود هذه السطر قد ينتهي البرنامج قبل تنفيذ الدوال
	// في حالة استخدام تطبيقات حقيقية قد تحتاج إلى تقنيات أخرى مثل WaitGroup
	select {} // تجعل البرنامج ينتظر حتى لا ينتهي فوراً
}
