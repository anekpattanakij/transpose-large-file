# transpose-large-file
Example to transpose large csv file by using GoLang

This program will transponse csv file from

```
,column1,column2
row1,value1,value2
row2,value3,value4
```

To

```
column1,row1,value1
column1,row2,value2
column2,row3,value3
column2,row4,value4
```

This program is using Reader.ReadLine(), low level functions, to handle large data set in one row, sinceScanner could not handle its due to memony limitation.
