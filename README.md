# Go concurrency test

This is a small project I made to study how Go's concurrency model works. The demos currently are:

## Terminal Race
You can choose the amount of cars that will race. The terminal buffer is updated each milissecond with the new cars' positions. For now, there are no ties.

_An important note: If your terminal starts "blinking" more than once, its because the amount of cars you provided to the program is greater than the amount that can show up in your terminal without scrolling down (because of its font size). You can bypass that by making the font size of your terminal smaller to fit all the cars on-screen at once, since the buffer has to be "re-rendered" and the cursor moved way down to the last character._

---------------------------------------------
### Feel free to make PRs and point some things out if you really understand the language; I'm open to learn it!