1. ไฟล์งานจะอยู่ใน server.go ครับ
2. ใช้ mysql เป็น database ครับ และใช้โปรแกรม db เป็น mysql workbench ครับ
3. ไฟล์ table ได้ทำการ export data ให้มัน create schema ไว้แล้วชื่อไฟล์ echoHW.sql ครับ
4. เวลา run บน broswer ใช้ http://localhost:1323/ ตาม coding นี้ :
   e.Logger.Fatal(e.Start(":1323"))
5. ในการเชื่อม db ใช้ id เป็น root ส่วน password เป็น 1234 แล้วชื่อ db ใช้เป็น godb ส่วนชื่อ host เป็น localhost port เป็น 3306 ตาม coding ที่ใช้เชื่อม db :
   db, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True")
6. ไฟล์ table ได้ทำการ export ไว้แล้วชื่อไฟล์ echoHW.sql
7. ผมได้ใช้ postman ในการเทส api แล้วโดยได้ลองในส่วนของ get , post แล้วก็ delete แล้วทำให้ข้อมูลใน table ทุก table เหลือเพียงข้อมูลเดียวโดย URL ที่เทสจะมีดังนี้ 
  เทส get
  http://localhost:1323/
  http://localhost:1323/bookstores/
  http://localhost:1323/bookstores/1
  http://localhost:1323/memberships/
  http://localhost:1323/memberships/1
  http://localhost:1323/orders/
  http://localhost:1323/orders/1
  เทส post 
  http://localhost:1323/bookstores/
  http://localhost:1323/memberships/
  http://localhost:1323/orders/
