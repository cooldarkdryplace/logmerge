# logmerge
Tiny tool for merging multiple log files. Reads stream, so should not use a lot of RAM.  
Logs are merged based on parsed timestamp. You should provide `-ts` timestamp format that corresponds to the one in your logs.
The default is: `Jan  2 15:04:05`.  

Please refer to time package documentation for more info.
https://golang.org/pkg/time/#pkg-constants

Tested with 3x12GB log files.  

Example:  
File #1
```
Aug  1 00:13:58 Preparation is finished. Wish us good luck.
Sep  25 11:50:04 Did I tell that the weather is just great? It feels more like a vacation.
Sep  30 06:30:41 Today I saw the last seagul, looks like we are far from big land now.
```

File #2
```
Sep  8 00:13:58 We have plenty of food and drinking water.
Sep  19 00:30:24 Thanks to a strong wind we quickly move tovards new adventures.
Oct  2 14:00:42 Captain and several sailors are injured. Doctor is here, hopefully we will be able to continue.
```

File #3
```
Oct  12 04:05:42 Storm was just incredible. And we lost both anchors.
Oct  20 18:44:08 Today is my birthday, and I am afraid that it is my last birthday.
```

Merge should result in:
```
Aug  1 00:13:58 Preparation is finished. Wish us good luck. 
Sep  8 00:13:58 We have plenty of food and drinking water.
Sep  19 00:30:24 Thanks to a strong wind we quickly move tovards new adventures.
Sep  25 11:50:04 Did I tell that the weather is just great? It feels more like a vacation.
Sep  30 06:30:41 Today I saw the last seagul, looks like we are far from big land now.
Oct  2 14:00:42 Captain and several sailors are injured. Doctor is here, hopefully we will be able to continue.
Oct  12 04:05:42 Storm was just incredible. And we lost both anchors.
Oct  20 18:44:08 Today is my birthday, and I am afraid that it is my last birthday.
```

## Installation
`go get github.com/cooldarkdryplace/logmerge`

## Usage
`logmerge -i your_input_directory -o output.file -tf Jan  2 15:04:05`  

`-i`  Input directory. Place where multiple files are located. All files considered to be logs.  
`-o`  Output file. Merge result.  
`-ts` Time format of your timestamps. Example: `Jan  2 15:04:05`.  
