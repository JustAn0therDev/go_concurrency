# Go concurrency test

This is a small project I made to study how Go's concurrency model works. The demos currently are:

## Terminal Race - Multithreading with multiple goroutines
You can choose the amount of cars that will race. The terminal buffer is updated each milissecond with the new cars' positions. For now, there are no ties.

The terminal race works by updating a string with various threads and checking which one is "complete". It shows a way to work with threads in Go.

_An important note: If your terminal starts "blinking" more than once, its because the amount of cars you provided to the program is greater than the amount that can show up in your terminal without scrolling down (because of its font size). You can bypass that by making the font size of your terminal smaller to fit all the cars on-screen at once, since the buffer has to be "re-rendered" and the cursor moved way down to the last character._

## Progress Bar - Multithreading with only two goroutines
It simulates a progress bar loading, which is random.

The only thing really multitask in this is the fact that the terminal buffer is updated on the main thread while another thread updates the progress bar string.

## Pass the Bat - Multithreading with multiple goroutines and channels
There are lines and "people" running from right to left on those lines and you can choose the amount of lines that will update. It works like Terminal Race, the difference is that the way the runners stop and let the next run uses channels, to show an example of what can be done with the feature.

---------------------------------------------
### Feel free to make PRs and point some things out if you really understand the language; I'm open to learn it!