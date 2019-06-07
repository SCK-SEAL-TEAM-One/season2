
# SCK SEAL TEAM ONE Season 2
สรุปการเรียน Week 1 ( 4-6 มิถุนายน 2562 )

## Project

- [start-go-with-greeting](https://github.com/SCK-SEAL-TEAM-One/season2/tree/master/start-go-with-greeting "start-go-with-greeting") เริ่มต้นภาษา GO
 - [VendingMachine](https://github.com/SCK-SEAL-TEAM-One/season2/tree/master/VendingMachine "VendingMachine") การทำโจทย์ตู้กดน้ำโดยใช้ TDD ด้วยภาษา GO
 - [go-simple-web](https://github.com/SCK-SEAL-TEAM-One/season2/tree/master/go-simple-web "go-simple-web") การสร้าง API สำหรับใช้งานจากโจทย์ตู้กดน้ำ

## ภาษาโก (Go language)
ภาษาโก เป็นภาษาที่ช่วยในการลดปัญหาต่างๆที่เกิดขึ้นในภาษาอื่น มีขนาดโปรแกรมที่เล็กส่งผลให้มีการใช้เวลาในการประมวลผลได้เร็ว และมีเครื่องมือที่ใช้ในการทดสอบติดมากับภาษาให้ด้วย
	
ได้รับความรู้เพิ่มเรื่อง 
	- function
	- method
	- struct
	- map

**คำสั่งในภาษาที่มีใช้บ่อย**
รันผลการทดสดสอบ

    go test 

 เพื่อแสดง coverage ของการทดสอบ

    go test -cover

build file จาก Go

    go build <file> 

 run Go file

    go run <file>
จัด format Go file

    go fmt <file> 
สร้าง module file

    go mod init <module> 
export ผลการทดสอบ coverage ไปยังไฟล์ count.out

    go test -coverprofile=count.out 
นำไฟล์ count.out มาแสดงผลบน html

    go tool cover -html=count.out 

	
## TDD : Test Driven Development
Test Driven Development คือการสร้าง Test Case ก่อนที่จะเริ่มการ Coding โดยเราจะทำการแก้ไข Code ของเราเพื่อให้ Test ที่สร้างไว้ก่อนนั้นผลการทดสอบเป็นผ่าน โดยแต่ละลำดับจะค่อยๆเห็นรูปแบบของ Code จนสามารถทำการ Refactor ได้

**Refactor Code**
	เป็นการปรับปรุงโค้ดที่เขียนแล้ว โดยยุบรวม หรือลบสิ่งที่ไม่จำเป็น เพื่อลดโค้ดซ้ำซ้อน และให้การทำงานมีคุณภาพ เช่น ลดเวลาในการ start time ลดขนาดของโปรแกรม


**Design**
	การออกแบบก่อนการเขียนโค้ดมีความสำคัญต่อการทำงาน ดังนี้
- จำกัดขอบเขตที่จำเป็นในการทำงานเพื่อให้การเขียนโค้ดไม่ให้ทำเกินที่ตั้งเป้าไว้ ทำให้เมื่อเริ่มเขียนโค้ดจะทำเฉพาะส่วนที่กำหนดและแก้ปัญหาจุดที่เราสนใจเท่านั้น จึงลดเวลาในการพัฒนาลงได้
- ทราบว่าจะใช้อะไรในการเขียนเป็นลำดับขั้นตอน 
- ทีมสามารถมองเป็นภาพรวมเดียวกันสามารถทำงานร่วมกันได้ โดยกำหนดมาตรฐานการตั้งชื่อตัวแปร ชื่อฟังก์ชัน ค่า Input Output ที่ออกมาสามารถนำไปใช้ได้ เมื่อรวมโค้ดกันจะช่วยให้ลดปัญหาที่อาจเกิดขึ้นได้
	เพราะการเขียนโดยไม่วางแผนเลยทำให้ระหว่างเขียนอาจใช้เวลามากกว่าที่ควร


## Git101
 Git ถูกสร้างและนำมาใช้เพื่อแก้ปัญหาความวุ่นวายของการจัดการ Source Code และยังสามารถติดตาม Version ตรวจดูการเปลี่ยนแปลงของ Source ได้

**คำสั่งที่มีใช้บ่อย**
ดูค่า config ของ git ใน list

    git config -l 

การเริ่มต้นการใช้ git

    git init 

ดูสถานะ

    git status 

กลับไป status เดิมตรง commit รอบที่แล้ว

    git checkout -- <file> 

ยืนยันไฟล์ที่ใช้เพื่อไป commit

    git add <file> 

ดูความแตกต่างว่ามีอะไรเปลี่ยนแปลงไปจากครั้งก่อน

    git diff 

ยกเลิกที่ใช้คำสั่ง git add เผื่อแก้ไขเพิ่มเติม 

    git reset HEAD <file> 

 แก้ commit message *แก้ได้เฉพาะ  commit ล่าสุด

    git commit --amend

 clone repository

    git clone <URL>

นำ commit ไปอยู่บน repository ถ้าไฟล์อัพซ้อนทับกันจะช่วยให้ทราบว่าไม่สามารถอัพได้ ต่อง git push ลงมาแล้วจัดการว่าต้องการการเปลี่ยนแปลงแบบใด หรือต้องการทั้งคู่มาต่อกัน

    git push 

 การสร้างไฟล์ .gitignore สำหรับเก็บที่อยู่ของไฟล์ที่เราต้องการยกเว้น (กรณียกเว้นไฟล์ .out ทั้งหมด)

    echo *.out > .gitignore

