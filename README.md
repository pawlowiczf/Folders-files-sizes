# Biggest folders, files, sorted by size 
Get list of biggest folders or biggest files in a particular directory:
- ***--amount*** - specifies the number of listed files, directories (default: 5)
- ***--mode*** - specifies the mode: **biggest-dirs** (get list of biggest dirs in current directory) or **biggest-files** (get list of biggest files in current directory) (default: biggest-dirs)
- ***--dir*** - specifies the directory to in which perform search (default: ".")
## Getting started
1. Download foldersize.exe or foldersize.  
2. Add path to foldersize.exe in environmental variables.  
3. Run terminal as admin.  
4. Navigate through folders and run following commands:
```sh
foldersize.exe --mode biggest-dirs --amount 5 --dir "C:\"
foldersize.exe --mode biggest-files --amount 6
```
